package middleware

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(ginContext *gin.Context) {
	file, err := ginContext.FormFile("file")
	if err != nil {
		log.Println("error forming file", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "audio file missing"})
		return
	}
	log.Println("got file ", file.Filename)

	fileExtension := filepath.Ext(file.Filename)
	filePath := fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), fileExtension)
	err = ginContext.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("error saving uploaded file", err)
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error saving file"})
		return
	}
	log.Println("saved uploaded file to ", filePath)

	ginContext.Set("uploadedFilePath", filePath)
	ginContext.Next()
}
