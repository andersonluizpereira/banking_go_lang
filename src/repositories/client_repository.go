package repositories

import (
	"banking/src/models"
	"database/sql"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (repo *ClientRepository) CreateClient(client *models.Client) error {
	_, err := repo.db.Exec("INSERT INTO clients (name, account_num, balance) VALUES (?, ?, ?)",
		client.Name, client.AccountNum, client.Balance)
	return err
}

func (repo *ClientRepository) GetClients() []models.Client {
	rows, _ := repo.db.Query("SELECT id, name, account_num, balance FROM clients")
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		rows.Scan(&client.ID, &client.Name, &client.AccountNum, &client.Balance)
		clients = append(clients, client)
	}
	return clients
}

func (repo *ClientRepository) GetClientByAccountNum(accountNum string) (*models.Client, error) {
	var client models.Client
	err := repo.db.QueryRow("SELECT id, name, account_num, balance FROM clients WHERE account_num = ?", accountNum).
		Scan(&client.ID, &client.Name, &client.AccountNum, &client.Balance)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
