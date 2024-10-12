// src/services/transfer_service_test.go
package test

import (
	"banking/src/models"
	"banking/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Testes

func TestTransferFunds_Success(t *testing.T) {
	mockClientRepo := new(MockClientRepository)
	mockTransferRepo := new(MockTransferRepository)
	transferService := services.NewTransferService(mockClientRepo, mockTransferRepo)

	fromClient := &models.Client{AccountNum: "123456", Balance: 5000}
	toClient := &models.Client{AccountNum: "654321", Balance: 1000}
	amount := 1000.0

	mockClientRepo.On("GetClientByAccountNum", "123456").Return(fromClient, nil)
	mockClientRepo.On("GetClientByAccountNum", "654321").Return(toClient, nil)
	mockClientRepo.On("UpdateClientBalance", fromClient).Return(nil)
	mockClientRepo.On("UpdateClientBalance", toClient).Return(nil)
	mockTransferRepo.On("CreateTransfer", mock.AnythingOfType("*models.Transfer")).Return(nil)

	err := transferService.TransferFunds("123456", "654321", amount)

	assert.NoError(t, err)
	assert.Equal(t, 4000.0, fromClient.Balance)
	assert.Equal(t, 2000.0, toClient.Balance)
	mockClientRepo.AssertExpectations(t)
	mockTransferRepo.AssertExpectations(t)
}

func TestTransferFunds_InsufficientBalance(t *testing.T) {
	mockClientRepo := new(MockClientRepository)
	mockTransferRepo := new(MockTransferRepository)
	transferService := services.NewTransferService(mockClientRepo, mockTransferRepo)

	fromClient := &models.Client{AccountNum: "123456", Balance: 500}
	toClient := &models.Client{AccountNum: "654321", Balance: 1000}
	amount := 1000.0

	mockClientRepo.On("GetClientByAccountNum", "123456").Return(fromClient, nil)
	mockClientRepo.On("GetClientByAccountNum", "654321").Return(toClient, nil)

	err := transferService.TransferFunds("123456", "654321", amount)

	assert.Error(t, err)
	assert.EqualError(t, err, "insufficient balance")
	mockClientRepo.AssertNotCalled(t, "UpdateClientBalance", fromClient)
	mockClientRepo.AssertNotCalled(t, "UpdateClientBalance", toClient)
	mockTransferRepo.AssertNotCalled(t, "CreateTransfer", mock.Anything)
}

func TestTransferFunds_AmountExceedsLimit(t *testing.T) {
	mockClientRepo := new(MockClientRepository)
	mockTransferRepo := new(MockTransferRepository)
	transferService := services.NewTransferService(mockClientRepo, mockTransferRepo)

	amount := 15000.0 // Excede o limite

	err := transferService.TransferFunds("123456", "654321", amount)

	assert.Error(t, err)
	assert.EqualError(t, err, "amount must be between 0 and 10,000")
	mockClientRepo.AssertNotCalled(t, "GetClientByAccountNum", "123456")
	mockClientRepo.AssertNotCalled(t, "UpdateClientBalance", mock.Anything)
	mockTransferRepo.AssertNotCalled(t, "CreateTransfer", mock.Anything)
}

func TestGetTransferHistory_Success(t *testing.T) {
	mockClientRepo := new(MockClientRepository)
	mockTransferRepo := new(MockTransferRepository)
	transferService := services.NewTransferService(mockClientRepo, mockTransferRepo)

	transfers := []models.Transfer{
		{FromAccountNum: "123456", ToAccountNum: "654321", Amount: 500, Status: "success"},
		{FromAccountNum: "654321", ToAccountNum: "123456", Amount: 300, Status: "success"},
	}

	mockTransferRepo.On("GetTransfersByAccountNum", "123456").Return(transfers, nil)

	result, err := transferService.GetTransferHistory("123456")

	assert.NoError(t, err)
	assert.Equal(t, transfers, result)
	mockTransferRepo.AssertExpectations(t)
}
