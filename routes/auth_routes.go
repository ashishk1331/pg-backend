package routes

import (
	"pg-backend/controller/auth"
	middleware "pg-backend/middelware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", auth.Login)
		authRouter.POST("/register", auth.Register)
		authRouter.POST("/forgot-password", auth.ForgotPassword)
		authRouter.POST("/reset-password", middleware.VerifyTokenMiddleware(), auth.ResetPassword)
	}
}
