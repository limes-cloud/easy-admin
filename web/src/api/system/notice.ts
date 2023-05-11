import axios from 'axios';

export function getNotices(params?: any) {
  return axios.get('/api/system/notices', { params });
}

export function getNotice(params?: any) {
  return axios.get('/api/system/notice', { params });
}

export function getUnreadNoticeNum() {
  return axios.get('/api/system/notice/unread_num');
}

export function addNotice(data: any) {
  return axios.post('/api/system/notice', data);
}

export function updateNotice(data: any) {
  return axios.put('/api/system/notice', data);
}

export function deleteNotice(data: any) {
  return axios.delete('/api/system/notice', { data });
}
