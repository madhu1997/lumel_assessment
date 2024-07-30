package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID      int     `gorm:"unique;not null"`
	ProductID    string  `gorm:"not null"`
	QuantitySold int     `gorm:"not null"`
	UnitPrice    float64 `gorm:"not null"`
	Discount     float64
	Order        Order   `gorm:"foreignKey:OrderID"`
	Product      Product `gorm:"foreignKey:ProductID;references:ProductID"`
}
