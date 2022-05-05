package app

import (
	"digibank/internal/frameworks/database"
	"digibank/internal/frameworks/dependencyInjection"
	"digibank/internal/frameworks/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func Setup() *gin.Engine {
	loadEnvVariables()

	routes := router.NewRouter()
	db := database.NewSession()

	di := dependencyInjection.NewDependencyInjection(routes, db)
	di.SetupDependencies()

	return routes
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
