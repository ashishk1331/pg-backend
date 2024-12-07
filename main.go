package main

import (
	"github.com/gin-gonic/gin"
	"pg-backend/routes"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	routes.RegisterV1Group(v1)

	router.Run(":8000")
}
