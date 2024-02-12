package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/db"
	"github.com/hrhridoy/event-booking-API/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
