package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	CreatedAt time.Time `gorm:"<-:create"`
	Name        string `gorm: "name"`
	Description string `gorm: "description"`
	MediumPrice string `gorm: "medium_price"`
	Author      string `gorm: "author"`
	ImageURL    string `gorm: "img_url`
	UserID      uint   `gorm:"user_id"`
}
