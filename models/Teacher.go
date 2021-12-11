package models

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdateDAt time.Time
	DeleteDAt gorm.DeletedAt
	UserName string `gorm:"size:20;not null;unique" form:"username"`
	Password string `gorm:"size:255;not null;" form:"password"`
}