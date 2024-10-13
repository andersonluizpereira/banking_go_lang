// src/services/client_service.go
package services

import (
	"banking/src/models"
	"banking/src/repositories"
	"errors"
)

// ClientServiceInterface define a interface para operações do cliente
type ClientServiceInterface interface {
	CreateClient(client *models.Client) error
	GetClients() ([]models.Client, error)
	GetClientByAccountNum(accountNum string) (*models.Client, error)
}

// ClientService é a implementação concreta que atende a ClientServiceInterface
type ClientService struct {
	repo repositories.ClientRepository // Interface do repositório de cliente
}

// Certifique-se de que ClientService implementa ClientServiceInterface
var _ ClientServiceInterface = (*ClientService)(nil)

// NewClientService cria uma nova instância de ClientService
func NewClientService(repo repositories.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

// CreateClient cria um novo cliente, verificando os campos necessários
func (s *ClientService) CreateClient(client *models.Client) error {
	if client.Name == "" || client.AccountNum == "" {
		return errors.New("missing required fields")
	}
	return s.repo.CreateClient(client)
}

// GetClients retorna todos os clientes
func (s *ClientService) GetClients() ([]models.Client, error) {
	return s.repo.GetClients()
}

func (s *ClientService) GetClientByAccountNum(accountNum string) (*models.Client, error) {
	return s.repo.GetClientByAccountNum(accountNum)
}
