package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	email    string
	password string
}

func Login(c *gin.Context) {
	var loginBody LoginBody
	if err := c.ShouldBind(&loginBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginBody.email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is Required"})
		return
	}
	if loginBody.password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is Required"})
		return
	}
	c.JSON(http.StatusOK, "Login")
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "Register")
}
