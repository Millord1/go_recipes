package models

import "gorm.io/gorm"

type Quantity struct {
	gorm.Model
	Num        uint16     `gorm:"type:uint"`
	Unit       string     `gorm:"type:varchar(45)"`
	Dish       Dish       `gorm:"foreignKey:ID;uniqueIndex:dish_ingredient"`
	Ingredient Ingredient `gorm:"foreignKey:ID;uniqueIndex:dish_ingredient"`
}
