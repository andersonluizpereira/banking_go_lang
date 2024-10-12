package main

import (
	"banking/src/controllers"
	"banking/src/database"
	"banking/src/repositories"
	"banking/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.InitDB("./bank.db")

	clientRepo := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo)

	transferRepo := repositories.NewTransferRepository(db)

	transferService := services.NewTransferService(
		clientRepo,
		transferRepo,
	)

	controllers.InitRoutes(r, clientService)
	controllers.InitTransferRoutes(r, transferService)

	print("running localhost:8080")
	r.Run(":8080")
}
