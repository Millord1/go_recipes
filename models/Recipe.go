package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Content string `gorm:"type:longtext"`
}
