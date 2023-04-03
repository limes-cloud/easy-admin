import Mock from 'mockjs';
import setupMock, {
  successResponseWrap,
  failResponseWrap,
} from '@/utils/setup-mock';

import { MockParams } from '@/types/mock';
import { isLogin } from '@/utils/auth';

setupMock({
  setup() {
    // Mock.XHR.prototype.withCredentials = true;

    // 登录
    Mock.mock(new RegExp('/api/system/user/login1'), (params: MockParams) => {
      console.log('/api/system/user/login');

      const { username, password } = JSON.parse(params.body);
      if (!username) {
        return failResponseWrap(null, '用户名不能为空', 50000);
      }
      if (!password) {
        return failResponseWrap(null, '密码不能为空', 50000);
      }
      if (username === 'admin' && password === 'admin') {
        window.localStorage.setItem('userRole', 'admin');
        return successResponseWrap({
          token: '12345',
          refresh_token: '123456',
        });
      }
      if (username === 'user' && password === 'user') {
        window.localStorage.setItem('userRole', 'user');
        return successResponseWrap({
          token: '54321',
          refresh_token: '123456',
        });
      }
      return failResponseWrap(null, '账号或者密码错误', 50000);
    });

    // 用户信息
    Mock.mock(new RegExp('/api/system/user/info'), () => {
      if (isLogin()) {
        return successResponseWrap({
          location: 'beijing',
          locationName: '北京',
          id: 1,
          team_id: 1,
          team: {
            name: '开发部',
            parent_id: 0,
            operator: 'system',
            operator_id: 1,
          },
          role_id: 1,
          role: {
            name: '超级管理员',
            parent_id: 0,
            operator: 'system',
            operator_id: 1,
          },
          name: '王立群',
          sex: true,
          phone: '18286219254',
          nickname: '柠檬很酸',
          avatar: 'logo.png',
          email: '1280291001@qq.com',
          status: true,
          last_login: 0,
          operator: 'system',
          operator_id: 0,
        });
      }
      return failResponseWrap(null, '未登录', 50008);
    });

    // 登出
    Mock.mock(new RegExp('/api/system/user/logout'), () => {
      return successResponseWrap(null);
    });

    // 用户的服务端菜单
    Mock.mock(new RegExp('/mock/api/system/user/menu'), () => {
      const menuList = {
        path: '/',
        name: 'P',
        component: '',
        redirect: false,
        hidden: false,
        type: '',
        permission: '',
        icon: 'icon-dashboard',
        order: 1,
        title: '菜单节点',
        weight: 10,
        children: [
          {
            path: '/dashboard',
            name: 'dashboard',
            component: 'Layout',
            redirect: '/dashboard/workplace',
            is_hidden: false,
            type: 'M',
            permission: '',
            icon: 'icon-dashboard',
            order: 1,
            is_home: true,
            title: '服务端',
            weight: 10,
            children: [
              {
                path: 'workplace',
                name: 'Workplace',
                component: 'dashboard/workplace/index',
                redirect: '',
                is_hidden: false,
                type: 'M',
                permission: '',
                title: '首页面板',
                icon: 'icon-dashboard',
                weight: 10,
              },
            ],
          },
          {
            path: '/system',
            name: 'System',
            component: 'Layout',
            redirect: '/system/user',
            is_hidden: false,
            type: 'M',
            permission: '',
            icon: 'icon-computer',
            order: 1,
            title: '系统管理',
            weight: 10,
            children: [
              {
                path: 'menu',
                name: 'Menu',
                component: 'system/menu/index',
                redirect: false,
                is_hidden: false,
                type: 'M',
                permission: '',
                icon: 'icon-menu',
                weight: 10,
                title: '菜单管理',
              },
              {
                path: 'role',
                name: 'Role',
                component: 'system/role/index',
                redirect: false,
                is_hidden: false,
                type: 'M',
                permission: '',
                icon: 'icon-relation',
                weight: 10,
                title: '角色管理',
              },
              {
                path: 'team',
                name: 'Team',
                component: 'system/team/index',
                redirect: false,
                is_hidden: false,
                type: 'M',
                permission: '',
                icon: 'icon-user-group',
                weight: 10,
                title: '部门管理',
              },
              {
                path: 'user',
                name: 'User',
                component: 'system/user/index',
                redirect: false,
                is_hidden: false,
                type: 'M',
                permission: '',
                icon: 'icon-user',
                weight: 10,
                title: '用户管理',
              },
              {
                path: 'login_log',
                name: 'LoginLog',
                component: 'system/login_log/index',
                redirect: false,
                is_hidden: false,
                type: 'M',
                permission: '',
                icon: 'icon-find-replace',
                weight: 10,
                title: '登陆日志',
              },
            ],
          },
        ],
      };
      return successResponseWrap(menuList);
    });
  },
});
