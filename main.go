package main

import (
	"github.com/gin-gonic/gin"
	"pg-backend/controller"
)

func main() {
	r := gin.Default()

	r.POST("/check", check.Get)
	r.Run(":8000")
}
