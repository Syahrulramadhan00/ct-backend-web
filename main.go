package main

import (
	"ct-backend/Config"
	"ct-backend/Route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	if loc, err := time.LoadLocation("Asia/Jakarta"); err != nil {
		panic(err)
	} else {
		time.Local = loc
	}

	db := Config.SetUpDatabaseConnection()
	defer Config.CloseDatabaseConnection(db)

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Origin,Content-Type,Accept,User-Agent,Content-Length,Authorization"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"},
	}))

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error loading .env file: %v", err)
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
