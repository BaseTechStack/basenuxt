import type { BaseItem } from '@@/app/stores/baseStore'
export interface entity extends BaseItem {
  test?: string;
  name?: string;
  email?: string;
  age?: number;
  active?: boolean;
  createdAt: string;
  updatedAt: string;
}
