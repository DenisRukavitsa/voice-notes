package routes

import (
	"log"
	"net/http"

	"github.com/DenisRukavitsa/voice-notes/models"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("error binding user data", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
		return
	}

	err = user.Save()
	if err != nil {
		log.Println("error saving user", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error saving user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userID": user.ID})
}