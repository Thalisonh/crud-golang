package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint   `gorm: "id" gorm: "primaryKey"`
	Name        string `gorm: "name"`
	Description string `gorm: "description"`
	MediumPrice string `gorm: "medium_price"`
	Author      string `gorm: "author"`
	ImageURL    string `gorm: "img_url`

}
