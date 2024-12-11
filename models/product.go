package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	CategoryID      uint              `json:"category_id"`
	Category        Category          `gorm:"foreignKey:CategoryID"`
	Reference       string            `gorm:"not null" json:"reference"`
	Name            string            `gorm:"not null" json:"name"`
	Description     string            `gorm:"not null" json:"description"`
	UniteVenteID    uint              `json:"unite_vente_id"` // Par Bouteille, Sachet, Carton, Paquet, ...
	UniteVente      UniteVente        `gorm:"foreignKey:UniteVenteID"`
	PrixVente       float64           `gorm:"not null" json:"prix_vente"`
	Signature       string            `json:"signature"`
	BonCommadeLines []BonCommandeLine `gorm:"foreignKey:ProductID"`
	Stocks          []Stock           `gorm:"foreignKey:ProductID"`
	PosID           uint              `json:"pos_id"`
	Pos             Pos               `gorm:"foreignKey:PosID"`
	CodeEntreprise  uint              `json:"code_entreprise"`
}
