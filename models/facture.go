package models

import "gorm.io/gorm"

type Facture struct {
	gorm.Model

	NFacture        float64  `gorm:"not null" json:"n_facture"`
	CommandeID      uint     `json:"commande_id"`
	Commande        Commande `gorm:"foreignKey:CommandeID"`
	Status          string   `gorm:"not null" json:"status"` // Cash ou Creance
	DelaiPaiement   string   `json:"delai_paiement"`
	Signature       string   `json:"signature"`
	PosID           uint     `json:"pos_id"`
	Pos             Pos      `gorm:"foreignKey:PosID"`
	CodeEentreprise float64  `json:"code_entreprise"`
}
