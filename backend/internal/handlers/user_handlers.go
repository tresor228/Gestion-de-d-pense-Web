package handlers

import (
	"gestion-de-depense/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type Gestion_Utilisateur struct {
	Service_de_user *services.Service_Utilisateur
}

// NewUserHandler crée un handler utilisateur
func NewGestion_Utilisateur(userService *services.Service_Utilisateur) *Gestion_Utilisateur {
	return &Gestion_Utilisateur{Service_de_user: userService}
}

// Register gère l'inscription
func (h *Gestion_Utilisateur) Inscription(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.Service_de_user.Inscription_User(req.Email, req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}

	return c.Status(201).JSON(user)
}

// Login gère la connexion
func (h *Gestion_Utilisateur) Connexion(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := h.Service_de_user.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
