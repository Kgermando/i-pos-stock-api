package models

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type Pos struct {
	gorm.Model

	// ID           uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	EntrepriseID uint          `json:"entreprise_id"`
	Entreprise   Entreprise    `gorm:"foreignKey:EntrepriseID"`
	Name         string        `gorm:"not null" json:"name"`
	Adresse      string        `json:"adresse"`
	Email        string        `json:"email"`
	Telephone    string        `json:"telephone"`
	Manager      string        `gorm:"not null" json:"manager"`
	Status       bool          `gorm:"not null" json:"status"` // Actif ou Inactif
	Signature    string        `json:"signature"`
	Stocks       []Stock       `gorm:"foreignKey:PosID" json:"stocks"`
	BonCommades  []BonCommande `gorm:"foreignKey:PosID" json:"bon_commandes"`
	Commandes    []Commande    `gorm:"foreignKey:PosID" json:"commandes"`
}
