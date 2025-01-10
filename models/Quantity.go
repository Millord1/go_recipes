package models

import "gorm.io/gorm"

type Quantity struct {
	gorm.Model
	Num          uint16 `gorm:"type:uint;uniqueIndex:quantityIndex"`
	Unit         string `gorm:"type:varchar(45);uniqueIndex:quantityIndex"`
	DishID       uint
	IngredientID uint
}
