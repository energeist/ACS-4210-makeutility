package helpers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadFromDotEnv tests the LoadFromDotEnv function
func TestLoadFromDotEnv(t *testing.T) {
	// Setup the .env file for testing
	os.WriteFile("../.env", []byte("TEST_KEY=success"), 0644)
	defer os.Remove("../.env")

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"Valid key", "TEST_KEY", "success"},
		{"Invalid key", "MISSING_KEY", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LoadFromDotEnv(tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestAligulacURL tests the AligulacURL function
func TestAligulacURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		apiKey   string
		id       int
		expected string
	}{
		{"Normal case", "player", "myapikey", 123, "https://api.aligulac.com/api/v1/player/123/?apikey=myapikey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := AligulacURL(tt.endpoint, tt.apiKey, tt.id)
			assert.Equal(t, tt.expected, url)
		})
	}
}

// TestServerURL tests the ServerURL function
func TestServerURL(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		serverPort string
		expected   string
	}{
		{"Default port", "api/data", "8080", "http://localhost:8080/api/data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := ServerURL(tt.endpoint, tt.serverPort)
			assert.Equal(t, tt.expected, url)
		})
	}
}

// TestGetRequest tests the GetRequest function
func TestGetRequest(t *testing.T) {
	tests := []struct {
		name         string
		serverStatus int
		serverBody   string
		expectError  bool
		expectedBody string
	}{
		{"Success", http.StatusOK, `{"message": "success"}`, false, `{"message": "success"}`},
		{"Server Error", http.StatusInternalServerError, "", true, ""},
		{"Client Error", http.StatusNotFound, "", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup a test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverBody))
			}))
			defer server.Close()

			// Call the function under test
			body, err := GetRequest(server.URL)

			// Assert expectations
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, string(body))
			}
		})
	}
}

// TestPostRequest tests the PostRequest function
func TestPostRequest(t *testing.T) {
	tests := []struct {
		name         string
		serverStatus int
		requestBody  interface{}
		responseBody string
		expectError  bool
	}{
		{"Success", http.StatusOK, map[string]string{"key": "value"}, "OK", false},
		{"Server Error", http.StatusInternalServerError, map[string]string{"key": "value"}, "", true},
		{"Invalid Data", http.StatusOK, make(chan int), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup a test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatus)
				body, _ := io.ReadAll(r.Body)
				assert.JSONEq(t, `{"key":"value"}`, string(body))
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			// Call the function under test
			resp, err := PostRequest(server.URL, tt.requestBody)

			// Assert expectations
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				responseBody, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.responseBody, string(responseBody))
			}
		})
	}
}
