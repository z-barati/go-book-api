package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	router := gin.New()
	router.POST("/register", Register)

	// Test cases
	tests := []struct {
		name           string
		payload        RegisterRequest
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "Valid registration",
			payload: RegisterRequest{
				Username:  "testuser",
				Email:     "test@example.com",
				Password:  "password123",
				FirstName: "Test",
				LastName:  "User",
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name: "Invalid email",
			payload: RegisterRequest{
				Username:  "testuser",
				Email:     "invalid-email",
				Password:  "password123",
				FirstName: "Test",
				LastName:  "User",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "Short password",
			payload: RegisterRequest{
				Username:  "testuser",
				Email:     "test@example.com",
				Password:  "123",
				FirstName: "Test",
				LastName:  "User",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Parse response
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, response, "error")
			} else {
				assert.NoError(t, err)
				assert.Contains(t, response, "message")
				assert.Contains(t, response, "token")
			}
		})
	}
}

func TestLogin(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	router := gin.New()
	router.POST("/login", Login)

	// Test cases
	tests := []struct {
		name           string
		payload        LoginRequest
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "Valid login",
			payload: LoginRequest{
				Username: "testuser",
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name: "Missing username",
			payload: LoginRequest{
				Password: "password123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "Missing password",
			payload: LoginRequest{
				Username: "testuser",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Parse response
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, response, "error")
			} else {
				assert.NoError(t, err)
				assert.Contains(t, response, "message")
				assert.Contains(t, response, "token")
			}
		})
	}
} 