package routes

import (
	"github.com/gin-gonic/gin"
	"pg-backend/controller/check"
	"pg-backend/controller/run"
)

func RegisterV1Group(router *gin.RouterGroup) {
	router.POST("/check", check.Post)
	router.POST("/run", run.Post)
}