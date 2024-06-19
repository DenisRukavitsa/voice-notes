package main

import "github.com/gin-gonic/gin"

func createServer() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.SetTrustedProxies([]string{"localhost"})
	ginEngine.Run()
	return ginEngine
}
