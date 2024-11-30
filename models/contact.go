package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model

	Fullname string `gorm:"not null" json:"fullname"`
	Email    string `gorm:"not null" json:"email"`
	Subject  string `gorm:"not null" json:"subject"`
	Message  string `gorm:"not null" json:"message"`
}
