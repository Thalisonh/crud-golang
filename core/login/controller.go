package login

import (
	"fmt"
	"net/http"

	"github.com/Thalisonh/crud-golang/database"
	"github.com/Thalisonh/crud-golang/database/entity"
	"github.com/Thalisonh/crud-golang/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDb()

	var p entity.Login

	err := c.ShouldBindJSON(p)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not find user",
		})
		return
	}

	var user entity.User 

	dbError := db.Where("email = ?", p.Email).First(&user).Error

	if dbError != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Can not find user",
			})
			return
	}

	if user.Password != services.Sha256Encoder(p.Password) {
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invalid credentials",
			})
		}
		return
	}

	token, err := services.NewJwtService().GeneretedToken(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
}