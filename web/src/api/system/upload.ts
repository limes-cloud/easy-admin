import axios from 'axios';

export function uploadFile(data: any) {
  return axios.post('/api/system/upload', data);
}

export default null;
