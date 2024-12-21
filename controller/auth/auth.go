package auth

import (
	"net/http"
	"pg-backend/config"
	"pg-backend/models"
	"pg-backend/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginBody models.User
	if err := c.ShouldBind(&loginBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginBody.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is Required"})
		return
	}
	if loginBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is Required"})
		return
	}
	c.JSON(http.StatusOK, "Login")
}

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is Required"})
		return
	}
	if !util.ValidateEmail(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is Invalid"})
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is Required"})
		return
	}
	// Check if the email already exists in the database
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}
	hashedPassword, err := util.Encrypt(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input.Password = hashedPassword
	if err := config.DB.Create(&input).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var userInfo models.UserInfo
	userInfo.UserID = input.Id
	userInfo.Username = strings.Split(input.Email, "@")[0]
	userInfo.Fullname = util.GenerateRandomString(10)
	if err := config.DB.Create(&userInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully",
		"userId":   input.Id,
		"userName": userInfo.Username,
		"fullName": userInfo.Fullname,
		"imageUrl": userInfo.ImageUrl,
		"email":    input.Email})
}
