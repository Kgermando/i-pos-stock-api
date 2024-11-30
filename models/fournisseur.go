package models

import "gorm.io/gorm"

type Fournisseur struct {
	gorm.Model

	Name      string  `gorm:"not null" json:"name"` // Name of entreprise
	Adresse   string  `json:"adresse"`
	Email     string  `json:"email"`
	Telephone string  `json:"telephone"`
	Signature string  `json:"signature"`
	Stocks     []Stock `gorm:"foreignKey:FournisseurID"`
}
