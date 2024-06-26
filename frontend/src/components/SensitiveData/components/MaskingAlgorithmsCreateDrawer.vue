<template>
  <Drawer :show="show" @close="$emit('dismiss')">
    <DrawerContent :title="$t('common.add')">
      <div
        class="w-[40rem] max-w-[calc(100vw-5rem)] space-y-6 divide-y divide-block-border"
      >
        <div class="space-y-6">
          <div class="sm:col-span-2 sm:col-start-1">
            <label for="title" class="textlabel">
              {{ $t("settings.sensitive-data.algorithms.table.title") }}
              <span class="text-red-600 mr-2">*</span>
            </label>
            <NInput
              v-model:value="state.title"
              :placeholder="t('settings.sensitive-data.algorithms.table.title')"
              :disabled="state.processing || readonly"
            />
          </div>
          <div class="sm:col-span-2 sm:col-start-1">
            <label for="description" class="textlabel">
              {{ $t("settings.sensitive-data.algorithms.table.description") }}
            </label>
            <NInput
              v-model:value="state.description"
              :placeholder="
                t('settings.sensitive-data.algorithms.table.description')
              "
              :disabled="state.processing || readonly"
            />
          </div>
          <div class="w-full mb-6 space-y-1">
            <label for="masking-type" class="textlabel">
              {{ $t("settings.sensitive-data.algorithms.table.masking-type") }}
              <span class="text-red-600 mr-2">*</span>
            </label>
            <RadioGrid
              :value="state.maskingType"
              :options="maskingTypeList"
              :disabled="readonly"
              class="grid-cols-3 gap-2"
              @update:value="onMaskingTypeChange($event as MaskingType)"
            >
              <template #item="{ option }: RadioGridItem<MaskingType>">
                {{ option.label }}
              </template>
            </RadioGrid>
          </div>
        </div>
        <div class="space-y-6 pt-6">
          <template v-if="state.maskingType === 'full-mask'">
            <div class="sm:col-span-2 sm:col-start-1">
              <label for="substitution" class="textlabel">
                {{
                  $t(
                    "settings.sensitive-data.algorithms.full-mask.substitution"
                  )
                }}
                <span class="text-red-600 mr-2">*</span>
              </label>
              <p class="textinfolabel">
                {{
                  $t(
                    "settings.sensitive-data.algorithms.full-mask.substitution-label"
                  )
                }}
              </p>
              <NInput
                v-model:value="state.fullMask.substitution"
                :placeholder="
                  t('settings.sensitive-data.algorithms.full-mask.substitution')
                "
                class="mt-2"
                :disabled="state.processing || readonly"
              />
            </div>
          </template>
          <template v-if="state.maskingType === 'range-mask'">
            <p class="textinfolabel">
              {{ $t("settings.sensitive-data.algorithms.range-mask.label") }}
            </p>
            <div
              v-for="(slice, i) in state.rangeMask.slices"
              :key="i"
              class="flex space-x-2 items-center"
            >
              <div class="flex-none flex flex-col gap-y-1">
                <label for="slice-start" class="textlabel flex">
                  {{
                    $t(
                      "settings.sensitive-data.algorithms.range-mask.slice-start"
                    )
                  }}
                  <span class="text-red-600 mr-2">*</span>
                </label>
                <NInputNumber
                  :value="slice.start"
                  :min="0"
                  :placeholder="
                    t(
                      'settings.sensitive-data.algorithms.range-mask.slice-start'
                    )
                  "
                  :disabled="state.processing || readonly"
                  style="width: 5rem"
                  @update:value="onSliceStartChange(i, $event)"
                />
              </div>
              <div class="flex-none flex flex-col gap-y-1">
                <label for="slice-end" class="textlabel">
                  {{
                    $t(
                      "settings.sensitive-data.algorithms.range-mask.slice-end"
                    )
                  }}
                  <span class="text-red-600 mr-2">*</span>
                </label>
                <NInputNumber
                  :value="slice.end"
                  :min="1"
                  :placeholder="
                    t('settings.sensitive-data.algorithms.range-mask.slice-end')
                  "
                  :disabled="state.processing || readonly"
                  style="width: 5rem"
                  @update:value="onSliceEndChange(i, $event)"
                />
              </div>
              <div class="flex-1 flex flex-col gap-y-1">
                <label for="substitution" class="textlabel">
                  {{
                    $t(
                      "settings.sensitive-data.algorithms.range-mask.substitution"
                    )
                  }}
                  <span class="text-red-600 mr-2">*</span>
                </label>
                <NInput
                  v-model:value="slice.substitution"
                  :placeholder="
                    t(
                      'settings.sensitive-data.algorithms.range-mask.substitution'
                    )
                  "
                  :disabled="state.processing || readonly"
                />
              </div>
              <div class="h-[34px] flex flex-row items-center self-end">
                <MiniActionButton
                  :disabled="state.processing || readonly"
                  @click.stop="removeSlice(i)"
                >
                  <TrashIcon class="w-4 h-4" />
                </MiniActionButton>
              </div>
            </div>
            <div v-if="rangeMaskErrorMessage" class="text-red-600">
              {{ rangeMaskErrorMessage }}
            </div>
            <NButton
              class="ml-auto"
              :disabled="state.processing || readonly"
              @click.prevent="addSlice"
            >
              {{ $t("common.add") }}
            </NButton>
          </template>
          <template v-if="state.maskingType === 'md5-mask'">
            <div class="sm:col-span-2 sm:col-start-1">
              <label for="salt" class="textlabel">
                {{ $t("settings.sensitive-data.algorithms.md5-mask.salt") }}
                <span class="text-red-600 mr-2">*</span>
              </label>
              <p class="textinfolabel">
                {{
                  $t("settings.sensitive-data.algorithms.md5-mask.salt-label")
                }}
              </p>
              <NInput
                v-model:value="state.md5Mask.salt"
                :placeholder="
                  t('settings.sensitive-data.algorithms.md5-mask.salt')
                "
                class="mt-2"
                :disabled="state.processing || readonly"
              />
            </div>
          </template>
        </div>
      </div>
      <template #footer>
        <div class="w-full flex justify-between items-center">
          <div class="w-full flex justify-end items-center gap-x-3">
            <NButton @click.prevent="$emit('dismiss')">
              {{ $t("common.cancel") }}
            </NButton>
            <NTooltip :disabled="!errorMessage">
              <template #trigger>
                <NButton
                  v-if="!readonly"
                  :disabled="isSubmitDisabled"
                  type="primary"
                  @click.prevent="onUpsert"
                >
                  {{ $t("common.confirm") }}
                </NButton>
              </template>
              {{ errorMessage }}
            </NTooltip>
          </div>
        </div>
      </template>
    </DrawerContent>
  </Drawer>
