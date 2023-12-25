package main

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/db"
	"github.com/learning-webserver/routes"
)

func main() {
	server := gin.Default()
	db.InitDB()

	routes.RegisterEventRoutes(server)

	server.Run(":4000")
}
