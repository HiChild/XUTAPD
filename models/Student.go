package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	StudentName string `gorm:"size:20;not null;unique" form:"student_name"`
	Password string `gorm:"size:255;not null;" form:"password"`
}