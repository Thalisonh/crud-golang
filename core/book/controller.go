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
	idUser := c.Param("id_user")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	newUserId, err := strconv.Atoi(idUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User id must be a integer",
		})
	}

	book, err := s.services.GetBook(int64(newId), int64(newUserId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not find book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

func (s *BookController) ShowBooks(c *gin.Context) {
	id := c.Param("id_user")
	idUser, err := strconv.Atoi(id)

	books, err := s.services.GetBooks(int64(idUser))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, &books)
}

func (s *BookController) CreateBook(c *gin.Context) {
	var newBook entity.Book
	idUser := c.Param("id_user")
	id, err := strconv.Atoi(idUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a Integer",
		})
		return
	}

	c.ShouldBindJSON(&newBook)

	newBook.UserID = int(id)

	user, err := s.services.CreateBook(&newBook)

	if err != nil{
		c.JSON(http.StatusNotModified, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *BookController) UpdateBook(c *gin.Context){
	var newBook entity.Book
	id := c.Param("id")
	idUser := c.Param("id_user")

	newId, errConv := strconv.Atoi(id)
	newIdUser, errUser := strconv.Atoi(idUser)

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	if errUser != nil {
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

	_, errFind := s.services.GetBook(int64(newId), int64(newIdUser))

	if errFind != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Entity not found",
		})
		return
	}

	newBook.UserID = int(newIdUser)

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
	idUser := c.Param("id_user")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID must be a integer",
		})
		return
	}

	newIdUser, errUser := strconv.Atoi(idUser)
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID must be a integer",
		})
		return
	}

	bookDeleted, err := s.services.GetBook(int64(newId), int64(newIdUser))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": bookDeleted,
		})
		return
	}

	deletedErr := s.services.DeleteBook(bookDeleted, int64(newIdUser))

	if deletedErr != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": deletedErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, deletedErr)
}
