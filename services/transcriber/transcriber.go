// +build !mock

package transcriber

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(filePath string) (string, error) {
  log.Println("transcribing audio")
  client := openai.NewClient(os.Getenv("OPENAI_AUTH_TOKEN"))
  context := context.Background()

  request := openai.AudioRequest{
    Model: openai.Whisper1,
    FilePath: filePath,
  }
  response, err := client.CreateTranscription(context, request)
  if err != nil {
    log.Println("error transcribing audio", err)
    return "", err
  }

  return response.Text, nil
}
