export type WorkspacePermission =
  | "bb.environments.create"
  | "bb.environments.delete"
  | "bb.environments.get"
  | "bb.environments.list"
  | "bb.environments.undelete"
  | "bb.environments.update"
  | "bb.externalVersionControls.create"
  | "bb.externalVersionControls.delete"
  | "bb.externalVersionControls.get"
  | "bb.externalVersionControls.list"
  | "bb.externalVersionControls.listProjects"
  | "bb.externalVersionControls.searchProjects"
  | "bb.externalVersionControls.update"
  | "bb.identityProviders.create"
  | "bb.identityProviders.delete"
  | "bb.identityProviders.get"
  | "bb.identityProviders.undelete"
  | "bb.identityProviders.update"
  | "bb.instanceRoles.create"
  | "bb.instanceRoles.delete"
  | "bb.instanceRoles.get"
  | "bb.instanceRoles.list"
  | "bb.instanceRoles.undelete"
  | "bb.instanceRoles.update"
  | "bb.instances.create"
  | "bb.instances.delete"
  | "bb.instances.get"
  | "bb.instances.list"
  | "bb.instances.sync"
  | "bb.instances.undelete"
  | "bb.instances.update"
  | "bb.policies.create"
  | "bb.policies.delete"
  | "bb.policies.get"
  | "bb.policies.list"
  | "bb.policies.update"
  | "bb.projects.create"
  | "bb.projects.delete"
  | "bb.projects.list"
  | "bb.projects.undelete"
  | "bb.risks.create"
  | "bb.risks.delete"
  | "bb.risks.list"
  | "bb.risks.update"
  | "bb.roles.create"
  | "bb.roles.delete"
  | "bb.roles.list"
  | "bb.roles.update"
  | "bb.settings.get"
  | "bb.settings.list"
  | "bb.settings.set";

export type ProjectPermission =
  | "bb.backups.create"
  | "bb.backups.list"
  | "bb.branches.create"
  | "bb.branches.delete"
  | "bb.branches.get"
  | "bb.branches.list"
  | "bb.branches.update"
  | "bb.changeHistories.get"
  | "bb.changeHistories.list"
  | "bb.changelists.create"
  | "bb.changelists.delete"
  | "bb.changelists.get"
  | "bb.changelists.list"
  | "bb.changelists.update"
  | "bb.databaseSecrets.delete"
  | "bb.databaseSecrets.list"
  | "bb.databaseSecrets.update"
  | "bb.databases.adviseIndex"
  | "bb.databases.export"
  | "bb.databases.get"
  | "bb.databases.getBackupSetting"
  | "bb.databases.getSchema"
  | "bb.databases.list"
  | "bb.databases.query"
  | "bb.databases.sync"
  | "bb.databases.update"
  | "bb.databases.updateBackupSetting"
  | "bb.issueComments.create"
  | "bb.issueComments.update"
  | "bb.issues.create"
  | "bb.issues.get"
  | "bb.issues.list"
  | "bb.issues.update"
  | "bb.planCheckRuns.list"
  | "bb.planCheckRuns.run"
  | "bb.plans.create"
  | "bb.plans.get"
  | "bb.plans.list"
  | "bb.plans.update"
  | "bb.projects.get"
  | "bb.projects.getIamPolicy"
  | "bb.projects.setIamPolicy"
  | "bb.projects.update"
  | "bb.rollouts.create"
  | "bb.rollouts.get"
  | "bb.rollouts.preview"
  | "bb.slowQueries.list"
  | "bb.taskRuns.list"
  | "bb.tasks.run"
  | "bb.tasks.skip";

export type Permission = WorkspacePermission | ProjectPermission;
