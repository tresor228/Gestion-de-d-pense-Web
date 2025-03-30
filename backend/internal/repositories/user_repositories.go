package repositories

import (
	"gestion-de-depense/backend/internal/models"

	"gorm.io/gorm"
)

// Structeure pour le dépôt d'utilisateur
type Depot_Utilisateur struct {
	DB *gorm.DB
}

// Nouvelle instance de Depot Utilisateur
func Initialisation_Depot_Utilisateur(db *gorm.DB) *Depot_Utilisateur {
	return &Depot_Utilisateur{DB: db}
}

// Ajout d'un Utilisateur
func (r *Depot_Utilisateur) Ajout_Utilisateur(user *models.Utilisateur) error {
	return r.DB.Create(user).Error
}

// Recherche un Utilisateur par email
func (r *Depot_Utilisateur) Recuperation_user_par_mail(email string) (*models.Utilisateur, error) {
	var user models.Utilisateur
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Recherche un Utilisateur par ID
func (r *Depot_Utilisateur) Recuperation_user_par_ID(id uint) (*models.Utilisateur, error) {
	var user models.Utilisateur
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Mise a jour d'un Utilisateur
func (r *Depot_Utilisateur) Mise_a_jour_Utilisateur(user *models.Utilisateur) error {
	return r.DB.Save(user).Error
}

// Supprimer un Utilisateur
func (r *Depot_Utilisateur) Supprimer_un_Utilisateur(id uint) error {
	return r.DB.Delete(&models.Utilisateur{}, id).Error
}
