<template>
  <component
    :is="link ? 'router-link' : tag"
    v-bind="bindings"
    :to="link && `/${instance.name}`"
    class="inline-flex items-center gap-x-1"
    :class="link && 'normal-link'"
  >
    <InstanceEngineIcon
      v-if="icon && iconPosition === 'prefix'"
      :instance="instance"
    />

    <span>{{ instanceName(instance) }}</span>

    <InstanceEngineIcon
      v-if="icon && iconPosition === 'suffix'"
      :instance="instance"
    />
  </component>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import InstanceEngineIcon from "@/components/InstanceEngineIcon.vue";
import type { Instance } from "@/types";
import { instanceName } from "@/utils";

const props = withDefaults(
  defineProps<{
    instance: Instance;
    tag?: string;
    link?: boolean;
    icon?: boolean;
    iconPosition?: "prefix" | "suffix";
  }>(),
  {
    tag: "span",
    link: true,
    icon: true,
    iconPosition: "prefix",
  }
);

const bindings = computed(() => {
  if (props.link) {
    return {
      to: `/${props.instance.name}`,
      activeClass: "",
      exactActiveClass: "",
      onClick: (e: MouseEvent) => {
        e.stopPropagation();
      },
    };
  }
  return {};
});
</script>
