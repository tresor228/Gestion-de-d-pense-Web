import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem("token"); // Vérifie si l'utilisateur est connecté
  return token ? children : <Navigate to="/login" />;
};

export default ProtectedRoute;
