package models

import "gorm.io/gorm"

type Commande struct {
	gorm.Model

	PosID        uint
	Pos          Pos    `gorm:"foreignKey:PosID"`
	NCommande    uint64 `gorm:"not null" json:"n_commande"` // Number Random
	ClientID     uint
	Client       Client         `gorm:"foreignKey:ClientID"`
	Signature    string         `json:"signature"`
	CommandeLines []CommandeLine `gorm:"foreignKey:CommandeID"`
}
