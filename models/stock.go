package models

import (
	"time"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	// ID             uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PosID          uint        `json:"pos_id"`
	Pos            Pos         `gorm:"foreignKey:PosID"`
	ProductID      uint        `json:"product_id"`
	Product        Product     `gorm:"foreignKey:ProductID"`
	Description    string      `json:"description"`
	Quantity       uint64      `gorm:"not null" json:"quantity"`
	FournisseurID  uint        `json:"fournisseur_id"`
	Fournisseur    Fournisseur `gorm:"foreignKey:FournisseurID"`
	PrixAchat      float64     `gorm:"not null" json:"prix_achat"`
	DateExpiration time.Time   `gorm:"not null" json:"date_expiration"`
	Signature      string      `json:"signature"`
	CodeEntreprise uint        `json:"code_entreprise"`
}
