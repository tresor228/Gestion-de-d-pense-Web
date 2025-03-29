package handlers

import (
	"expense-tracker/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/register", func(c *fiber.Ctx) error {
		return services.RegisterUser(c, db)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return services.LoginUser(c, db)
	})

	app.Use(AuthMiddleware)

	app.Get("/transactions", func(c *fiber.Ctx) error {
		return services.GetTransactions(c, db)
	})

	app.Post("/transactions", func(c *fiber.Ctx) error {
		return services.AddTransaction(c, db)
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	return services.AuthMiddleware(c)
}
