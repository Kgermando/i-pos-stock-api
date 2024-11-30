package models

import "gorm.io/gorm"

type UniteVente struct {
	gorm.Model

	Name      string    `gorm:"not null" json:"name"`
	Signature string    `json:"signature"`
	Products  []Product `gorm:"foreignKey:UniteVenteID"`
}
