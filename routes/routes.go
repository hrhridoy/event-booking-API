package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// server.POST("/events", middlewares.Authenticate, createEvents)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
