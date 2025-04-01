import API from './api';

export const getTransactions = async () => {
  const response = await API.get('/transactions');
  return response.data;
};

export const addTransaction = async (transaction) => {
  return await API.post('/transactions', transaction);
};

export const deleteTransaction = async (id) => {
  return await API.delete(`/transactions/${id}`);
};
