package routes

import (
	"pg-backend/controller/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", auth.Login)
		authRouter.POST("/register", auth.Register)
		authRouter.POST("/forgot-password", auth.ForgotPassword)
	}
}
