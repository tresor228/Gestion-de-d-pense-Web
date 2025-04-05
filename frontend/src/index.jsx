import { useEffect, useState } from "react";
import axios from "axios";

function ExpensesList() {
    const [expenses, setExpenses] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        // Utilisation d'axios au lieu de fetch pour la cohérence avec le reste du projet
        setIsLoading(true);
        axios.get("/api/transactions") // Utilise le proxy configuré dans package.json
            .then(response => {
                setExpenses(response.data);
                setIsLoading(false);
            })
            .catch(error => {
                console.error("Erreur lors du chargement des dépenses:", error);
                setError("Impossible de charger les dépenses");
                setIsLoading(false);
            });
    }, []);

    if (isLoading) return <p>Chargement des dépenses...</p>;
    if (error) return <p className="text-red-500">{error}</p>;
    if (expenses.length === 0) return <p>Aucune dépense enregistrée.</p>;

    return (
        <div className="mt-4">
            <h2 className="text-xl font-semibold mb-2">Liste des dépenses</h2>
            <ul className="divide-y divide-gray-200">
                {expenses.map(expense => (
                    <li key={expense.id} className="py-2 flex justify-between">
                        <span>{expense.description || expense.name}</span>
                        <span className={expense.amount < 0 ? "text-red-500" : "text-green-500"}>
                            {expense.amount.toFixed(2)} €
                        </span>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default ExpensesList;