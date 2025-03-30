package models

import "gorm.io/gorm"

// Transaction représente une transaction financière
type Transaction struct {
	gorm.Model
	UserID  uint    `gorm:"not null"`
	Montant float64 `gorm:"not null"`
	Type    string  `gorm:"not null"` // "Revenue" ou "Depense"
}
