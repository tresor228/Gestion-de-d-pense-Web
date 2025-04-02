import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
    const isAuthenticated = !!localStorage.getItem("token"); // Exemple avec un token
    return isAuthenticated ? children : <Navigate to="/login" />;
};

export default ProtectedRoute;
