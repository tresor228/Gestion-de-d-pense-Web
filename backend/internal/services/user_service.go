package services

import (
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
	"gestion-de-depense/backend/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

// logique des utilisateurs
type Service_Utilisateur struct {
	UserRepo *repositories.Depot_Utilisateur
}

// NewUserService crée une nouvelle instance
func Initilisation_Service_Utilisteur(userRepo *repositories.Depot_Utilisateur) *Service_Utilisateur {
	return &Service_Utilisateur{UserRepo: userRepo}
}

// Inscription d'un utilisateur
func (s *Service_Utilisateur) Inscription_Utilisateur(email, password string) (*models.Utilisateur, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Utilisateur{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.UserRepo.Ajout_Utilisateur(user)
	return user, err
}

// Connexion Utilisateur authentifie l'utilisateur et génère un jeton JWT
func (s *Service_Utilisateur) Connexion_utilisateur(email, password string) (string, error) {
	user, err := s.UserRepo.Recuperation_user_par_mail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	// Création du token JWT
	tokenString, err := auth.GenerateToken(user.ID)

	return tokenString, nil
}
