package main

import "github.com/gofor-little/env"

func main() {
  loadEnvFile()
  createServer()
}

func loadEnvFile() {
  if err := env.Load(".env"); err != nil {
    panic(err)
  }
}
