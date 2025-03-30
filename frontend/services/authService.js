import API from './api';

export const login = async (email, password) => {
  const response = await API.post('/auth/login', { email, password });
  if (response.data.token) {
    localStorage.setItem('token', response.data.token);
  }
  return response.data;
};

export const register = async (name, email, password) => {
  return await API.post('/auth/register', { name, email, password });
};

export const logout = () => {
  localStorage.removeItem('token');
};
