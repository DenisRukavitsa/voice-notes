package routes

import (
	"net/http"
	"path/filepath"
	"slices"

	"github.com/DenisRukavitsa/voice-notes/services/filemanager"
	"github.com/DenisRukavitsa/voice-notes/services/transcriber"
	"github.com/gin-gonic/gin"
)

func transcribe(ginContext *gin.Context) {
	uploadedFilePath := ginContext.GetString("uploadedFilePath")
  defer filemanager.Remove(uploadedFilePath)

  fileExtension := filepath.Ext(uploadedFilePath)
  allowedFileExtensions := []string{".mp3", ".mp4", ".mpeg", ".mpga", ".m4a", ".wav", ".webm"}
  if !slices.Contains(allowedFileExtensions, fileExtension) {
    ginContext.JSON(http.StatusBadRequest, gin.H{"error": "invalid audio file extension"})
    return
  }

	transcription, err := transcriber.Transcribe(uploadedFilePath)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "error transcribing audio"})
    return
	}

  ginContext.JSON(http.StatusOK, gin.H{"transcription": transcription})
}
