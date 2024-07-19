package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/DenisRukavitsa/voice-notes/server"
)

func SendRequest(method string, route string, body []byte) (int, map[string]string, error) {
	server := server.Create()
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest(method, route, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	server.ServeHTTP(recorder, request)

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	return recorder.Code, response, err
}