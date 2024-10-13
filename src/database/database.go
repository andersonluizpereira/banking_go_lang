package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(db_name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		return nil, err
	}

	// Chama a função para criar a tabela clients
	err = createClientsTable(db)
	if err != nil {
		return nil, err
	}

	// Chama a função para criar a tabela transfers
	err = createTransfersTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createClientsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		account_num TEXT NOT NULL UNIQUE,
		balance REAL NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating clients table: %v", err)
		return err
	}
	return nil
}

func createTransfersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS transfers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		from_account_num TEXT NOT NULL,
		to_account_num TEXT NOT NULL,
		amount REAL NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (from_account_num) REFERENCES clients(account_num),
		FOREIGN KEY (to_account_num) REFERENCES clients(account_num)
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating transfers table: %v", err)
		return err
	}
	return nil
}
