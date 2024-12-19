package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(60);uniqueIndex:dish_name"`
	Descritpion string       `gorm:"type:text"`
	Duration    string       `gorm:"type:varchar(45)"`
	Categories  []Category   `gorm:"many2many:dish_categories"`
	Ingredients []Ingredient `gorm:"many2many:dish_ingredients"`
	Recipes     []Recipe     `gorm:"foreignKey:ID"`
}
