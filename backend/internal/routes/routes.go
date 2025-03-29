package routes

import (
	"expense-tracker/internal/handlers"
	"expense-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initialise les routes
func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, transactionHandler *handlers.TransactionHandler) {
	api := app.Group("/api")

	// Routes publiques
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Routes protégées
	api.Use(middleware.AuthMiddleware)
	api.Post("/transaction", transactionHandler.AddTransaction)
}
