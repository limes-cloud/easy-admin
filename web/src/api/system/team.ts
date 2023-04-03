import axios from 'axios';

export function getTeams(params?: any) {
  return axios.get('/api/system/teams', { params });
}

export function addTeam(data: any) {
  return axios.post('/api/system/team', data);
}

export function importTeam(data: any) {
  return axios.post('/api/system/teams', data);
}

export function updateTeam(data: any) {
  return axios.put('/api/system/team', data);
}

export function deleteTeam(data: any) {
  return axios.delete('/api/system/team', { data });
}
export default null;
