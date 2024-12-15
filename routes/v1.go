package routes

import (
	"github.com/gin-gonic/gin"
	"pg-backend/controller/check"
	"pg-backend/controller/run"
	"pg-backend/controller/generate"
)

func RegisterV1Group(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/check", check.Post)
		v1.POST("/run", run.Post)
		v1.GET("/generate/:id", generate.Get)
	}
}