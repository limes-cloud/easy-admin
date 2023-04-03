import axios from 'axios';

export function getMenus(params?: any) {
  return axios.get('/api/system/menus', { params });
}

export function addMenu(data: any) {
  return axios.post('/api/system/menu', data);
}

export function importMenu(data: any) {
  return axios.post('/api/system/menus', data);
}

export function updateMenu(data: any) {
  return axios.put('/api/system/menu', data);
}

export function deleteMenu(data: any) {
  return axios.delete('/api/system/menu', { data });
}
export default null;
