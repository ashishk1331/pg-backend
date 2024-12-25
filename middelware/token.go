package middleware

import (
	"net/http"
	"pg-backend/config"
	"pg-backend/models"
	"pg-backend/util"

	"github.com/gin-gonic/gin"
)

func VerifyTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}
		token, ok := util.ValidateHeaderToken(authHeader)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		claims, err := util.ParseJWTToken(token, config.JwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", &models.User{
			Id:    claims.Id,
			Email: claims.Email,
			Role:  claims.Role,
		})
	}
}

func VerifyResetTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}
		token, ok := util.ValidateHeaderToken(authHeader)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		claims, err := util.ParseJWTToken(token, config.ResetSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", &models.User{
			Id:    claims.Id,
			Email: claims.Email,
			Role:  claims.Role,
		})
	}
}
