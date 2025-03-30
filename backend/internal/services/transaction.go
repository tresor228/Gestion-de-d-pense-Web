package services

import (
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
)

// structure pour le service de transaction
type Transaction_Service struct {
	TransactionRepo *repositories.Gestionnaire_Transaction
}

// Initialisation de Transaction Service
func Initialisation_Transaction_Service(transactionRepo *repositories.Gestionnaire_Transaction) *Transaction_Service {
	return &Transaction_Service{TransactionRepo: transactionRepo}
}

// Ajoutez une nouvelle transaction
func (s *Transaction_Service) Create_Transaction(userID uint, montant float64, transType string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID:  userID,
		Montant: montant,
		Type:    transType,
	}

	err := s.TransactionRepo.Create_Transaction(transaction)
	return transaction, err
}
