package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrhridoy/event-booking-API/utils"
)

// Another way of doing authentication check before handling request method.

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request."})
		return
	}
	userId, err := utils.VerifyTOken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request.", "err": err.Error()})
		return
	}
	context.Set("userId", userId)
	context.Next()

}

// One way  of doing the Auth check

// func Authenticate(context *gin.Context) (int64, error) {
// 	token := context.Request.Header.Get("Authorization")
// 	if token == "" {
// 		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request."})
// 		return 0, errors.New("Empty token")
// 	}
// 	userId, err := utils.VerifyTOken(token)
// 	if err != nil {
// 		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request.", "err": err.Error()})
// 		return 0, err
// 	}
// 	return userId, nil
// }
