import axios from 'axios';
import type { AxiosRequestConfig, AxiosResponse } from 'axios';
import { Message, Modal } from '@arco-design/web-vue';
import { useUserStore, useAppStore } from '@/store';
import { getToken, isLogin, setToken } from '@/utils/auth';
import { refreshToken } from '@/api/system/user';

export interface HttpResponse<T = unknown> {
  msg: string;
  code: number;
  data: T;
  total: number;
  page: number;
  page_size: number;
}

if (import.meta.env.VITE_API_BASE_URL) {
  axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
  axios.defaults.withCredentials = true;
  axios.defaults.timeout = 10000;
}

axios.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    if (isLogin()) {
      const token = getToken();
      if (!config.headers) {
        config.headers = {};
      }
      config.headers.Authorization = token as string;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 是否正在刷新的标记
let isRefresh = false;
// 重试队列，每一项将是一个待执行的函数形式
let requests: any = [];
// 是否正在退出登陆
let isLogout = false;

axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>) => {
    const res = response.data;
    if (res.code === useAppStore().$state.successCode) {
      return res;
    }

    // 处理需要重新登陆的错误
    if (res.code === 4000) {
      if (!isLogout) {
        isLogout = true;
        Modal.error({
          title: '重新登陆提醒',
          content: res.msg,
          okText: '重新登陆',
          async onOk() {
            const userStore = useUserStore();
            await userStore.logout();
            window.location.reload();
            isLogout = false;
          },
        });
      }
      return Promise.reject(new Error(res.msg || 'Error'));
    }

    // 重新登陆过期处理
    if (res.code === 4001) {
      const { config } = response;
      if (!isRefresh) {
        isRefresh = true;
        return refreshToken()
          .then((resToken) => {
            // 处理刷新成功
            setToken(resToken.data.token);
            requests.forEach((cb: any) => cb(resToken.data.token));
            requests = [];
            return axios(config);
          })
          .finally(() => {
            isRefresh = false;
          });
      }

      return new Promise((resolve) => {
        requests.push(() => {
          resolve(axios(config));
        });
      });
    }

    // 通用错误处理逻辑
    Message.error({
      content: res.msg || 'Error',
      duration: 5 * 1000,
    });
    return Promise.reject(new Error(res.msg || 'Error'));
  },
  (error) => {
    Message.error({
      content: error.msg || 'Request Error',
      duration: 5 * 1000,
    });
    return Promise.reject(error);
  }
);
