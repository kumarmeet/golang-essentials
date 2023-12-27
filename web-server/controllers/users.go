package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/models"
	"github.com/learning-webserver/utils"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	u, err := user.GetUser(id)

	userResponse := map[string]interface{}{
		"id":        u.ID,
		"name":      u.Name,
		"email":     u.Email,
		"createdat": u.Created_At,
		"updatedat": u.Updated_At,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Data successfully saved.", "user": userResponse})
}

func Login(ctx *gin.Context) {
	var user models.User

	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := user.ValidCreds(credentials.Email, credentials.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(u.Email, u.ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	userResponse := map[string]interface{}{
		"id":        u.ID,
		"email":     u.Email,
		"createdat": u.Created_At,
		"updatedat": u.Updated_At,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": true,
		"message": "Data successfully saved.",
		"user":    userResponse,
		"token":   token,
	})
}
