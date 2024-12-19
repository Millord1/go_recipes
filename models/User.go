package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"-"`
	Username  string `gorm:"type:varchar(45);uniqueIndex:username_unique;not null;<-" json:"username" fake:"{username}"`
	Email     string `gorm:"type:varchar(60);uniqueIndex:email_unique;not null;<-" json:"-" fake:"email"`
	Password  string `gorm:"type:varchar(65);not null;<-" json:"-" fake:"password"`
	Totp      string `gorm:"type:varchar(60);<-" json:"-"`
	Dishes    []Dish `gorm:"many2many:user_dishes"`
	Favorites []Dish `gorm:"many2many:user_favorites"`
}
