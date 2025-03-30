package models

import "gorm.io/gorm"

// User représente un utilisateur du système
type Utilisateur struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
