package models

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type Commande struct {
	gorm.Model
	
	// ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PosID          uint           `json:"pos_id"`
	Pos            Pos            `gorm:"foreignKey:PosID"`
	Ncommande      uint64         `gorm:"not null" json:"ncommande"` // Number Random
	Status         string         `json:"status"`                     // Ouverte et Fermée
	ClientID       uint           `json:"client_id"`
	Client         Client         `gorm:"foreignKey:ClientID"`
	Signature      string         `json:"signature"`
	CodeEntreprise uint           `json:"code_entreprise"`
	CommandeLines  []CommandeLine `gorm:"foreignKey:CommandeID"`
}
