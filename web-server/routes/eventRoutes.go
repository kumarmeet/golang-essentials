package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
	"github.com/learning-webserver/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", controllers.GetEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events/:id", controllers.GetEventById)
	authenticated.POST("/events", controllers.InsertEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEventById)
	authenticated.DELETE("/events/:id", controllers.DeleteEventById)
}
