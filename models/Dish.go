package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Name        string      `gorm:"type:varchar(60);uniqueIndex:dish_name"`
	Descritpion string      `gorm:"type:text"`
	Duration    string      `gorm:"type:varchar(45)"`
	Users       []*User     `gorm:"many2many:user_dishes"`
	Favorites   []*User     `gorm:"many2many:user_favorites"`
	Categories  []*Category `gorm:"many2many:dish_categories"`
	Recipes     []*Recipe
	Quantities  []Quantity
}
