package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token found!"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}
