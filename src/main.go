package main

import (
	"banking/src/controllers"
	"banking/src/database"
	"banking/src/repositories"
	"banking/src/services"
	"fmt"
	"os"

	_ "banking/docs" // Importa a documentação gerada pelo Swag

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "bankingapp",
		Short: "Banking App CLI",
		Long:  "This is a CLI application for managing a banking application server.",
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the banking server",
		Long:  "Starts the banking server on localhost:8080",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}

	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Long:  "Applies the latest database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runMigrations("./bank.db"); err != nil {
				fmt.Println("Failed to run migrations:", err)
				os.Exit(1)
			}
			fmt.Println("Migrations applied successfully")
		},
	}

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runServer() {
	r := gin.Default()
	db, err := database.InitDB("./bank.db")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}
	defer db.Close()

	clientRepo := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo)

	transferRepo := repositories.NewTransferRepository(db)
	transferService := services.NewTransferService(clientRepo, transferRepo)

	controllers.InitRoutes(r, clientService)
	controllers.InitTransferRoutes(r, transferService)

	// Rota Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Running on localhost:8080")
	r.Run(":8080")
}

func runMigrations(dbPath string) error {
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("sqlite3://%s", dbPath),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
