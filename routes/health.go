package routes

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func health(ginContext *gin.Context) {
  ginContext.JSON(http.StatusOK, gin.H{"message": "server healthy"})
}