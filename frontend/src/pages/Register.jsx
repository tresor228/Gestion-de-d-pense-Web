import { useState } from "react";
import { useNavigate } from "react-router-dom";
import authService from "../services/authService";

const Register = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      await authService.register(name, email, password);
      navigate("/login");
    } catch (err) {
      setError("Échec de l'inscription. Essayez encore.");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-4 text-center">Inscription</h2>
        {error && <p className="text-red-500 text-sm">{error}</p>}
        <form onSubmit={handleRegister}>
          <input
            type="text"
            placeholder="Nom"
            className="w-full p-2 mb-2 border rounded"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <input
            type="email"
            placeholder="Email"
            className="w-full p-2 mb-2 border rounded"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            placeholder="Mot de passe"
            className="w-full p-2 mb-2 border rounded"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button className="w-full bg-green-500 text-white p-2 rounded hover:bg-green-600">
            S'inscrire
          </button>
        </form>
        <p className="text-sm mt-4 text-center">
          Déjà un compte ? <a href="/login" className="text-blue-500">Connectez-vous</a>
        </p>
      </div>
    </div>
  );
};

export default Register;
