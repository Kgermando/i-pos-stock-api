package models

import (
	"time"

	"gorm.io/gorm"
)

type BonCommande struct {
	gorm.Model

	PosID           uint              `gorm:"index" json:"pos_id"`
	Pos             Pos               `gorm:"foreignKey:PosID"`
	NCommande       uint64            `gorm:"not null" json:"n_commande"` // Number Random
	DateCommande    time.Time         `gorm:"not null" json:"date_commande"`
	DateLivraison   time.Time         `gorm:"not null" json:"date_livraison"`
	FournisseurID   uint              `json:"fournisseur_id"`
	Fournisseur     Fournisseur       `gorm:"foreignKey:FournisseurID"`
	Status          string            `gorm:"not null" json:"status"` // brouillon, valide, annule
	MontantTotal    float64           `gorm:"not null" json:"montant_total"`
	Notes           string            `json:"notes"`
	BonCommadeLines []BonCommandeLine `gorm:"foreignKey:BonCommandeID"`
	Signature       string            `json:"signature"`
}
