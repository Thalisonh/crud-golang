package routes

import (
	"github.com/Thalisonh/crud-golang/core/book"
	"github.com/Thalisonh/crud-golang/database"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	bookRepository := book.NewBookRepository(database.GetDb())
	bookService := book.NewBookService(bookRepository)
	bookController := book.NewBookController(bookService)

	main := router.Group("api/v1")
	{
		books := main.Group("book")
		{
			books.GET("/", bookController.ShowBooks)
			books.GET("/:id", bookController.ShowBook)
			books.POST("/", bookController.CreateBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}
	}
	return router
}
