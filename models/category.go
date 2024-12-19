package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name      string    `gorm:"not null" json:"category"`
	Signature string    `json:"signature"`
	// Products  []Product `gorm:"foreignKey:CategoryID"`
}
