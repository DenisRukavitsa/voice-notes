package user

import (
	"log"
	"net/http"

	"github.com/DenisRukavitsa/voice-notes/modules/auth"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user UserDto
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("error binding user data", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
		return
	}

	_, err = findUserByEmail(user.Email)
	if err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user email already registered"})
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

func Login(context *gin.Context) {
	var user UserDto
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("error binding user data", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
		return
	}

	userId, err := checkUserCredentials(user)
	if userId == "" || err != nil {
		log.Println("error checking user credentials", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user credentials"})
		return
	}

	jwtToken, err := auth.GenerateToken(user.Email, userId)
	if err != nil {
		log.Println("error generating jwt token", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error generating jwt token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": jwtToken})
}
