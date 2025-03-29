package repositories

import (
	"gestion-de-depense/backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository définit les méthodes d'accès aux données utilisateur.
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository crée une nouvelle instance de UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser ajoute un nouvel utilisateur dans la base de données.
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetUserByEmail recherche un utilisateur par email.
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID recherche un utilisateur par ID.
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser met à jour un utilisateur existant.
func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

// DeleteUser supprime un utilisateur.
func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
