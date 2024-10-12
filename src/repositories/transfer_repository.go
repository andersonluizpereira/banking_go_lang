package repositories

import (
	"banking/src/models"
	"database/sql"
)

type TransferRepository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{db: db}
}

func (repo *TransferRepository) CreateTransfer(transfer *models.Transfer) error {
	_, err := repo.db.Exec("INSERT INTO transfers (from_account_num, to_account_num, amount, status) VALUES (?, ?, ?, ?)",
		transfer.FromAccountNum, transfer.ToAccountNum, transfer.Amount, transfer.Status)
	return err
}

func (repo *TransferRepository) GetTransfersByAccountNum(accountNum string) ([]models.Transfer, error) {
	rows, err := repo.db.Query("SELECT from_account_num, to_account_num, amount, status FROM transfers WHERE from_account_num = ? OR to_account_num = ? ORDER BY created_at DESC", accountNum, accountNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		rows.Scan(&transfer.FromAccountNum, &transfer.ToAccountNum, &transfer.Amount, &transfer.Status)
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}

func (repo *ClientRepository) UpdateClientBalance(client *models.Client) error {
	_, err := repo.db.Exec("UPDATE clients SET balance = ? WHERE account_num = ?", client.Balance, client.AccountNum)
	return err
}
