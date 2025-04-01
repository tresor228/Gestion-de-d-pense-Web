package repositories

import (
	"gestion-de-depense/backend/internal/models"

	"gorm.io/gorm"
)

// Struture pour le dépôt de transaction
type Gestionnaire_Transaction struct {
	DB *gorm.DB
}

// Nouvelle instance de Gestionnaire de Transaction
func Initialisation_Gestionnaire_Transaction(db *gorm.DB) *Gestionnaire_Transaction {
	return &Gestionnaire_Transaction{DB: db}
}

// Ajout d'une transaction
func (r *Gestionnaire_Transaction) Create_Transaction(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

// Récupération d'une transaction par Mail
func (r *Gestionnaire_Transaction) Recuperation_Trans_Mail(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Recuperation Mail par ID
func (r *Gestionnaire_Transaction) Recuperation_Trans_ID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Mise à jour d'une transaction
func (r *Gestionnaire_Transaction) Mise_a_jour_Transaction(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

// Suppression d'une transaction
func (r *Gestionnaire_Transaction) Suppression_Transaction(id string) error {
	return r.DB.Delete(&models.Transaction{}, id).Error
}
