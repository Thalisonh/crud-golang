package book

import "github.com/Thalisonh/crud-golang/database/entity"

type IBookService interface {
	GetBook (bookId int64) (*entity.Book, error)
	GetBooks() (*[]entity.Book, error)
	UpdateBook (bookId int64, book *entity.Book) (*entity.Book, error)
	DeleteBook (book *entity.Book) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type BookService struct {
	repository IBookRepository
}

func NewBookService(repository IBookRepository) IBookService {
	return &BookService{repository: repository}
}

func (s *BookService) GetBook (bookId int64) (*entity.Book, error){
	book, err := s.repository.GetBook(bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBooks() (*[]entity.Book, error){
	books, err := s.repository.GetBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookService) UpdateBook (bookId int64, book *entity.Book) (*entity.Book, error){
	book, err := s.repository.UpdateBook(bookId, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) DeleteBook (book *entity.Book) error{
	book, err := s.repository.GetBook(int64(book.ID))
	if err != nil {
		return err
	}

	bookError := s.repository.DeleteBook(book)
	if bookError != nil {
		return bookError
	}

	return nil
}

func (s *BookService) CreateBook (book *entity.Book) (*entity.Book, error){
	book, err := s.repository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, err
}


