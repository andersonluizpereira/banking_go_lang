package main

import (
	"banking/src/controllers"
	"banking/src/database"
	"banking/src/repositories"
	"banking/src/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

	rootCmd.AddCommand(runCmd)

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

	fmt.Println("Running on localhost:8080")
	r.Run(":8080")
}
