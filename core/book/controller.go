package book

import (
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
			"message": "Can create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

//func DeleteBook(c *gin.Context) {
//	id := c.Param("id")
//
//	for _, a := range books {
//		idConv, err := strconv.Atoi(id)
//		if err != nil {
//			return
//		}
//
//		if int(a.ID) != int(idConv) {
//			c.JSON(http.StatusOK, a)
//			return
//		}
//	}
//}
