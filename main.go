package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/db"
	"github.com/hrhridoy/event-booking-API/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")

}
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
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
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
