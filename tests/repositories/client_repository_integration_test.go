// src/repositories/client_repository_integration_test.go
package test

import (
	"banking/src/models"
	"banking/src/repositories"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite para Go
	"github.com/stretchr/testify/assert"
)

func TestClientRepository_CreateClient(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewClientRepository(db)
	client := &models.Client{
		Name:       "John Doe",
		AccountNum: "123456",
		Balance:    1000.0,
	}

	err := repo.CreateClient(client)
	assert.NoError(t, err)

	// Verifica se o cliente foi realmente criado
	storedClient, err := repo.GetClientByAccountNum("123456")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", storedClient.Name)
	assert.Equal(t, "123456", storedClient.AccountNum)
	assert.Equal(t, 1000.0, storedClient.Balance)
}

func TestClientRepository_GetClientByAccountNum_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewClientRepository(db)

	client, err := repo.GetClientByAccountNum("999999")
	assert.Error(t, err)
	assert.Nil(t, client)
	assert.EqualError(t, err, "client not found")
}

func TestClientRepository_UpdateClientBalance(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewClientRepository(db)

	// Cria um cliente para o teste
	client := &models.Client{
		Name:       "Jane Doe",
		AccountNum: "654321",
		Balance:    200.0,
	}
	err := repo.CreateClient(client)
	assert.NoError(t, err)

	// Atualiza o saldo
	client.Balance = 500.0
	err = repo.UpdateClientBalance(client)
	assert.NoError(t, err)

	// Verifica se o saldo foi atualizado corretamente
	updatedClient, err := repo.GetClientByAccountNum("654321")
	assert.NoError(t, err)
	assert.Equal(t, 500.0, updatedClient.Balance)
}

func TestClientRepository_GetClients(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repositories.NewClientRepository(db)

	// Insere alguns clientes
	clients := []models.Client{
		{Name: "Alice", AccountNum: "111111", Balance: 1000.0},
		{Name: "Bob", AccountNum: "222222", Balance: 1500.0},
	}
	for _, client := range clients {
		err := repo.CreateClient(&client)
		assert.NoError(t, err)
	}

	// Verifica se GetClients retorna os clientes criados
	storedClients, err := repo.GetClients()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(storedClients))
	assert.Equal(t, "Alice", storedClients[0].Name)
	assert.Equal(t, "Bob", storedClients[1].Name)
}
