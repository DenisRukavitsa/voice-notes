package tests

import (
	"net/http"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	code, response, err := helpers.SendRequest(http.MethodGet, "/health", nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "server healthy", response["message"])
}
