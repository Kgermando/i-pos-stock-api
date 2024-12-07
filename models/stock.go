package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	PosID          uint        `json:"pos_id"`
	Pos            Pos         `gorm:"foreignKey:PosID"`
	ProductID      uint        `json:"product_id"`
	Product        Product     `gorm:"foreignKey:ProductID"`
	Description    string      `json:"description"`
	FournisseurID  uint        `json:"fournisseur_id"`
	Fournisseur    Fournisseur `gorm:"foreignKey:FournisseurID"`
	Quantity       uint64      `gorm:"not null" json:"quantity"`
	PrixAchat      float64     `gorm:"not null" json:"prix_achat"`
	PrixVente      float64     `gorm:"not null" json:"prix_vente"`
	DateExpiration time.Time   `gorm:"not null" json:"date_expiration"`
	Signature      string      `json:"signature"`
	CodeEntreprise uint        `json:"code_entreprise"`
}
