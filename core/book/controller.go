package book

import (
	"fmt"
	"github.com/Thalisonh/crud-golang/database/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	services IBookService
}

func NewBookController(services IBookService) BookController {
	return BookController{services: services}
}

func (s *BookController) ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	book, err := s.services.GetBook(int64(newId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not find book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

func (s *BookController) ShowBooks(c *gin.Context) {
	books, err := s.services.GetBooks()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, &books)
}

func (s *BookController) CreateBook(c *gin.Context) {
	var newBook entity.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can bind Json: " + err.Error(),
		})
		return
	}

	book, err := s.services.CreateBook(&newBook)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (s *BookController) UpdateBook(c *gin.Context){
	var newBook entity.Book
	id := c.Param("id")

	newId, errConv := strconv.Atoi(id)

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	err := c.ShouldBindJSON(&newBook)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, errFind := s.services.GetBook(int64(newId))

	if errFind != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Entity not found",
		})
		return
	}

	bookUpdated, err := s.services.UpdateBook(int64(newId), &newBook)

	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "Entity not modified",
		})
		return
	}

	c.JSON(http.StatusOK, bookUpdated)
}

func (s *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	fmt.Errorf("id:", id)
	//
	newId, err := strconv.Atoi(id)
	//
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID must be a integer",
		})
		return
	}

	bookDeleted, err := s.services.GetBook(int64(newId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": bookDeleted,
		})
		return
	}

	deletedErr := s.services.DeleteBook(bookDeleted)

	if deletedErr != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": deletedErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, deletedErr)
	//c.JSON(http.StatusOK, deletedErr)
}
