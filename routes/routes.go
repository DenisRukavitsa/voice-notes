package routes

import (
	"github.com/DenisRukavitsa/voice-notes/middleware"
	"github.com/gin-gonic/gin"
)

func Register(ginEngine *gin.Engine) {
	ginEngine.Use(middleware.CORS())
  ginEngine.GET("/health", health)
  ginEngine.POST("/transcribe", middleware.SaveUploadedFile, transcribe)

	ginEngine.POST("/register", register)
}