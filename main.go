package main

import (
	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/server"
	"github.com/gofor-little/env"
)

func main() {
	loadEnvFile()
	databaseClient := database.Connect()
	defer database.Disconnect(databaseClient)

	ginEngine := server.Create()
	server.Run(ginEngine)
}

func loadEnvFile() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}
}
