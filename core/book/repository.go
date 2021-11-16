package book

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

type IBookRepository interface {
	GetBook (bookId int64, userId int64) (*entity.Book, error)
	GetBooks(userId int64) (*[]entity.Book, error)
	UpdateBook (bookId int64, userId int64, book *entity.Book) (*entity.Book, error)
	DeleteBook(book *entity.Book, userId int64) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type RepositoryBook struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &RepositoryBook{db: db}
}


func (r *RepositoryBook) GetBook(bookId int64, userId int64) (*entity.Book, error){
	var bookModel entity.Book
	err := r.db.Where("id = ? AND user_id = ?", bookId, userId).First(&bookModel).Error

	return &bookModel, err
}

func (r *RepositoryBook) GetBooks(userId int64) (*[]entity.Book, error){
	var books []entity.Book
	return &books, r.db.Where("user_id = ?", userId).Find(&books).Error
}

func (r *RepositoryBook) UpdateBook (bookId int64, userId int64, book *entity.Book) (*entity.Book, error) {
	return book, r.db.Where("id = ? AND user_id = ?", bookId, userId).Save(&book).Error
}

func (r *RepositoryBook) DeleteBook(book *entity.Book, userId int64) error {
	return r.db.Where("id = ? AND user_id = ?", book.ID, userId).Delete(&entity.Book{}).Error
}

func (r *RepositoryBook) CreateBook (book *entity.Book) (*entity.Book, error) {
	return book, r.db.Create(&book).Error
}