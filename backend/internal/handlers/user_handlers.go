package handlers

import (
	"gestion-de-depense/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type Gestion_Utilisateur struct {
	Service_de_user *services.Service_Utilisateur
}

// Constructeur pour le gestionnaire d'utilisateur
func NewGestion_Utilisateur(userService *services.Service_Utilisateur) *Gestion_Utilisateur {
	return &Gestion_Utilisateur{Service_de_user: userService}
}

// Fonctionnnalité pour la gestion de l'inscription
func (h *Gestion_Utilisateur) Inscription(c *fiber.Ctx) error {
	type Requête struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Requête
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.Service_de_user.Inscription_Utilisateur(req.Email, req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}

	return c.Status(201).JSON(user)
}

// Fonctionnnalité pour la gestion de la connexion
func (h *Gestion_Utilisateur) Connexion(c *fiber.Ctx) error {
	type Requête struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Requête
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := h.Service_de_user.Connexion_utilisateur(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
