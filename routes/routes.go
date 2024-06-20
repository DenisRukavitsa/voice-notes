package routes

import "github.com/gin-gonic/gin"

func Register(ginEngine *gin.Engine) {
  ginEngine.GET("/health", health)
}