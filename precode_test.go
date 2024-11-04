package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"fmt"
	"strings"
)
type CafeClient struct {
	client *http.Client
	baseURL string
}
var cafeClient CafeClient

func init() {
	server := httptest.NewServer(http.HandlerFunc(mainHandle))
	cafeClient = CafeClient{server.Client(), server.URL}
}
func TestCorrectRequest(t *testing.T){
	response := cafeClient.sentRequest("/cafe?count=4&city=moscow")

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	require.Equal(t, 200, response.StatusCode)
	assert.NotEmpty(t, body)
}
func TestWrongCity(t *testing.T){
	response := cafeClient.sentRequest("/cafe?count=4&city=Sevastopol")
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	errorMessage := string(body)

	require.Equal(t, 400, response.StatusCode)
	require.Equal(t, errorMessage, "wrong city value")
}

func TestMoreThanExistCount(t *testing.T) {
	response := cafeClient.sentRequest("/cafe?count=10&city=moscow")
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

    listResponse := strings.Split(string(body), ",")
	cafeListMoscow:= cafeList["moscow"]

	require.Equal(t, 200, response.StatusCode)
	require.Len(t, cafeListMoscow, len(listResponse))
}

func (c *CafeClient) sentRequest(endpoint string) *http.Response{
	request, _ := http.NewRequest("GET", c.baseURL + endpoint, nil)
	response, err := c.client.Do(request)
	fmt.Println(err)
	return response
}
