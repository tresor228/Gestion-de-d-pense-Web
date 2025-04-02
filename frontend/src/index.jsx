import { useEffect, useState } from "react";

function ExpensesList() {
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/expenses")
            .then(response => response.json())
            .then(data => setExpenses(data))
            .catch(error => console.error("Erreur :", error));
    }, []);

    return (
        <ul>
            {expenses.map(expense => (
                <li key={expense.id}>
                    {expense.name}: {expense.amount} €
                </li>
            ))}
        </ul>
    );
}

export default ExpensesList;