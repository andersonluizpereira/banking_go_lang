// src/services/mock_repositories.go
package test

import (
	"banking/src/models"

	"github.com/stretchr/testify/mock"
)

// Definindo MockClientRepository uma vez neste arquivo
type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) GetClientByAccountNum(accountNum string) (*models.Client, error) {
	args := m.Called(accountNum)
	return args.Get(0).(*models.Client), args.Error(1)
}

func (m *MockClientRepository) UpdateClientBalance(client *models.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) CreateClient(client *models.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) GetClients() ([]models.Client, error) {
	args := m.Called()
	return args.Get(0).([]models.Client), args.Error(1)
}

// Definindo MockTransferRepository uma vez neste arquivo
type MockTransferRepository struct {
	mock.Mock
}

func (m *MockTransferRepository) CreateTransfer(transfer *models.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func (m *MockTransferRepository) GetTransfersByAccountNum(accountNum string) ([]models.Transfer, error) {
	args := m.Called(accountNum)
	return args.Get(0).([]models.Transfer), args.Error(1)
}