</template>
<script setup lang="ts">
import { cloneDeep } from "lodash-es";
import { TrashIcon } from "lucide-vue-next";
import { NInput, NInputNumber } from "naive-ui";
import { computed, watch, reactive } from "vue";
import { useI18n } from "vue-i18n";
import type { RadioGridOption, RadioGridItem } from "@/components/v2";
import {
  Drawer,
  DrawerContent,
  RadioGrid,
  MiniActionButton,
} from "@/components/v2";
import { pushNotification, useSettingV1Store } from "@/store";
import {
  MaskingAlgorithmSetting_Algorithm as Algorithm,
  MaskingAlgorithmSetting_Algorithm_FullMask as FullMask,
  MaskingAlgorithmSetting_Algorithm_RangeMask as RangeMask,
  MaskingAlgorithmSetting_Algorithm_MD5Mask as MD5Mask,
  MaskingAlgorithmSetting_Algorithm_RangeMask_Slice as RangeMask_Slice,
} from "@/types/proto/v1/setting_service";
import type { MaskingType } from "./utils";
import { getMaskingType } from "./utils";

interface MaskingTypeOption extends RadioGridOption<MaskingType> {
  value: MaskingType;
  label: string;
}

interface LocalState {
  processing: boolean;
  maskingType: MaskingType;
  title: string;
  description: string;
  fullMask: FullMask;
  rangeMask: RangeMask;
  md5Mask: MD5Mask;
}

const props = defineProps<{
  show: boolean;
  readonly: boolean;
  algorithm: Algorithm;
}>();

const emit = defineEmits<{
  (event: "dismiss"): void;
}>();

const defaultRangeMask = computed(() =>
  RangeMask.fromPartial({
    slices: [
      RangeMask_Slice.fromPartial({
        start: 0,
        end: 1,
        substitution: "*",
      }),
    ],
  })
);

const state = reactive<LocalState>({
  processing: false,
  maskingType: "full-mask",
  title: "",
  description: "",
  fullMask: FullMask.fromPartial({}),
  rangeMask: cloneDeep(defaultRangeMask.value),
  md5Mask: MD5Mask.fromPartial({}),
});
const { t } = useI18n();
const settingStore = useSettingV1Store();

const maskingTypeList = computed((): MaskingTypeOption[] => [
  {
    value: "full-mask",
    label: t("settings.sensitive-data.algorithms.full-mask.self"),
  },
  {
    value: "range-mask",
    label: t("settings.sensitive-data.algorithms.range-mask.self"),
  },
  {
    value: "md5-mask",
    label: t("settings.sensitive-data.algorithms.md5-mask.self"),
  },
]);

const algorithmList = computed((): Algorithm[] => {
  return (
    settingStore.getSettingByName("bb.workspace.masking-algorithm")?.value
      ?.maskingAlgorithmSettingValue?.algorithms ?? []
  );
});

watch(
  () => props.algorithm,
  (algorithm) => {
    state.title = algorithm.title;
    state.description = algorithm.description;
    state.maskingType = getMaskingType(algorithm) ?? "full-mask";
    state.fullMask = algorithm.fullMask ?? FullMask.fromPartial({});
    state.rangeMask = algorithm.rangeMask ?? cloneDeep(defaultRangeMask.value);
    state.md5Mask = algorithm.md5Mask ?? MD5Mask.fromPartial({});
  }
);

