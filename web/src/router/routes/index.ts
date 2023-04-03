import type { RouteRecordNormalized } from 'vue-router';
import { DEFAULT_LAYOUT } from './base';

const modules = {
  ...import.meta.glob('@/views/*/*.vue'),
  ...import.meta.glob('@/views/*/*/*.vue'),
  ...import.meta.glob('@/views/*/*/*/*.vue'),
};

let homePath = '';
let index = 0;

async function appRouteAndPermission(
  routes: RouteRecordNormalized[],
  menus: any[],
  permissions: string[]
) {
  menus.forEach((item) => {
    if (item.permission === 'baseApi' || item.name === 'baseApi') {
      return;
    }
    // 处理菜单
    let component: any;
    if (item.component && item.component !== 'Layout') {
      component = modules[`/src/views/${item.component}.vue`]();
    }
    const route: any = {
      path: item.path,
      name: item.name,
      component: item.component === 'Layout' ? DEFAULT_LAYOUT : () => component,
      redirect: item.redirect,
      activeMenu: item.name,
      children: [],
      meta: {
        requiresAuth: true,
        title: item.title,
        icon: `icon-${item.icon}`,
        // hideChildrenInMenu: !item.children || item.children.length === 1,
        hideInMenu: item.is_hidden,
        order: -item.weight,
        ignoreCache: !item.is_cache,
      },
    };

    if (item.type === 'M') {
      if (index === 0) {
        homePath = item.path;
        index = 1;
      }
      if (item.is_home) homePath = item.path;
      routes.push(route as RouteRecordNormalized);
    }

    // 添加指令权限
    if (item.permission) {
      permissions.push(item.permission);
    }

    // 处理子菜单
    if (item.children && !item.is_hidden) {
      appRouteAndPermission(
        route.children as RouteRecordNormalized[],
        item.children,
        permissions
      );
    }
  });
}

export function getAppRouteAndPermission(menu: any) {
  const routes: RouteRecordNormalized[] = [];
  const permissions: string[] = [];
  appRouteAndPermission(routes, menu.children, permissions);
  // const formatRoutes = formatModules(routes, []);
  return {
    routes,
    permissions,
    homePath,
  };
}

// function formatModules(_modules: any, result: RouteRecordNormalized[]) {
//   Object.keys(_modules).forEach((key) => {
//     const defaultModule = _modules[key].default;
//     if (!defaultModule) return;
//     const moduleList = Array.isArray(defaultModule)
//       ? [...defaultModule]
//       : [defaultModule];
//     result.push(...moduleList);
//   });
//   return result;
// }

export const appRoutes: RouteRecordNormalized[] = [] as RouteRecordNormalized[];

// export const appRoutes: RouteRecordNormalized[] = formatModules(modules, []);

export default null;
