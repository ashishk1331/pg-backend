package main

import (
	"pg-backend/config"
	"pg-backend/models"
	"pg-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.ConnectSQLDB()
	models.MigrateUser(config.DB)
	models.MigrateUserInfo(config.DB)
	router := gin.Default()
	routes.RegisterV1Group(router)
	router.Run(":8000")
}
