package routes

import (
	"github.com/DenisRukavitsa/voice-notes/middleware"
	"github.com/DenisRukavitsa/voice-notes/modules/file"
	"github.com/DenisRukavitsa/voice-notes/modules/transcription"
	"github.com/DenisRukavitsa/voice-notes/modules/user"
	"github.com/gin-gonic/gin"
)

func Register(ginEngine *gin.Engine) {
	ginEngine.Use(middleware.CORS())
  ginEngine.GET("/health", health)
  ginEngine.POST("/transcribe", file.SaveUploadedFile, transcription.Transcribe)
	ginEngine.POST("/register", user.Register)
}