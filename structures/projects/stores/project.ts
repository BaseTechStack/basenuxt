import type { BaseItem } from '@@/app/stores/baseStore'
export interface Project extends BaseItem {
  name?: string,description;
  /users/flakerimismani/base/basenuxt/testProject?: string;
  createdAt: string;
  updatedAt: string;
}
