package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/learning-webserver/config"
	"github.com/learning-webserver/db"
	"github.com/learning-webserver/routes"
)

func main() {
	server := gin.Default()

	cfg, err := config.LoadConfig()

	fmt.Println(cfg)

	if err != nil {
		log.Fatal("Error loading config file")
	}

	err = godotenv.Load("mysql.env")

	db.InitDB(cfg.AppDB)

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.MainRoutes(server)

	server.Run(":" + cfg.AppPort)
}
