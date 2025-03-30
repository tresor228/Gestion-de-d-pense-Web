package routes

import (
	"gestion-de-depense/backend/internal/handlers"
	"gestion-de-depense/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initialise les routes
func SetupRoutes(app *fiber.App, userHandler *handlers.Gestion_Utilisateur, transactionHandler *handlers.TransactionHandler) {
	api := app.Group("/api")

	// Routes publiques
	api.Post("/register", userHandler.Inscription)
	api.Post("/login", userHandler.Connexion)

	// Routes protégées
	api.Use(middleware.AuthMiddleware)
	api.Post("/transaction", transactionHandler.AddTransaction)
}
