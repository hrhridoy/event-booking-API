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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later..."})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data..."})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		// fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later..."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "even created", "event:": event})

}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID..."})
		return
	}
	eventById, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event..."})
		return
	}
	context.JSON(http.StatusOK, eventById)
}
