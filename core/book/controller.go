package book

import (
	"github.com/Thalisonh/crud-golang/database"
	"github.com/Thalisonh/crud-golang/database/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []entity.Book{
	{ID: 1, Author: "Thalison", Description: "teste", MediumPrice: "3,15"},
}

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not Found",
		})
		return
	}

	db := database.GetDb()

	var book entity.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not find book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func ShowBooks(c *gin.Context) {
	//db := database.GetDb()

	//var books []models.Book
	//err := db.First(&books).Error

	//if err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "Can list books: " + err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, &books)
}

func CreateBook(c *gin.Context) {
	//db := database.GetDb()
	var newBook entity.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can bind Json: " + err.Error(),
		})
		return
	}

	//err = db.Create(&book).Error

	books = append(books, newBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, books)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		idConv, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		if int(a.ID) != int(idConv) {
			c.JSON(http.StatusOK, a)
			return
		}
	}
}
