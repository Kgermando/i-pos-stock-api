package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	CategoryID      uint
	Category        Category          `gorm:"foreignKey:CategoryID"`
	Reference       string            `gorm:"not null" json:"reference"`
	Name            string            `gorm:"not null" json:"name"`
	Description     string            `gorm:"not null" json:"description"`
	UniteVenteID    uint              // Par Bouteille, Sachet, Carton, Paquet, ...
	UniteVente      UniteVente        `gorm:"foreignKey:UniteVenteID"`
	Signature       string            `json:"signature"`
	BonCommadeLines []BonCommandeLine `gorm:"foreignKey:ProductID"`
	Stocks          []Stock           `gorm:"foreignKey:ProductID"`
}
