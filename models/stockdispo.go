package models

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type StockDispo struct {
	gorm.Model

	// ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PosID          uint      `json:"pos_id"`
	Pos            Pos       `gorm:"foreignKey:PosID"`
	ProductID      uint      `json:"product_id"`
	Product        Product   `gorm:"foreignKey:ProductID"`
	QtyStock       uint64    `gorm:"not null" json:"qty_stock"`
	QtyCmdline     uint64    `gorm:"not null" json:"qty_cmdline"`
	PriceStock     float64   `gorm:"not null" json:"price_stock"`
	CodeEntreprise uint      `json:"code_entreprise"`

	Dispo uint64 `gorm:"-" json:"dispo"`
	Pourcent float64 `gorm:"-" json:"pourcent"`
}
