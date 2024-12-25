package auth

import (
	"fmt"
	"net/http"
	"pg-backend/config"
	"pg-backend/models"
	"pg-backend/repository"
	"pg-backend/util"
	"strings"
	"time"

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
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := util.ComparePassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	user.Password = ""
	token, err := util.GenerateJWTToken(user, time.Hour*24, config.JwtSecret)
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
	_, err := repository.GetUserByEmail(input.Email)
	if err == nil {
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

func ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
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
	// check if user with email exits
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// iF email exits send password reset email
	// Generate JWT token with user email for 15 minutes
	token, err := util.GenerateJWTToken(user, time.Minute*15, config.ResetSecret)
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	if err := util.SendPasswordResetEmail(user, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "resetLink": resetLink})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset mail sent successfully", "resetLink": resetLink})
}

func ResetPassword(c *gin.Context) {
	var input struct {
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is Required"})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userData, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	existingUser, err := repository.GetUserById(userData.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	hashedPassword, err := util.Encrypt(input.Password)
	if err != nil {
		fmt.Print("Error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	existingUser.Password = hashedPassword
	if err := config.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func GetUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userData, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	existingUser, err := repository.GetUserById(userData.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	existingUser.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "User found", "user": existingUser})
}
