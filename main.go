package main

import (
	_ "digibank/docs"
	"digibank/internal/frameworks/app"
	_ "github.com/swaggo/files" // swagger embed files
	"log"
)

// @title           Digibank API
// @version         1.0
// @description     This is a code challenge that manages some bank routines.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Caique Nunes
// @contact.url    https://www.linkedin.com/in/caique-nunes/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	server := app.Setup()
	if err := server.Run(); err != nil {
		log.Fatalf("failed to start server - err %v", err)
	}
}
