package models

import "gorm.io/gorm"

type BonCommandeLine struct {
	gorm.Model

	BonCommandeID uint
	BonCommande   BonCommande `gorm:"foreignKey:BonCommandeID"`
	ProductID     uint
	Product       Product `gorm:"foreignKey:ProductID"`
	Quantity      uint64  `gorm:"not null" json:"quantity"`
	PriceUnit     float64 `gorm:"not null" json:"price_unit"`
}
