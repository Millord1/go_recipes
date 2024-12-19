package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string `gorm:"type:varchar(60); uniqueIndex:ing_name"`
}
