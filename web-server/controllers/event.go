package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/models"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data, try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func GetEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	events, err := event.GetEvent(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func DeleteEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed request data!"})
		return
	}

	events, err := event.DeleteEvent(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}
func UpdateEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed request data!"})
		return
	}

	events, err := event.UpdateEvent(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func InsertEvent(ctx *gin.Context) {
	id := ctx.GetInt64("userId")

	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.UserId = id

	eventId, err := event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	eventData, _ := event.GetEvent(eventId)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Data successfully saved.", "event": eventData})
}
