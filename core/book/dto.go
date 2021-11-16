package book

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"time"
)

type BookRequest struct {
	ID int64 `json:"id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	MediumPrice string `json: "medium_price"`
	Author      string `json: "author"`
	ImageURL    string `json: "img_url`
	UserID      int   `json:"user_id"`
}

type BookResponse struct {
	ID int64 `json:"id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	MediumPrice string `json: "medium_price"`
	Author      string `json: "author"`
	ImageURL    string `json: "img_url`
	UserID      int   `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *BookRequest) ToEntity() *entity.Book {
	return &entity.Book{
		Name:        r.Name,
		Description: r.Description,
		MediumPrice: r.MediumPrice,
		Author:      r.Author,
		ImageURL:    r.ImageURL,
		UserID:      r.UserID,
	}
}

func NewBookResponseEntity(b entity.Book) *BookResponse {
	return &BookResponse{
		ID:          int64(b.UserID),
		Name:        b.Name,
		Description: b.Description,
		MediumPrice: b.MediumPrice,
		Author:      b.Author,
		ImageURL:    b.ImageURL,
		UserID:      b.UserID,
	}
}