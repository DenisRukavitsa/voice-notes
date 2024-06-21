package filemanager

import (
	"log"
	"os"
)

func Remove(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("error removing uploaded file", err)
	}
}
