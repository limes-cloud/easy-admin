const TOKEN_KEY = 'token';

const isLogin = () => {
  return !!localStorage.getItem(TOKEN_KEY);
};

const getToken = () => {
  return localStorage.getItem(TOKEN_KEY);
};

const setToken = (data: string) => {
  localStorage.setItem(TOKEN_KEY, data);
};

const clearToken = () => {
  localStorage.removeItem(TOKEN_KEY);
};

export { isLogin, getToken, setToken, clearToken };
