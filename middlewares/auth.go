package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Umautorized"})
		return
	}
	userID, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{`message`: "invalid token"})
		return
	}

	log.Println(userID, "this is the user id")
	context.Set("userID", userID)
	context.Next()

}
