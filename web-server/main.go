package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/learning-webserver/db"
	"github.com/learning-webserver/routes"
)

func main() {
	server := gin.Default()
	db.InitDB()

	err := godotenv.Load(".env")

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.RegisterEventRoutes(server)

	server.Run(":4000")
}
