package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/db"
	"github.com/learning-webserver/models"
)

func main() {
	server := gin.Default()
	db.InitDB()

	server.GET("/events", func(ctx *gin.Context) {
		events, err := models.GetAllEvents()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data, try again later"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
	})

	server.POST("/events", func(ctx *gin.Context) {
		var event models.Event

		err := ctx.ShouldBindJSON(&event)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed request data!"})
			return
		}

		_, err = event.Save()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not saved data!"})
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Data successfully saved.", "event": event})
	})

	server.Run(":4000")
}
