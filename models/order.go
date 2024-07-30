package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID       int       `gorm:"unique;not null"`
	CustomerID    string    `gorm:"not null"`
	RegionID      int       `gorm:"not null"`
	DateOfSale    time.Time `gorm:"not null"`
	PaymentMethod string    `gorm:"not null"`
	ShippingCost  float64   `gorm:"not null"`
	Customer      Customer  `gorm:"foreignKey:CustomerID"`
	Region        Region    `gorm:"foreignKey:RegionID"`
}
