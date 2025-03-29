package handlers

import (
	"expense-tracker/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

// NewTransactionHandler crée un handler transaction
func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

// AddTransaction ajoute une transaction
func (h *TransactionHandler) AddTransaction(c *fiber.Ctx) error {
	type Request struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
		Type   string  `json:"type"` // "income" ou "expense"
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	transaction, err := h.TransactionService.AddTransaction(req.UserID, req.Amount, req.Type)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add transaction"})
	}

	return c.Status(201).JSON(transaction)
}
