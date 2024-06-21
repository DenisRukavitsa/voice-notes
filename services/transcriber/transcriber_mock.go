// +build mock

package transcriber

func Transcribe(filePath string) (string, error) {
  return "mocked transcription", nil
}
