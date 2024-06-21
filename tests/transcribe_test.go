package tests

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/server"
	"github.com/stretchr/testify/assert"
)

func TestTranscribeRouteOkResponse(t *testing.T) {
  var buffer bytes.Buffer
  multipartWriter := multipart.NewWriter(&buffer)
  fileWriter, _ := multipartWriter.CreateFormFile("file", "jackhammer.wav")
  fileData, _ := os.ReadFile("testdata/jackhammer.wav")
  fileWriter.Write(fileData)

  request, _ := http.NewRequest(http.MethodPost, "/transcribe", &buffer)
  request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
  multipartWriter.Close()

  server := server.Create()
  recorder := httptest.NewRecorder()
  server.ServeHTTP(recorder, request)

  var response map[string]string
  err := json.Unmarshal(recorder.Body.Bytes(), &response)

  assert.Nil(t, err)
  assert.Equal(t, recorder.Code, http.StatusOK)
  assert.Equal(t, response["transcription"], "mocked transcription")
}

func TestTranscribeRouteFileExtensionValidation(t *testing.T) {
  var buffer bytes.Buffer
  multipartWriter := multipart.NewWriter(&buffer)
  fileWriter, _ := multipartWriter.CreateFormFile("file", "test.txt")
  fileWriter.Write([]byte("test"))

  request, _ := http.NewRequest(http.MethodPost, "/transcribe", &buffer)
  request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
  multipartWriter.Close()

  server := server.Create()
  recorder := httptest.NewRecorder()
  server.ServeHTTP(recorder, request)

  var response map[string]string
  err := json.Unmarshal(recorder.Body.Bytes(), &response)

  assert.Nil(t, err)
  assert.Equal(t, http.StatusBadRequest, recorder.Code)
  assert.Equal(t, "invalid audio file extension", response["error"])
}

