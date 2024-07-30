package models

import "github.com/jinzhu/gorm"

type Region struct {
	gorm.Model
	RegionName string `gorm:"not null"`
}
