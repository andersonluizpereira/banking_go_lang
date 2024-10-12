// test/client_service_test.go
package test

import (
	"banking/src/models"
	"banking/src/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientService(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	assert.NotNil(t, clientService, "Expected NewClientService to return a non-nil ClientService instance")
}

func TestCreateClient_Success(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	client := &models.Client{Name: "John Doe", AccountNum: "123456"}

	mockRepo.On("CreateClient", client).Return(nil)

	err := clientService.CreateClient(client)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateClient_MissingFields(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	client := &models.Client{Name: "", AccountNum: "123456"}

	err := clientService.CreateClient(client)

	assert.Error(t, err)
	assert.EqualError(t, err, "missing required fields")
	mockRepo.AssertNotCalled(t, "CreateClient", client)
}

func TestGetClients_Success(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	clients := []models.Client{
		{Name: "John Doe", AccountNum: "123456", Balance: 100.0},
		{Name: "Jane Doe", AccountNum: "654321", Balance: 200.0},
	}

	mockRepo.On("GetClients").Return(clients, nil)

	result, err := clientService.GetClients()

	assert.NoError(t, err)
	assert.Equal(t, clients, result)
	mockRepo.AssertExpectations(t)
}

func TestGetClientByAccountNum_Success(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	client := &models.Client{Name: "John Doe", AccountNum: "123456", Balance: 100.0}

	mockRepo.On("GetClientByAccountNum", "123456").Return(client, nil)

	result, err := clientService.GetClientByAccountNum("123456")

	assert.NoError(t, err)
	assert.Equal(t, client, result)
	mockRepo.AssertExpectations(t)
}

func TestGetClientByAccountNum_NotFound(t *testing.T) {
	mockRepo := new(MockClientRepository)
	clientService := services.NewClientService(mockRepo)

	mockRepo.On("GetClientByAccountNum", "999999").Return((*models.Client)(nil), errors.New("client not found"))

	result, err := clientService.GetClientByAccountNum("999999")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "client not found")
	mockRepo.AssertExpectations(t)
}
