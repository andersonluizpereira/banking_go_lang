package controllers

import (
	"banking/src/models"
	"banking/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, clientService services.ClientServiceInterface) {
	r.POST("/v1/clients", func(c *gin.Context) {
		var client models.Client
		if err := c.ShouldBindJSON(&client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := clientService.CreateClient(&client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, client)
	})

	r.GET("/v1/clients", func(c *gin.Context) {
		clients, err := clientService.GetClients()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Certifique-se de retornar o erro
			return
		}
		c.JSON(http.StatusOK, clients)
	})
	r.GET("/v1/clients/:accountNum", func(c *gin.Context) {
		accountNum := c.Param("accountNum")
		client, err := clientService.GetClientByAccountNum(accountNum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
			return
		}
		c.JSON(http.StatusOK, client)
	})
}
