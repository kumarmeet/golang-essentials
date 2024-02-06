package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/learning-webserver/config"
	"github.com/learning-webserver/db"
	"github.com/learning-webserver/routes"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
)

func main() {
	server := gin.Default()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	})

	server.Use(func(ctx *gin.Context) {
		corsMiddleware.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	})

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

	// Create an HTTP/2 server with the default transport
	http2Server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: server,
	}

	http2.ConfigureServer(http2Server, &http2.Server{})

	http2Server.ListenAndServe()
}
