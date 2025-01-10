package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	Order   int    `gorm:"type:int"`
	Title   string `gorm:"type:varchar(45)"`
	Content string `gorm:"type:longtext"`
	DishID  uint
}
