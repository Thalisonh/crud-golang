package user

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerUser struct {
	services IUserService
}

func NewUserController(service IUserService) ControllerUser {
	return ControllerUser{services: service}
}

func (r *ControllerUser) GetAll(c *gin.Context){
	users, err := r.services.GetUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (r *ControllerUser) GetUser(c *gin.Context) {
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

func (r *ControllerUser) CreateUser(c *gin.Context) {
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

func (r *ControllerUser) Delete(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	user, errFind := r.services.GetUser(int64(idConv))

	if errFind != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Entity not found",
		})
		return
	}

	errDelete := r.services.DeleteUser(user)

	if errDelete != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": errDelete,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}

func (r *ControllerUser) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUser entity.User

	err := c.ShouldBindJSON(&updateUser)
	idUpd, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a Integer",
		})
		return
	}

	user, errGet := r.services.GetUser(int64(idUpd))

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Entity not found",
		})
		return
	}

	newUser, err := r.services.UpdateUser(int64(user.ID), &updateUser)

	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, &newUser)
}