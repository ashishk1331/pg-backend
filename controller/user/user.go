package user

import (
	"net/http"
	"pg-backend/config"
	"pg-backend/models"
	"pg-backend/repository"
	"pg-backend/util"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	token, ok := util.ValidateHeaderToken(token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}
	x, err := util.ParseJWTToken(token, config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}
	user, err := repository.GetUserById(x.Id)
	user.Password = ""
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	var userInfo models.UserInfo
	if err := config.DB.Where("user_id = ?", x.Id).First(&userInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "User found",
		"user_info": userInfo,
		"user":      user,
	})
}
