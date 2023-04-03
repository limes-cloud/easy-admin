import type { Router } from 'vue-router';
import NProgress from 'nprogress'; // progress bar

import { useAppStore } from '@/store';
import { getUserMenus } from '@/api/system/user';
import { getAppRouteAndPermission } from '../routes';
import { REDIRECT_MAIN, NOT_FOUND_ROUTE } from '../routes/base';
import { WHITE_LIST } from '../constants';

export default function setupPermissionGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const appStore = useAppStore();
    // const userStore = useUserStore();
    if (appStore.menuFromServer) {
      // 判读是否存在菜单
      if (
        !appStore.appAsyncMenus.length &&
        !WHITE_LIST.find((el) => el.name === to.name)
      ) {
        appStore.startLoading('系统初始化中');
        const { data } = await getUserMenus();
        const info = getAppRouteAndPermission(data);
        info.routes.forEach((item) => {
          router.addRoute(item);
        });
        // 添加到路由中
        router.addRoute(REDIRECT_MAIN);
        router.addRoute(NOT_FOUND_ROUTE);

        appStore.setServerMenu(info);
        appStore.stopLoading();
        // 路由跳转
        if (to.path === '/') {
          if (info.homePath) {
            next({ path: info.homePath, replace: true });
          } else {
            next({ path: info.routes[0].path, replace: true });
          }
        } else {
          next({ ...to, replace: true });
        }
      } else {
        next();
      }
    } else {
      next();
    }
    NProgress.done();
  });
}
