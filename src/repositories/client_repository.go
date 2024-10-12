package repositories

import (
	"banking/src/models"
	"database/sql"
	"errors"
)

type ClientRepository interface {
	GetClientByAccountNum(accountNum string) (*models.Client, error)
	UpdateClientBalance(client *models.Client) error
	CreateClient(client *models.Client) error
	GetClients() ([]models.Client, error)
}

// ClientRepositoryImpl é a implementação concreta do repositório
type ClientRepositoryImpl struct {
	db *sql.DB // Supondo que você tenha uma conexão de banco de dados SQL
}

// NewClientRepository cria uma nova instância de ClientRepositoryImpl
func NewClientRepository(db *sql.DB) *ClientRepositoryImpl {
	return &ClientRepositoryImpl{db: db}
}

// Implementação do método GetClientByAccountNum
func (repo *ClientRepositoryImpl) GetClientByAccountNum(accountNum string) (*models.Client, error) {
	var client models.Client
	err := repo.db.QueryRow("SELECT id, name, account_num, balance FROM clients WHERE account_num = ?", accountNum).
		Scan(&client.ID, &client.Name, &client.AccountNum, &client.Balance)
	if err == sql.ErrNoRows {
		return nil, errors.New("client not found")
	} else if err != nil {
		return nil, err
	}
	return &client, nil
}

// Implementação do método UpdateClientBalance
func (repo *ClientRepositoryImpl) UpdateClientBalance(client *models.Client) error {
	_, err := repo.db.Exec("UPDATE clients SET balance = ? WHERE account_num = ?", client.Balance, client.AccountNum)
	return err
}

// Implementação do método CreateClient
func (repo *ClientRepositoryImpl) CreateClient(client *models.Client) error {
	_, err := repo.db.Exec("INSERT INTO clients (name, account_num, balance) VALUES (?, ?, ?)",
		client.Name, client.AccountNum, client.Balance)
	return err
}

// Implementação do método GetClients
func (repo *ClientRepositoryImpl) GetClients() ([]models.Client, error) {
	rows, err := repo.db.Query("SELECT id, name, account_num, balance FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.AccountNum, &client.Balance); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
