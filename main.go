package main

import (
	"ct-backend/Config"
	"ct-backend/Route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	db := Config.SetUpDatabaseConnection()
	defer Config.CloseDatabaseConnection(db)

	server := gin.Default()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	// init route and DI
	Route.Init(server, db)

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
