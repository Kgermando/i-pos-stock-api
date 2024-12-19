package models

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type CommandeLine struct {
	gorm.Model
	
	// ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CommandeID     uint      `json:"commande_id"`
	Commande       Commande  `gorm:"foreignKey:CommandeID"`
	ProductID      uint      `json:"product_id"`
	Product        Product   `gorm:"foreignKey:ProductID"`
	Quantity       uint64    `gorm:"not null" json:"quantity"`
	CodeEntreprise uint      `json:"code_entreprise"`

	Mb float64 `gorm:"-" json:"mb"`
}
