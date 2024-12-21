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
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := util.ComparePassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	user.Password = ""
	token, err := util.GenerateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, gin.H{"message": "Logged In Successfully", "token": token})
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
