package user

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	services IUserService
}

func NewUserController(service IUserService) UserController {
	return UserController{services: service}
}

func (r *UserController) GetAll(c *gin.Context){
	users, err := r.services.GetUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (r *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	user, err := r.services.GetUser(int64(newId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not find user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *UserController) CreateUser(c *gin.Context) {
	var newUser entity.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := r.services.CreateUser(&newUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
