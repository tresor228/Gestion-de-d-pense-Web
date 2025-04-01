import { useState } from "react";
import transactionService from "../services/transactionService";

const AddTransaction = ({ onTransactionAdded }) => {
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState("");
  const [type, setType] = useState("income");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!description || !amount) return;

    const newTransaction = { description, amount: parseFloat(amount), type };
    
    try {
      await transactionService.addTransaction(newTransaction);
      setDescription("");
      setAmount("");
      onTransactionAdded();
    } catch (error) {
      console.error("Erreur lors de l'ajout de la transaction :", error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white p-4 rounded shadow-md">
      <h3 className="text-xl font-semibold mb-2">Ajouter une transaction</h3>
      <input
        type="text"
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        className="w-full p-2 border rounded mb-2"
      />
      <input
        type="number"
        placeholder="Montant"
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        className="w-full p-2 border rounded mb-2"
      />
      <select
        value={type}
        onChange={(e) => setType(e.target.value)}
        className="w-full p-2 border rounded mb-2"
      >
        <option value="income">Revenu</option>
        <option value="expense">Dépense</option>
      </select>
      <button
        type="submit"
        className="w-full bg-green-500 text-white p-2 rounded hover:bg-green-700"
      >
        Ajouter
      </button>
    </form>
  );
};

export default AddTransaction;
