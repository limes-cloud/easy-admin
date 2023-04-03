import axios from 'axios';

export function getLoginLog(params?: any) {
  return axios.get('/api/system/login/log', { params });
}

export default null;
