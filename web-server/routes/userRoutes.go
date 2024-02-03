package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/controllers"
	"github.com/learning-webserver/middlewares"
)

func RegisterUserRoutes(server *gin.Engine) {
	users := server.Group("/users")
	users.Use(middlewares.Authenticate)
	users.GET("", controllers.Users)
}
