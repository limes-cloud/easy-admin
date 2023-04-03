import axios from 'axios';

export function getRoles(params?: any) {
  return axios.get('/api/system/roles', { params });
}

export function getRole(params: any) {
  return axios.get('/api/system/role', { params });
}

export function addRole(data: any) {
  return axios.post('/api/system/role', data);
}

export function updateRole(data: any) {
  return axios.put('/api/system/role', data);
}

export function deleteRole(data: any) {
  return axios.delete('/api/system/role', { data });
}

export function getRoleMenu(params: any) {
  return axios.get('/api/system/role/menu', { params });
}

export function updateRoleMenu(data: any) {
  return axios.put('/api/system/role/menu', data);
}

export function getRoleMenuids(params: any) {
  return axios.get('/api/system/role/menu/ids', { params });
}

export default null;
