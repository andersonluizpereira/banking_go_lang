package repositories

import (
	"banking/src/models"
	"database/sql"
)

// TransferRepository define a interface para operações de transferência
type TransferRepository interface {
	CreateTransfer(transfer *models.Transfer) error
	GetTransfersByAccountNum(accountNum string) ([]models.Transfer, error)
}

type TransferRepositoryImpl struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepositoryImpl {
	return &TransferRepositoryImpl{db: db}
}

// Implementação do método CreateTransfer
func (repo *TransferRepositoryImpl) CreateTransfer(transfer *models.Transfer) error {
	_, err := repo.db.Exec("INSERT INTO transfers (from_account_num, to_account_num, amount, status) VALUES (?, ?, ?, ?)",
		transfer.FromAccountNum, transfer.ToAccountNum, transfer.Amount, transfer.Status)
	return err
}

// Implementação do método GetTransfersByAccountNum
func (repo *TransferRepositoryImpl) GetTransfersByAccountNum(accountNum string) ([]models.Transfer, error) {
	rows, err := repo.db.Query("SELECT from_account_num, to_account_num, amount, status FROM transfers WHERE from_account_num = ? OR to_account_num = ? ORDER BY created_at DESC", accountNum, accountNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		if err := rows.Scan(&transfer.FromAccountNum, &transfer.ToAccountNum, &transfer.Amount, &transfer.Status); err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
