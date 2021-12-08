package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `form:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `form:"password" gorm:"size:255"`
}
