package book

import "github.com/Thalisonh/crud-golang/database/entity"

type IBookService interface {
	GetBook (bookId int64, userId int64) (*entity.Book, error)
	GetBooks(userId int64) (*[]entity.Book, error)
	UpdateBook (bookId int64, userId int64, book *entity.Book) (*entity.Book, error)
	DeleteBook (book *entity.Book, userId int64) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type ServiceBook struct {
	repository IBookRepository
}

func NewBookService(repository IBookRepository) IBookService {
	return &ServiceBook{repository: repository}
}

func (s *ServiceBook) GetBook (bookId int64, userId int64) (*entity.Book, error){
	book, err := s.repository.GetBook(bookId, userId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBook) GetBooks(userId int64) (*[]entity.Book, error){
	books, err := s.repository.GetBooks(userId)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *ServiceBook) UpdateBook (bookId int64, userId int64, book *entity.Book) (*entity.Book, error){
	book, err := s.repository.UpdateBook(bookId, userId, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBook) DeleteBook (book *entity.Book, userId int64) error{
	book, err := s.repository.GetBook(int64(book.ID), userId)
	if err != nil {
		return err
	}

	bookError := s.repository.DeleteBook(book, userId)
	if bookError != nil {
		return bookError
	}

	return nil
}

func (s *ServiceBook) CreateBook (book *entity.Book) (*entity.Book, error){
	book, err := s.repository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, err
}


