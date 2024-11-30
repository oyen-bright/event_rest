package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/models"
	"github.com/oyen-bright/event_REST/utils"
)

func signup(contex *gin.Context) {
	var user models.User
	err := contex.BindJSON(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err = user.Save()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	contex.JSON(http.StatusCreated, gin.H{"message": "User created"})

}

func login(context *gin.Context) {

	var user models.User

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToke(user.Email, user.ID)
	if err != nil {

		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal sever error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "oka", "token": token})

}
