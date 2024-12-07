package models

import "gorm.io/gorm"

type BonCommandeLine struct {
	gorm.Model

	BonCommandeID uint        `json:"bon_commande_id"`
	BonCommande   BonCommande `gorm:"foreignKey:BonCommandeID" json:"bon_commande"`
	ProductID     uint        `json:"product_id"`
	Product       Product     `gorm:"foreignKey:ProductID" json:"product"`
	Quantity      uint64      `gorm:"not null" json:"quantity"`
	PriceUnit     float64     `gorm:"not null" json:"price_unit"`
}
