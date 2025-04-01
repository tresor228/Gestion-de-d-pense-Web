import { useEffect, useState } from 'react';
import { getTransactions, addTransaction, deleteTransaction } from '../services/transactionService';

const Dashboard = () => {
  const [transactions, setTransactions] = useState([]);
  const [description, setDescription] = useState('');
  const [amount, setAmount] = useState('');

  useEffect(() => {
    loadTransactions();
  }, []);

  const loadTransactions = async () => {
    const data = await getTransactions();
    setTransactions(data);
  };

  const handleAddTransaction = async () => {
    await addTransaction({ description, amount });
    setDescription('');
    setAmount('');
    loadTransactions();
  };

  const handleDelete = async (id) => {
    await deleteTransaction(id);
    loadTransactions();
  };

  const handleDeleteTransaction = async (id) => {
    try {
      await transactionService.deleteTransaction(id);
      fetchTransactions(); // Recharge les transactions après suppression
    } catch (error) {
      console.error("Erreur lors de la suppression :", error);
    }
  };
  return (
    <div className="p-6">
      <h2 className="text-xl font-bold">Gestion des Dépenses</h2>

      <div className="my-4">
        <input type="text" placeholder="Description" value={description} onChange={(e) => setDescription(e.target.value)} className="border p-2 mr-2" />
        <input type="number" placeholder="Montant" value={amount} onChange={(e) => setAmount(e.target.value)} className="border p-2 mr-2" />
        <button onClick={handleAddTransaction} className="bg-green-500 text-white px-4 py-2">Ajouter</button>
      </div>

      <ul className="mt-4">
        {transactions.map((t) => (
          <li key={t.id} className="border p-2 flex justify-between">
            <span>{t.description} - {t.amount} FCFA</span>
            <button onClick={() => handleDelete(t.id)} className="bg-red-500 text-white px-2 py-1">Supprimer</button>
          </li>
        ))}
      </ul>
    </div>
  );
};
<ul className="mt-2">
  {transactions.map((transaction) => (
    <li
      key={transaction.id}
      className={`p-2 border-b flex justify-between items-center ${
        transaction.type === "income" ? "text-green-500" : "text-red-500"
      }`}
    >
      <span>{transaction.description} - {transaction.amount} FCFA</span>
      <button
        onClick={() => handleDeleteTransaction(transaction.id)}
        className="bg-red-500 text-white p-1 rounded hover:bg-red-700"
      >
        Supprimer
      </button>
    </li>
  ))}
</ul>

export default Dashboard;
