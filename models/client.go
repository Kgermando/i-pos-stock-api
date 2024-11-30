package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	FullName  string     `gorm:"not null" json:"fullname"`
	Telephone string     `gorm:"not null" json:"telephone"`
	Email     string     `json:"email"`
	Signature string     `json:"signature"`
	Commandes []Commande `gorm:"foreignKey:ClientID"`
}

