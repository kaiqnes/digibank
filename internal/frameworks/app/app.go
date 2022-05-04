package app

import (
	"digibank/internal/frameworks/database"
	"digibank/internal/frameworks/dependencyInjection"
	"digibank/internal/frameworks/router"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	routes := router.NewRouter()
	db := database.NewSession()

	di := dependencyInjection.NewDependencyInjection(routes, db)
	di.SetupDependencies()

	return routes
}
