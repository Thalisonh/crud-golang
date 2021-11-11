package book

import "github.com/Thalisonh/crud-golang/database/entity"

type IBookService interface {
	GetBook (book_id int64) (*entity.Book, error)
	GetBooks() (*[]entity.Book, error)
	UpdateBook (book *entity.Book) (*entity.Book, error)
	DeleteBook (book_id int64) error
	CreateBook (book *entity.Book) (*entity.Book, error)
}

type BookService struct {
	repository IBookRepository
}

func NewBookService(repository IBookRepository) IBookService {
	return &BookService{repository: repository}
}

func (s *BookService) GetBook (book_id int64) (*entity.Book, error){
	book, err := s.repository.GetBook(book_id)
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

func (s *BookService) UpdateBook (book *entity.Book) (*entity.Book, error){
	book, err := s.repository.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) DeleteBook (book_id int64) error{
	book, err := s.repository.GetBook(book_id)
	if err != nil {
		return err
	}

	bookError := s.DeleteBook(int64(book.ID))
	if bookError != nil {
		return bookError
	}

	return nil
}

func (s *BookService) CreateBook (book *entity.Book) (*entity.Book, error){
	book, err := s.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, err
}


