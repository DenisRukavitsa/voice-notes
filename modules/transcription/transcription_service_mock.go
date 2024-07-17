// +build mock

package transcription

func transcribeAudioFile(filePath string) (string, error) {
  return "mocked transcription", nil
}
