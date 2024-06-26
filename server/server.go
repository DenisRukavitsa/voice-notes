package server

import (
	"github.com/DenisRukavitsa/voice-notes/routes"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	ginEngine := gin.Default()
	err := ginEngine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	routes.Register(ginEngine)
	return ginEngine
}

func Run(ginEngine *gin.Engine) {
	err := ginEngine.Run()
	if err != nil {
		panic(err)
	}
}
