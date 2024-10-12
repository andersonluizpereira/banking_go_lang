package test

import (
	"banking/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := models.Client{
		ID:         1,
		Name:       "John Doe",
		AccountNum: "123456",
		Balance:    1000.0,
	}

	assert.Equal(t, 1, client.ID)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "123456", client.AccountNum)
	assert.Equal(t, 1000.0, client.Balance)
}

func TestClient_UpdateBalance(t *testing.T) {
	client := models.Client{
		ID:         2,
		Name:       "Jane Doe",
		AccountNum: "654321",
		Balance:    500.0,
	}

	// Atualizar o saldo
	client.Balance += 250.0
	assert.Equal(t, 750.0, client.Balance)

	client.Balance -= 100.0
	assert.Equal(t, 650.0, client.Balance)
}

func TestClient_InvalidData(t *testing.T) {
	client := models.Client{
		ID:         0,      // ID inválido
		Name:       "",     // Nome inválido
		AccountNum: "",     // Número da conta inválido
		Balance:    -500.0, // Saldo inválido
	}

	assert.Equal(t, 0, client.ID)
	assert.Equal(t, "", client.Name)
	assert.Equal(t, "", client.AccountNum)
	assert.Equal(t, -500.0, client.Balance)
}