const maskingAlgorithm = computed((): Algorithm => {
  const result = Algorithm.fromPartial({
    id: props.algorithm.id,
    title: state.title,
    description: state.description,
    category: state.maskingType === "md5-mask" ? "HASH" : "MASK",
  });

  switch (state.maskingType) {
    case "full-mask":
      result.fullMask = state.fullMask;
      break;
    case "range-mask":
      result.rangeMask = state.rangeMask;
      break;
    case "md5-mask":
      result.md5Mask = state.md5Mask;
      break;
  }

  return result;
});

const rangeMaskErrorMessage = computed(() => {
  if (state.rangeMask.slices.length === 0) {
    return t("settings.sensitive-data.algorithms.error.slice-required");
  }
  for (let i = 0; i < state.rangeMask.slices.length; i++) {
    const slice = state.rangeMask.slices[i];
    if (
      Number.isNaN(slice.start) ||
      Number.isNaN(slice.end) ||
      slice.start < 0 ||
      slice.end <= 0
    ) {
      return t("settings.sensitive-data.algorithms.error.slice-invalid-number");
    }
    if (slice.start >= slice.end) {
      return t("settings.sensitive-data.algorithms.error.slice-number-range");
    }

    for (let j = 0; j < i; j++) {
      const pre = state.rangeMask.slices[j];
      if (slice.start >= pre.end || pre.start >= slice.end) {
        continue;
      }
      return t("settings.sensitive-data.algorithms.error.slice-overlap");
    }

    if (!slice.substitution) {
      return t(
        "settings.sensitive-data.algorithms.error.substitution-required"
      );
    }
    if (slice.substitution.length > 16) {
      return t("settings.sensitive-data.algorithms.error.substitution-length");
    }
  }
  return "";
});

const errorMessage = computed(() => {
  if (!state.title) {
    return t("settings.sensitive-data.algorithms.error.title-required");
  }

  switch (state.maskingType) {
    case "full-mask":
      if (!state.fullMask.substitution) {
        return t(
          "settings.sensitive-data.algorithms.error.substitution-required"
        );
      }
      if (state.fullMask.substitution.length > 16) {
        return t(
          "settings.sensitive-data.algorithms.error.substitution-length"
        );
      }
      return "";
    case "md5-mask":
      if (!state.md5Mask.salt) {
        return t("settings.sensitive-data.algorithms.error.salt-required");
      }
      return "";
    case "range-mask":
      return rangeMaskErrorMessage.value;
  }
  return "";
});

const isSubmitDisabled = computed(() => {
  if (props.readonly || state.processing) {
    return true;
  }

  if (!state.title) {
    return true;
  }

  return !!errorMessage.value;
});

const onUpsert = async () => {
  state.processing = true;

  const index = algorithmList.value.findIndex(
    (item) => item.id === maskingAlgorithm.value.id
  );
  const newList = [...algorithmList.value];
  if (index < 0) {
    newList.push(maskingAlgorithm.value);
  } else {
    newList[index] = maskingAlgorithm.value;
  }

  try {
    await settingStore.upsertSetting({
      name: "bb.workspace.masking-algorithm",
      value: {
        maskingAlgorithmSettingValue: {
          algorithms: newList,
        },
      },
    });

    pushNotification({
      module: "bytebase",
      style: "SUCCESS",
      title: t("common.updated"),
    });
    emit("dismiss");
  } finally {
    state.processing = false;
  }
};

const onMaskingTypeChange = (maskingType: MaskingType) => {
  if (state.processing || props.readonly) {
    return;
  }
  switch (maskingType) {
    case "full-mask":
      state.fullMask = FullMask.fromPartial({});
      break;
    case "range-mask":
      state.rangeMask = cloneDeep(defaultRangeMask.value);
      break;
    case "md5-mask":
      state.md5Mask = MD5Mask.fromPartial({});
      break;
  }
  state.maskingType = maskingType;
};

const addSlice = () => {
  const last = state.rangeMask.slices[state.rangeMask.slices.length - 1];
  state.rangeMask.slices.push(
    RangeMask_Slice.fromPartial({
      start: (last?.start ?? -1) + 1,
      end: (last?.end ?? 0) + 1,
      substitution: "*",
    })
  );
};

const removeSlice = (i: number) => {
  state.rangeMask.slices.splice(i, 1);
};

const onSliceStartChange = (index: number, val: number | null) => {
  if (val === null || Number.isNaN(val)) {
    return;
  }
  const slice = state.rangeMask.slices[index];
  slice.start = Math.max(0, val);
  if (slice.end <= slice.start) {
    slice.end = slice.start + 1;
  }
};

const onSliceEndChange = (index: number, val: number | null) => {
  if (val === null || Number.isNaN(val)) {
    return;
  }
  const slice = state.rangeMask.slices[index];
  slice.end = Math.max(1, val);
  if (slice.start >= slice.end) {
    slice.start = slice.end - 1;
  }
};
</script>
