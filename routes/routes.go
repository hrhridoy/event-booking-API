package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvents)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
