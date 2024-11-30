package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/models"
)

func registerForEvent(context *gin.Context) {

	userId := context.MustGet("userID").(int64)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event", "event": event})

}

func unregisterFromEvent(context *gin.Context) {

	userId := context.MustGet("userID").(int64)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	var event models.Event

	event.ID = eventId

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Unregistered from event", "event": event})

}
