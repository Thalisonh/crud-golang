package book

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

type IBookRepository interface {
	GetBook (bookId int64, userId int64) (*entity.Book, error)
	GetBooks(userId int64) (*[]entity.Book, error)
	UpdateBook (bookId int64, book *entity.Book) (*entity.Book, error)
	DeleteBook(book *entity.Book, userId int64) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{db: db}
}


func (r *BookRepository) GetBook(bookId int64, userId int64) (*entity.Book, error){
	var bookModel entity.Book
	err := r.db.Where("id = ? AND user_id = ?", bookId, userId).First(&bookModel).Error

	return &bookModel, err
}

func (r *BookRepository) GetBooks(userId int64) (*[]entity.Book, error){
	var books []entity.Book
	return &books, r.db.Where("user_id = ?", userId).Find(&books).Error
}

func (r *BookRepository) UpdateBook (id int64, book *entity.Book) (*entity.Book, error) {
	return book, r.db.Where("id = ?", id).Save(&book).Error
}

func (r *BookRepository) DeleteBook(book *entity.Book, userId int64) error {
	return r.db.Where("id = ? AND user_id = ?", book.ID, userId).Delete(&entity.Book{}).Error
}

func (r *BookRepository) CreateBook (book *entity.Book) (*entity.Book, error) {
	return book, r.db.Create(&book).Error
}