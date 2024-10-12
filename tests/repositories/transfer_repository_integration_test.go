// src/repositories/transfer_repository_integration_test.go
package test

import (
	"banking/src/models"
	"banking/src/repositories"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestTransferRepository_CreateTransfer(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewTransferRepository(db)
	transfer := &models.Transfer{
		FromAccountNum: "123456",
		ToAccountNum:   "654321",
		Amount:         100.0,
		Status:         "success",
		CreatedAt:      time.Now(),
	}

	err := repo.CreateTransfer(transfer)
	assert.NoError(t, err)

	// Verifica se a transferência foi realmente criada
	rows, err := db.Query("SELECT from_account_num, to_account_num, amount, status FROM transfers WHERE from_account_num = ?", "123456")
	assert.NoError(t, err)
	defer rows.Close()

	assert.True(t, rows.Next())
	var fromAccountNum, toAccountNum, status string
	var amount float64
	err = rows.Scan(&fromAccountNum, &toAccountNum, &amount, &status)
	assert.NoError(t, err)
	assert.Equal(t, "123456", fromAccountNum)
	assert.Equal(t, "654321", toAccountNum)
	assert.Equal(t, 100.0, amount)
	assert.Equal(t, "success", status)
}

func TestTransferRepository_GetTransfersByAccountNum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewTransferRepository(db)

	// Insere transferências para teste
	transfers := []models.Transfer{
		{FromAccountNum: "123456", ToAccountNum: "654321", Amount: 50.0, Status: "success", CreatedAt: time.Now()},
		{FromAccountNum: "654321", ToAccountNum: "123456", Amount: 75.0, Status: "failed", CreatedAt: time.Now()},
	}
	for _, transfer := range transfers {
		err := repo.CreateTransfer(&transfer)
		assert.NoError(t, err)
	}

	// Testa GetTransfersByAccountNum
	storedTransfers, err := repo.GetTransfersByAccountNum("123456")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(storedTransfers))

	// Verifica o conteúdo do resultado
	assert.Equal(t, "123456", storedTransfers[0].FromAccountNum)
	assert.Equal(t, "654321", storedTransfers[1].FromAccountNum)
	assert.Equal(t, 75.0, storedTransfers[1].Amount)
	assert.Equal(t, "failed", storedTransfers[1].Status)
}
