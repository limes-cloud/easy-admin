import { createRouter, createWebHistory } from 'vue-router';
import NProgress from 'nprogress'; // progress bar
import 'nprogress/nprogress.css';

import createRouteGuard from './guard';
import { DEFAULT_LAYOUT } from './routes/base';

NProgress.configure({ showSpinner: false }); // NProgress Configuration

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/index.vue'),
      meta: {
        requiresAuth: false,
      },
    },

    {
      path: '/userinfo',
      name: 'Userinfo',
      component: DEFAULT_LAYOUT,
      redirect: '/userinfo/detail',
      meta: {
        requiresAuth: true,
      },
      children: [
        {
          path: 'detail',
          name: 'UserinfoDetail',
          component: () => import('@/views/system/user/info/index.vue'),
        },
      ],
    },
  ],
  scrollBehavior() {
    return { top: 0 };
  },
});

createRouteGuard(router);
export default router;
