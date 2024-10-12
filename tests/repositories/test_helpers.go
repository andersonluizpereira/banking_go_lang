// src/repositories/test_helpers.go
package test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB configura um banco de dados SQLite em memória para testes
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:") // Banco de dados em memória
	if err != nil {
		t.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	// Cria as tabelas `clients` e `transfers`
	createClientsTable := `
    CREATE TABLE clients (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        account_num TEXT UNIQUE,
        balance REAL
    );`
	_, err = db.Exec(createClientsTable)
	if err != nil {
		t.Fatalf("Erro ao criar a tabela clients: %v", err)
	}

	createTransfersTable := `
    CREATE TABLE transfers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        from_account_num TEXT,
        to_account_num TEXT,
        amount REAL,
        status TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createTransfersTable)
	if err != nil {
		t.Fatalf("Erro ao criar a tabela transfers: %v", err)
	}

	return db
}
