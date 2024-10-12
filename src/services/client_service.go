package services

import (
	"banking/src/models"
	"banking/src/repositories"
	"errors"
)

type ClientService struct {
	repo *repositories.ClientRepository
}

func NewClientService(repo *repositories.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	if client.Name == "" || client.AccountNum == "" {
		return errors.New("missing required fields")
	}
	return s.repo.CreateClient(client)
}

func (s *ClientService) GetClients() ([]models.Client, error) {
	return s.repo.GetClients(), nil
}

func (s *ClientService) GetClientByAccountNum(accountNum string) (*models.Client, error) {
	return s.repo.GetClientByAccountNum(accountNum)
}
