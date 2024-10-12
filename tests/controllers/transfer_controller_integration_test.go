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

// MockTransferService implementa a interface TransferServiceInterface para testes
type MockTransferService struct {
	mock.Mock
}

func (m *MockTransferService) TransferFunds(fromAccount, toAccount string, amount float64) error {
	args := m.Called(fromAccount, toAccount, amount)
	return args.Error(0)
}

func (m *MockTransferService) GetTransferHistory(accountNum string) ([]models.Transfer, error) {
	args := m.Called(accountNum)
	return args.Get(0).([]models.Transfer), args.Error(1)
}

func setupRouterTranferIntegration(mockService *MockTransferService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	controllers.InitTransferRoutes(r, mockService) // Passando o mock que implementa TransferServiceInterface
	return r
}
func TestTransferFunds_Success(t *testing.T) {
	mockService := new(MockTransferService)
	router := setupRouterTranferIntegration(mockService)

	transferRequest := map[string]interface{}{
		"from_account": "123456",
		"to_account":   "654321",
		"amount":       100.0,
	}
	mockService.On("TransferFunds", "123456", "654321", 100.0).Return(nil)

	body, _ := json.Marshal(transferRequest)
	req, _ := http.NewRequest("POST", "/v1/transfer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "transfer successful", response["status"])

	mockService.AssertExpectations(t)
}

func TestTransferFunds_BadRequest_InvalidJSON(t *testing.T) {
	mockService := new(MockTransferService)
	router := setupRouterTranferIntegration(mockService)

	req, _ := http.NewRequest("POST", "/v1/transfer", bytes.NewBuffer([]byte(`{invalid_json}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid character")

	mockService.AssertNotCalled(t, "TransferFunds")
}

func TestTransferFunds_FailedTransfer(t *testing.T) {
	mockService := new(MockTransferService)
	router := setupRouterTranferIntegration(mockService)

	transferRequest := map[string]interface{}{
		"from_account": "123456",
		"to_account":   "654321",
		"amount":       10000.0,
	}
	mockService.On("TransferFunds", "123456", "654321", 10000.0).Return(assert.AnError)

	body, _ := json.Marshal(transferRequest)
	req, _ := http.NewRequest("POST", "/v1/transfer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, assert.AnError.Error(), response["error"])

	mockService.AssertExpectations(t)
}

func TestGetTransferHistory_Success(t *testing.T) {
	mockService := new(MockTransferService)
	router := setupRouterTranferIntegration(mockService)

	transfers := []models.Transfer{
		{FromAccountNum: "123456", ToAccountNum: "654321", Amount: 50.0, Status: "success"},
		{FromAccountNum: "654321", ToAccountNum: "123456", Amount: 75.0, Status: "failed"},
	}
	mockService.On("GetTransferHistory", "123456").Return(transfers, nil)

	req, _ := http.NewRequest("GET", "/v1/transfers/123456", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseTransfers []models.Transfer
	err := json.Unmarshal(w.Body.Bytes(), &responseTransfers)
	assert.NoError(t, err)
	assert.Equal(t, len(transfers), len(responseTransfers))

	for i, transfer := range transfers {
		assert.Equal(t, transfer.FromAccountNum, responseTransfers[i].FromAccountNum)
		assert.Equal(t, transfer.ToAccountNum, responseTransfers[i].ToAccountNum)
		assert.Equal(t, transfer.Amount, responseTransfers[i].Amount)
		assert.Equal(t, transfer.Status, responseTransfers[i].Status)
	}

	mockService.AssertExpectations(t)
}

func TestGetTransferHistory_InternalServerError(t *testing.T) {
	mockService := new(MockTransferService)
	router := setupRouterTranferIntegration(mockService)

	mockService.On("GetTransferHistory", "123456").Return(nil, assert.AnError)

	req, _ := http.NewRequest("GET", "/v1/transfers/123456", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
