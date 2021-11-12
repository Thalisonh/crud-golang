package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	CreatedAt time.Time `gorm:"<-:create"`
	Name string `gorm:"name"`
	Email string `gorm:"email"`
	Password string `gorm:"password"`
	Book []Book
}