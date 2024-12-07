package routes

import (
	"github.com/gin-gonic/gin"
	"pg-backend/controller"
)

func RegisterV1Group(router *gin.RouterGroup) {
	router.POST("/check", check.Get)
}