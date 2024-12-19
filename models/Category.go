package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(45);uniqueIndex:cat_name"`
}
