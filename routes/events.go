package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
		})
		return
	}
	if events == nil {
		events = []models.Event{}
	}
	context.JSON(http.StatusOK, gin.H{
		"data": events,
	})

}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	userID := context.MustGet("userID").(int64)

	event.UserID = userID
	err = event.Save()

	fmt.Println("Event error", err)

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "Event": event})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "bad reques"},
		)
		return
	}

	_event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("event with %d not found", id)})
	}
	userID := context.MustGet("userID").(int64)
	if _event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	event.ID = id

	err = event.Upate()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return

	}
	context.JSON(http.StatusOK, gin.H{"message": "Event created", "Event": event})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "bad reques"},
		)
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("event with %d not found", id)})
	}

	userID := context.MustGet("userID").(int64)
	if event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return

	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "bad reques"},
		)
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("event with %d not found", id)})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": event,
	})

}
