package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB agora retorna um erro
func InitDB(db_name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		return nil, err
	}
	return db, nil
}
