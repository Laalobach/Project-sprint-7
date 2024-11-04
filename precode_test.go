package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorrectRequest(t *testing.T) {
	response := sentRequest("/cafe?count=4&city=moscow")
	body := response.Body.String()

	require.Equal(t, http.StatusOK, response.Code)
	assert.NotEmpty(t, body)
}
func TestWrongCity(t *testing.T) {
	response := sentRequest("/cafe?count=4&city=Sevastopol")
	body := response.Body.String()
	errorMessage := string(body)

	require.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, errorMessage, "wrong city value")
}

func TestMoreThanExistCount(t *testing.T) {
	response := sentRequest("/cafe?count=10&city=moscow")
	body := response.Body.String()
	totalCount := 4
	listResponse := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, response.Code)
	assert.Len(t, listResponse, totalCount)
}

func sentRequest(endpoint string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", endpoint, nil)

	responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, request)
	
	return responseRecorder
}
