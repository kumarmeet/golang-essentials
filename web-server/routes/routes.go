package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventById)
	server.POST("/events", controllers.InsertEvent)
	server.PUT("/events/:id", controllers.UpdateEventById)
	server.DELETE("/events/:id", controllers.DeleteEventById)
}
