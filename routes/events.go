package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		// fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events..."})
	}
	if events == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No Events Created by Users."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	// one way of doing the Auth check
	// userId, err := middlewares.Authenticate(context)
	// if err != nil {
	// 	return
	// }
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data..."})
		return
	}
	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		// fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events..."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created.", "event": event})

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID..."})
		return
	}
	eventById, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event..."})
		return
	}
	context.JSON(http.StatusOK, eventById)
}
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID..."})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event..."})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event data..."})
		return
	}
	updatedEvent.ID = eventId

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update events..."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully..."})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID..."})
		return
	}
	userId := context.GetInt64("userId")
	eventById, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event..."})
		return
	}
	if eventById.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delate event."})
		return
	}
	err = eventById.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete events..."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Delete Successfully..."})
}
