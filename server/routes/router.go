package routes

import (
	"github.com/Thalisonh/crud-golang/core/book"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("book")
		{
			books.GET("/:id", book.ShowBook)
			books.GET("/", book.ShowBooks)
			books.POST("/", book.CreateBook)
			books.DELETE("/:id", book.DeleteBook)
		}
	}
	return router
}
