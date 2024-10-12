// src/services/transfer_service.go
package services

import (
	"banking/src/models"
	"banking/src/repositories"
	"errors"
	"sync"
)

// TransferService é o serviço que gerencia as transferências entre contas
type TransferService struct {
	clientRepo    repositories.ClientRepository
	transferRepo  repositories.TransferRepository
	transferMutex sync.Mutex
}

// NewTransferService cria uma nova instância de TransferService
func NewTransferService(clientRepo repositories.ClientRepository, transferRepo repositories.TransferRepository) *TransferService {
	return &TransferService{
		clientRepo:   clientRepo,
		transferRepo: transferRepo,
	}
}

// TransferFunds realiza uma transferência entre duas contas
func (s *TransferService) TransferFunds(fromAccountNum, toAccountNum string, amount float64) error {
	if amount <= 0 || amount > 10000 {
		return errors.New("amount must be between 0 and 10,000")
	}

	s.transferMutex.Lock()
	defer s.transferMutex.Unlock()

	fromClient, err := s.clientRepo.GetClientByAccountNum(fromAccountNum)
	if err != nil {
		return err
	}

	if fromClient.Balance < amount {
		return errors.New("insufficient balance")
	}

	toClient, err := s.clientRepo.GetClientByAccountNum(toAccountNum)
	if err != nil {
		return err
	}

	fromClient.Balance -= amount
	toClient.Balance += amount

	err = s.clientRepo.UpdateClientBalance(fromClient)
	if err != nil {
		return err
	}

	err = s.clientRepo.UpdateClientBalance(toClient)
	if err != nil {
		return err
	}

	transfer := models.Transfer{
		FromAccountNum: fromAccountNum,
		ToAccountNum:   toAccountNum,
		Amount:         amount,
		Status:         "success",
	}
	return s.transferRepo.CreateTransfer(&transfer)
}

// GetTransferHistory retorna o histórico de transferências de uma conta específica
func (s *TransferService) GetTransferHistory(accountNum string) ([]models.Transfer, error) {
	return s.transferRepo.GetTransfersByAccountNum(accountNum)
}
