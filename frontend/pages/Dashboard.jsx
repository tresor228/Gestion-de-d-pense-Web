import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import transactionService from "../services/transactionService";
import Chart from "../components/Chart";
import AddTransaction from "../components/AddTransaction";

const Dashboard = () => {
  const [transactions, setTransactions] = useState([]);
  const [balance, setBalance] = useState(0);
  const navigate = useNavigate();

  const fetchTransactions = async () => {
    try {
      const data = await transactionService.getTransactions();
      setTransactions(data);
      calculateBalance(data);
    } catch (error) {
      console.error("Erreur lors du chargement des transactions :", error);
    }
  };

  useEffect(() => {
    fetchTransactions();
  }, []);

  const calculateBalance = (transactions) => {
    const total = transactions.reduce((acc, transaction) => {
      return transaction.type === "income" ? acc + transaction.amount : acc - transaction.amount;
    }, 0);
    setBalance(total);
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <h1 className="text-3xl font-bold text-center">Tableau de bord</h1>
      <div className="max-w-4xl mx-auto bg-white shadow-md rounded-lg p-6 mt-4">
        <h2 className="text-2xl font-semibold text-gray-700">Solde: {balance} FCFA</h2>
        <Chart transactions={transactions} />
        <AddTransaction onTransactionAdded={fetchTransactions} />
        <h3 className="text-xl font-semibold mt-4">Historique des transactions</h3>
        <ul className="mt-2">
          {transactions.map((transaction) => (
            <li
              key={transaction.id}
              className={`p-2 border-b ${
                transaction.type === "income" ? "text-green-500" : "text-red-500"
              }`}
            >
              {transaction.description} - {transaction.amount} FCFA
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Dashboard;
