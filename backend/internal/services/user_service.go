package services

import (
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
	"gestion-de-depense/backend/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

// UserService gère la logique utilisateur
type UserService struct {
	UserRepo *repositories.UserRepository
}

// NewUserService crée une nouvelle instance
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// RegisterUser enregistre un utilisateur
func (s *UserService) RegisterUser(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.UserRepo.CreateUser(user)
	return user, err
}

// LoginUser vérifie l'utilisateur et génère un token JWT
func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
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
