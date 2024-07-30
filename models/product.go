package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductID   string `gorm:"unique;not null"`
	ProductName string `gorm:"not null"`
	Category    string
}
