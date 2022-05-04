package main

import (
	"digibank/internal/frameworks/app"
	"log"
)

func main() {
	server := app.Setup()
	if err := server.Run(); err != nil {
		log.Fatalf("failed to start server - err %v", err)
	}
}
