package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/models"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User Name or Password."})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created successfully."})
}
