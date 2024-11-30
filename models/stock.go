package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	PosID          uint
	Pos            Pos `gorm:"foreignKey:PosID"`
	ProductID      uint
	Product        Product `gorm:"foreignKey:ProductID"`
	Description    string  `json:"description"`
	FournisseurID  uint
	Fournisseur    Fournisseur `gorm:"foreignKey:FournisseurID"`
	Quantity       uint64      `gorm:"not null" json:"quantity"`
	PrixAchat      float64     `gorm:"not null" json:"prix_achat"`
	PrixVente      float64     `gorm:"not null" json:"prix_vente"`
	DateExpiration time.Time   `gorm:"not null" json:"date_expiration"`
	Signature      string      `json:"signature"`
}
