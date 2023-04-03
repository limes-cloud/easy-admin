import { defineStore } from 'pinia';
import {
  login as userLogin,
  logout as userLogout,
  getUser,
} from '@/api/system/user';
import { setToken, clearToken } from '@/utils/auth';
import { removeRouteListener } from '@/utils/route-listener';
import { UserState } from './types';
import useAppStore from '../app';

const useUserStore = defineStore('user', {
  state: (): UserState => ({
    id: undefined,
    team_id: undefined,
    team: undefined,
    role_id: undefined,
    role: undefined,
    name: undefined,
    sex: undefined,
    phone: undefined,
    nickname: undefined,
    avatar: undefined,
    email: undefined,
    status: undefined,
    last_login: undefined,
    operator: undefined,
    operator_id: undefined,
  }),

  getters: {
    userInfo(state: UserState): UserState {
      return { ...state };
    },
  },

  actions: {
    // switchRoles() {
    //   return new Promise((resolve) => {
    //     this.role = this.role === 'user' ? 'admin' : 'user';
    //     resolve(this.role);
    //   });
    // },
    // Set user's information
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial);
    },

    // Reset user's information
    resetInfo() {
      this.$reset();
    },

    // Get user's information
    async info() {
      const res = await getUser();

      this.setInfo(res.data);
    },

    // Login
    async login(loginForm: any) {
      try {
        const res = await userLogin(loginForm);
        setToken(res.data.token);
      } catch (err) {
        clearToken();
        throw err;
      }
    },
    logoutCallBack() {
      const appStore = useAppStore();
      this.resetInfo();
      clearToken();
      removeRouteListener();
      appStore.clearServerMenu();
    },
    // Logout
    async logout() {
      try {
        await userLogout();
      } finally {
        this.logoutCallBack();
      }
    },
  },
});

export default useUserStore;
