import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL + "/auth";

const login = async (email, password) => {
  const response = await axios.post(`${API_URL}/login`, { email, password });
  return response.data;
};

const register = async (name, email, password) => {
  const response = await axios.post(`${API_URL}/register`, { name, email, password });
  return response.data;
};

export default { login, register };
