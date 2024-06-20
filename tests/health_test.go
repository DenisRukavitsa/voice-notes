package tests

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/DenisRukavitsa/voice-notes/server"
)

func TestHealthRoute(t *testing.T) {
  server := server.Create()
  recorder := httptest.NewRecorder()

  request, err := http.NewRequest(http.MethodGet, "/health", nil)
  if err != nil {
    t.Fatal(err)
  }
  server.ServeHTTP(recorder, request)

  var response map[string]string
  err = json.Unmarshal(recorder.Body.Bytes(), &response)
  if err != nil {
    t.Fatal(err)
  }

  if recorder.Code != http.StatusOK {
    t.Fatalf("Expected status %d, got %d", http.StatusOK, recorder.Code)
  }
  expectedMessage := "server healthy"
  if response["message"] != expectedMessage {
    t.Fatalf("Expected response %s, got %s", expectedMessage, response["message"])
  }
}
