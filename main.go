package main

import (
	"github.com/gin-gonic/gin"
	"pg-backend/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterV1Group(router)

	router.Run(":8000")
}
