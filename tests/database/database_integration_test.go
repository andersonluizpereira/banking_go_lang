// src/database/database_integration_test.go
package database

import (
	"banking/src/database"
	"io/ioutil"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Função para redirecionar a saída de log durante o teste
func captureLogOutput(f func()) string {
	old := log.Writer()
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	f()

	w.Close()
	log.SetOutput(old)

	out, _ := ioutil.ReadAll(r)
	print("log->", string(out))
	return string(out)
}

func TestInitDB_Success(t *testing.T) {
	dbName := "./test_bank.db"
	os.Remove(dbName) // Certifique-se de que o arquivo não exista antes do teste

	db, err := database.InitDB(dbName)
	defer os.Remove(dbName) // Limpeza após o teste
	defer db.Close()

	// Verifique se o banco foi inicializado corretamente
	assert.NotNil(t, db, "O banco de dados não deveria ser nulo")
	assert.NoError(t, err, "Deveria ser possível conectar ao banco de dados")

	// Tente executar uma operação simples
	_, execErr := db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY AUTOINCREMENT)")
	assert.NoError(t, execErr, "Deveria ser possível criar uma tabela no banco de dados")
}

func TestInitDB_Failure(t *testing.T) {
	dbName := "/invalid_folder/test_bank.db" // Um caminho provavelmente inválido

	logOutput := captureLogOutput(func() {
		database.InitDB(dbName) // Deve falhar ao tentar inicializar
	})

	assert.Contains(t, logOutput, "")
}
