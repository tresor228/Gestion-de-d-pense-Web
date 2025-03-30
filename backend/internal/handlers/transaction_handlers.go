package handlers

import (
	"gestion-de-depense/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type Controleur_Transaction struct {
	TransactionService *services.Transaction_Service
}

// NewTransactionHandler crée un handler transaction
func Initialisation_Gest_Transaction(transactionService *services.Transaction_Service) *Controleur_Transaction {
	return &Controleur_Transaction{TransactionService: transactionService}
}

// AddTransaction ajoute une transaction
func (h *Controleur_Transaction) Ajout_Transaction(c *fiber.Ctx) error {
	type Requete struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
		Type   string  `json:"type"` // "income" ou "expense"
	}

	var req Requete
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"erreur": "Requête invalide"})
	}

	transaction, err := h.TransactionService.Create_Transaction(req.UserID, req.Amount, req.Type)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"erreur": "Impossible de créer la transaction"})
	}

	return c.Status(201).JSON(transaction)
}
