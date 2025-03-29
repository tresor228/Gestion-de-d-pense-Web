package handlers

import (
	"gestion-de-depense/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

// NewUserHandler crée un handler utilisateur
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Register gère l'inscription
func (h *UserHandler) Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.UserService.RegisterUser(req.Email, req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}

	return c.Status(201).JSON(user)
}

// Login gère la connexion
func (h *UserHandler) Login(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := h.UserService.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
