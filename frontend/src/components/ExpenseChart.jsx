import React from "react";
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";

const ExpenseChart = ({ transactions }) => {
  const data = [
    { type: "Revenus", total: transactions.filter(t => t.type === "income").reduce((sum, t) => sum + t.amount, 0) },
    { type: "Dépenses", total: transactions.filter(t => t.type === "expense").reduce((sum, t) => sum + t.amount, 0) },
  ];

  return (
    <div className="bg-white p-4 rounded-lg shadow-md">
      <h2 className="text-lg font-semibold mb-2">Aperçu des finances</h2>
      <ResponsiveContainer width="100%" height={250}>
        <BarChart data={data}>
          <XAxis dataKey="type" />
          <YAxis />
          <Tooltip />
          <Bar dataKey="total" fill="#4F46E5" />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default ExpenseChart;
