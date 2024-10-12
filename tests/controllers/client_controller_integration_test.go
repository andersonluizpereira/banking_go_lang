// src/controllers/client_controller_integration_test.go
package controllers

import (
	"banking/src/controllers"
	"banking/src/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientService struct {
	mock.Mock
}

func (m *MockClientService) CreateClient(client *models.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientService) GetClients() ([]models.Client, error) {
	args := m.Called()
	return args.Get(0).([]models.Client), args.Error(1)
}

func setupRouter(mockService *MockClientService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	controllers.InitRoutes(r, mockService)
	return r
}

func TestCreateClient_Success(t *testing.T) {
	mockService := new(MockClientService)
	router := setupRouter(mockService)

	client := models.Client{Name: "John Doe", AccountNum: "123456", Balance: 1000.0}
	mockService.On("CreateClient", &client).Return(nil)

	clientJSON, _ := json.Marshal(client)
	req, _ := http.NewRequest("POST", "/v1/clients", bytes.NewBuffer(clientJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseClient models.Client
	err := json.Unmarshal(w.Body.Bytes(), &responseClient)
	assert.NoError(t, err)
	assert.Equal(t, client.Name, responseClient.Name)
	assert.Equal(t, client.AccountNum, responseClient.AccountNum)
	assert.Equal(t, client.Balance, responseClient.Balance)

	mockService.AssertExpectations(t)
}

func TestCreateClient_BadRequest(t *testing.T) {
	mockService := new(MockClientService)
	router := setupRouter(mockService)

	req, _ := http.NewRequest("POST", "/v1/clients", bytes.NewBuffer([]byte(`{invalid_json}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid character")

	mockService.AssertNotCalled(t, "CreateClient")
}

func TestGetClients_Success(t *testing.T) {
	mockService := new(MockClientService)
	router := setupRouter(mockService)

	clients := []models.Client{
		{Name: "John Doe", AccountNum: "123456", Balance: 1000.0},
		{Name: "Jane Doe", AccountNum: "654321", Balance: 2000.0},
	}
	mockService.On("GetClients").Return(clients, nil)

	req, _ := http.NewRequest("GET", "/v1/clients", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseClients []models.Client
	err := json.Unmarshal(w.Body.Bytes(), &responseClients)
	assert.NoError(t, err)
	assert.Equal(t, len(clients), len(responseClients))

	for i, client := range clients {
		assert.Equal(t, client.Name, responseClients[i].Name)
		assert.Equal(t, client.AccountNum, responseClients[i].AccountNum)
		assert.Equal(t, client.Balance, responseClients[i].Balance)
	}

	mockService.AssertExpectations(t)
}

func TestGetClients_InternalServerError(t *testing.T) {
	mockService := new(MockClientService)
	router := setupRouter(mockService)

	mockService.On("GetClients").Return(nil, assert.AnError)

	req, _ := http.NewRequest("GET", "/v1/clients", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}
