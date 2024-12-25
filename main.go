package main

import (
	"pg-backend/config"
	middleware "pg-backend/middelware"
	"pg-backend/models"
	"pg-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.InitJWTSecret()
	config.ConnectSQLDB()
	models.MigrateUser(config.DB)
	models.MigrateUserInfo(config.DB)
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.RegisterV1Group(router)
	router.Run(":8000")
}
