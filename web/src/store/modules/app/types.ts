import type { RouteRecordNormalized } from 'vue-router';

export interface AppState {
  title: string;
  desc: string;
  theme: string;
  colorWeak: boolean;
  navbar: boolean;
  menu: boolean;
  topMenu: boolean;
  hideMenu: boolean;
  menuCollapse: boolean;
  footer: boolean;
  themeColor: string;
  menuWidth: number;
  globalSettings: boolean;
  device: string;
  tabBar: boolean;
  menuFromServer: boolean;
  serverMenu: RouteRecordNormalized[];
  isLoading: boolean;
  permissions: string[];
  loadingTitle: string;
  debug: boolean;
  errorCode: number;
  successCode: number;
  copyright: string;
  staticUrl: string;
  logo: string;
  [key: string]: unknown;
}
