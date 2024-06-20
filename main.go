package main

import (
	"github.com/DenisRukavitsa/voice-notes/server"
	"github.com/gofor-little/env"
)

func main() {
  loadEnvFile()
  ginEngine := server.Create()
  server.Run(ginEngine)
}

func loadEnvFile() {
  if err := env.Load(".env"); err != nil {
    panic(err)
  }
}
