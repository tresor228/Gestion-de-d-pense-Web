import { Line } from "react-chartjs-2";
import { Chart as ChartJS, LineElement, PointElement, LinearScale, Title, Tooltip } from "chart.js";

ChartJS.register(LineElement, PointElement, LinearScale, Title, Tooltip);

const Chart = ({ transactions }) => {
  const data = {
    labels: transactions.map((t) => t.date),
    datasets: [
      {
        label: "Dépenses et Revenus",
        data: transactions.map((t) => t.amount),
        borderColor: "#4CAF50",
        backgroundColor: "rgba(76, 175, 80, 0.2)",
      },
    ],
  };

  return <Line data={data} />;
};

export default Chart;