import axios from 'axios';

export function getConfig() {
  return axios.get('/api/system/config');
}
export default null;
