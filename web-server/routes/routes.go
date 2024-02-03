package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
	"github.com/learning-webserver/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.MaxMultipartMemory = 10 << 20 // 10mb

	server.POST("/signups", controllers.Signup)
	server.POST("/login", controllers.Login)
	server.GET("/events", controllers.GetEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events/:id", controllers.GetEventById)
	authenticated.POST("/events", controllers.InsertEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEventById)
	authenticated.DELETE("/events/:id", controllers.DeleteEventById)

	users := server.Group("/users")
	users.Use(middlewares.Authenticate)
	users.GET("", controllers.Users)

	server.POST("/xlsx/upload", middlewares.UploadMultipleFilesMiddleware([]string{".xls", ".xlsx", ".csv"}), controllers.ImportCsvXlsx)
	server.POST("/image/upload/:event_id", middlewares.UploadMultipleFilesMiddleware([]string{".png", ".jpg", ".jpeg", ".gif"}), controllers.UploadFile)
}
