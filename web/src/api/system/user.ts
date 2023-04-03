import rsa from '@/utils/rsa';
import axios from 'axios';

export function login(data: any) {
  const info = {
    password: rsa.encrypt({
      password: data.password,
      time: new Date().getTime(),
    }),
    phone: data.username,
    captcha: data.captcha,
    captcha_id: data.captcha_id,
  };
  return axios.post('/api/system/user/login', info);
}

export function captcha() {
  return axios.post('/api/system/captcha');
}

export function logout() {
  return axios.post('/api/system/user/logout');
}

export function refreshToken() {
  return axios.post('/api/system/token/refresh');
}

export function getUserMenus() {
  return axios.get('/api/system/user/menus');
}

export function getUsers(params?: any) {
  return axios.get('/api/system/users', { params });
}

export function getUser() {
  return axios.get('/api/system/user');
}

export function addUser(data: any) {
  return axios.post('/api/system/user', data);
}

export function updateUser(data: any) {
  return axios.put('/api/system/user', data);
}

export function deleteUser(data: any) {
  return axios.delete('/api/system/user', { data });
}

// 修改用户信息
export function updateUserInfo(data: any) {
  return axios.put('/api/system/user/info', data);
}

// 修改用户密码
export function updateUserPassword(data: any) {
  return axios.put('/api/system/user/password', data);
}

// 获取邮箱验证码
export function getEmailCaptcha() {
  return axios.post('/api/system/email/captcha');
}

export default null;
