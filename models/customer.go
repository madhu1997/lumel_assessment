package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	CustomerID   string `gorm:"unique;not null"`
	CustomerName string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Address      string
}
