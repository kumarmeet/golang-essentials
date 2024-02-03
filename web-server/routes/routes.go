package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
	"github.com/learning-webserver/middlewares"
)

func MainRoutes(server *gin.Engine) {
	server.MaxMultipartMemory = 10 << 20 // 10mb

	server.Use(middlewares.RateLimit())

	server.POST("/xlsx/upload", middlewares.UploadMultipleFilesMiddleware([]string{".xls", ".xlsx", ".csv"}), controllers.ImportCsvXlsx)
	server.POST("/image/upload/:event_id", middlewares.UploadMultipleFilesMiddleware([]string{".png", ".jpg", ".jpeg", ".gif"}), controllers.UploadFile)

	RegisterAuthRoutes(server)
	RegisterUserRoutes(server)
	RegisterEventRoutes(server)
}
