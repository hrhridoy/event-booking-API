package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/models"
	"github.com/hrhridoy/event-booking-API/utils"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		// fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events..."})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request. Empty token"})
		return
	}
	UserId, err := utils.VerifyTOken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request.", "err": err.Error()})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data..."})
		return
	}
	event.UserID = UserId
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

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event..."})
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

	eventById, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event..."})
		return
	}
	err = eventById.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete events..."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Delete Successfully..."})
}
