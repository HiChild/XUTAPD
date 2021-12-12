package models

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	TeacherName string `gorm:"size:20;not null;unique" form:"teacher_name"`
	Password string `gorm:"size:255;not null;" form:"password"`
}