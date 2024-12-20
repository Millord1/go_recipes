package models

import "gorm.io/gorm"

type Quantity struct {
	gorm.Model
	Num          uint16 `gorm:"type:uint"`
	Unit         string `gorm:"type:varchar(45)"`
	DishID       uint
	IngredientID uint
}
