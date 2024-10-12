// src/models/transfer_test.go
package test

import (
	"banking/src/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	createdAt := time.Now()
	transfer := models.Transfer{
		ID:             1,
		FromAccountNum: "123456",
		ToAccountNum:   "654321",
		Amount:         500.0,
		Status:         "success",
		CreatedAt:      createdAt,
	}

	assert.Equal(t, 1, transfer.ID)
	assert.Equal(t, "123456", transfer.FromAccountNum)
	assert.Equal(t, "654321", transfer.ToAccountNum)
	assert.Equal(t, 500.0, transfer.Amount)
	assert.Equal(t, "success", transfer.Status)
	assert.Equal(t, createdAt, transfer.CreatedAt)
}

func TestTransfer_InvalidData(t *testing.T) {
	invalidTime := time.Time{} // Tempo zero inválido
	transfer := models.Transfer{
		ID:             0,
		FromAccountNum: "",     // Conta de origem inválida
		ToAccountNum:   "",     // Conta de destino inválida
		Amount:         -100.0, // Valor negativo
		Status:         "",     // Status vazio
		CreatedAt:      invalidTime,
	}

	assert.Equal(t, 0, transfer.ID)
	assert.Equal(t, "", transfer.FromAccountNum)
	assert.Equal(t, "", transfer.ToAccountNum)
	assert.Equal(t, -100.0, transfer.Amount)
	assert.Equal(t, "", transfer.Status)
	assert.Equal(t, invalidTime, transfer.CreatedAt)
}

func TestTransfer_UpdateAmountAndStatus(t *testing.T) {
	createdAt := time.Now()
	transfer := models.Transfer{
		ID:             2,
		FromAccountNum: "123456",
		ToAccountNum:   "654321",
		Amount:         300.0,
		Status:         "pending",
		CreatedAt:      createdAt,
	}

	// Atualizando o valor e o status
	transfer.Amount = 400.0
	transfer.Status = "failed"

	assert.Equal(t, 400.0, transfer.Amount)
	assert.Equal(t, "failed", transfer.Status)
}
