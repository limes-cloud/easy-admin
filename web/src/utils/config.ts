import { getConfig } from '@/api/system/config';
import { AppState } from '@/store/modules/app/types';
import defaultSettings from '@/config/settings.json';
import { useAppStore } from '@/store';

export default async function systemConfig() {
  const { data } = await getConfig();
  const config = { ...defaultSettings, ...data };
  const app = useAppStore();
  app.setting(config as AppState);
  document.title = app.$state.title;
  return data;
}
