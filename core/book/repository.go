package book

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

type IBookRepository interface {
	GetBook (book_id int64) (*entity.Book, error)
	GetBooks() (*[]entity.Book, error)
	UpdateBook (book *entity.Book) (*entity.Book, error)
	DeleteBook(book_id int64) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{db: db}
}


func (r *BookRepository) GetBook(book_id int64) (*entity.Book, error){
	var bookModel entity.Book
	err := r.db.Where("id = ?", book_id).First(&bookModel).Error

	return &bookModel, err
}

func (r *BookRepository) GetBooks() (*[]entity.Book, error){
	var books []entity.Book
	return &books, r.db.First(&books).Error
}

func (r *BookRepository) UpdateBook (book *entity.Book) (*entity.Book, error) {
	return book, r.db.Save(book).Error
}

func (r *BookRepository) DeleteBook(book_id int64) error {
	return r.db.Where("id = ?", book_id).Delete(&entity.Book{}).Error
}

func (r *BookRepository) CreateBook (book *entity.Book) (*entity.Book, error) {
	return book, r.db.Create(&book).Error
}