package routes

import (
	"gestion-de-depense/backend/internal/handlers"
	"gestion-de-depense/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initialise les routes
func SetupRoutes(app *fiber.App, userHandler *handlers.Gestion_Utilisateur, transactionHandler *handlers.Controleur_Transaction) {
	api := app.Group("/api")

	// Routes publiques
	api.Post("/Inscription", userHandler.Inscription)
	api.Post("/Connexion", userHandler.Connexion)

	// Routes protégées
	api.Use(middleware.AuthMiddleware)
	api.Post("/transaction", transactionHandler.Ajout_Transaction)
}
