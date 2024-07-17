package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user UserModel
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("error binding user data", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
		return
	}

	userId, err := save(user)
	if err != nil {
		log.Println("error saving user", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error saving user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userId": userId})
}