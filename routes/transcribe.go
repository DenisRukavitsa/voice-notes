package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/DenisRukavitsa/voice-notes/services/transcriber"
	"github.com/gin-gonic/gin"
)

func transcribe(ginContext *gin.Context) {
	uploadedFilePath := ginContext.GetString("uploadedFilePath")
	transcription, err := transcriber.Transcribe(uploadedFilePath)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "error transcribing audio"})
		return
	}

	err = os.Remove(uploadedFilePath)
	if err != nil {
		log.Println("error removing uploaded file", err)
	}

	ginContext.JSON(http.StatusOK, gin.H{"transcription": transcription})
}
