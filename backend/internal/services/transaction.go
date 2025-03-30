package services

import (
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
)

// TransactionService gère la logique métier des transactions
type TransactionService struct {
	TransactionRepo *repositories.TransactionRepository
}

// NewTransactionService crée un nouveau service transaction
func NewTransactionService(transactionRepo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{TransactionRepo: transactionRepo}
}

// AddTransaction ajoute une transaction
func (s *TransactionService) AddTransaction(userID uint, montant float64, transType string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID:  userID,
		Montant: montant,
		Type:    transType,
	}

	err := s.TransactionRepo.CreateTransaction(transaction)
	return transaction, err
}
