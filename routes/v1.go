package routes

import (
	"pg-backend/controller/check"
	"pg-backend/controller/generate"
	"pg-backend/controller/run"
	"pg-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func RegisterV1Group(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/check", check.Post)
		v1.POST("/run", run.Post)
		v1.GET("/generate/:id", generate.Get)
		v1.GET("/user", user.Get)
	}
	RegisterAuthRoutes(v1)
}
