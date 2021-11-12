package routes

import (
	"github.com/Thalisonh/crud-golang/core/book"
	"github.com/Thalisonh/crud-golang/core/user"
	"github.com/Thalisonh/crud-golang/database"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	bookRepository := book.NewBookRepository(database.GetDb())
	bookService := book.NewBookService(bookRepository)
	bookController := book.NewBookController(bookService)

	userRepository := user.NewUserRepository(database.GetDb())
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	main := router.Group("api/v1")
	{
		users := main.Group("users/")
		{
			users.GET("/", userController.GetAll)
			users.GET("/:id", userController.GetUser)
			users.POST("/", userController.CreateUser)

			books := main.Group("book")
			{
				books.GET("/", bookController.ShowBooks)
				books.GET("/:id", bookController.ShowBook)
				books.POST("/", bookController.CreateBook)
				books.PUT("/:id", bookController.UpdateBook)
				books.DELETE("/:id", bookController.DeleteBook)
			}
		}

	}
	return router
}
