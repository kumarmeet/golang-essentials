package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
	"github.com/learning-webserver/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events/:id", controllers.GetEventById)
	authenticated.POST("/events", controllers.InsertEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEventById)
	authenticated.DELETE("/events/:id", controllers.DeleteEventById)

	server.POST("/upload", middlewares.UploadMultipleFilesMiddleware(), controllers.UploadFile)

	server.GET("/events", controllers.GetEvents)
	server.POST("/signups", controllers.Signup)
	server.POST("/login", controllers.Login)
}
