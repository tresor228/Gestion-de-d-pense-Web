import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL + "/transactions";

const getTransactions = async () => {
  const response = await axios.get(API_URL, {
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
  });
  return response.data;
};

const addTransaction = async (transaction) => {
  await axios.post(API_URL, transaction, {
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
  });
};

const deleteTransaction = async (id) => {
  await axios.delete(`${API_URL}/${id}`, {
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
  });
};

export default { getTransactions, addTransaction, deleteTransaction };