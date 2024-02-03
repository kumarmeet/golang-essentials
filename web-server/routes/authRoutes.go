package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
)

func RegisterAuthRoutes(server *gin.Engine) {
	server.POST("/signups", controllers.Signup)
	server.POST("/login", controllers.Login)
}
