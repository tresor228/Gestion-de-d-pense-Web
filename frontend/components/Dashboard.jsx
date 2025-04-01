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

export default Dashboard;
