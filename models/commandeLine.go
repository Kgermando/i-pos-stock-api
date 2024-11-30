package models

import "gorm.io/gorm"

type CommandeLine struct {
	gorm.Model

	CommandeID uint
	Commande   Commande `gorm:"foreignKey:CommandeID"`
	ProductID  uint
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   uint64   `gorm:"not null" json:"quantity"`
	PrixVente  float64  `gorm:"not null" json:"prix_vente"`
}
