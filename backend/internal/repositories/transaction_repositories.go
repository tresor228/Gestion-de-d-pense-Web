package repositories

import (
	"expense-tracker/internal/models"

	"gorm.io/gorm"
)

// TransactionRepository gère les opérations liées aux transactions.
type TransactionRepository struct {
	DB *gorm.DB
}

// NewTransactionRepository crée une nouvelle instance de TransactionRepository.
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// CreateTransaction ajoute une nouvelle transaction.
func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

// GetTransactionByID récupère une transaction par ID.
func (r *TransactionRepository) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetTransactionsByUserID récupère toutes les transactions d'un utilisateur.
func (r *TransactionRepository) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// UpdateTransaction met à jour une transaction existante.
func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

// DeleteTransaction supprime une transaction.
func (r *TransactionRepository) DeleteTransaction(id uint) error {
	return r.DB.Delete(&models.Transaction{}, id).Error
}
