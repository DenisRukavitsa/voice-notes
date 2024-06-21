package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/server"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	server := server.Create()
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	server.ServeHTTP(recorder, request)

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "server healthy", response["message"])
}
