export type RoleType = '' | '*' | 'admin' | 'user';
export interface UserState {
  id?: number;
  team_id?: number;
  team?: any;
  role_id?: number;
  role?: any;
  name?: string;
  sex?: boolean;
  phone?: string;
  nickname?: string;
  avatar?: string;
  email?: string;
  status?: boolean;
  last_login?: number;
  operator?: string;
  operator_id?: number;
}
