package controllers

import (
	"banking/src/models"
	"banking/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, clientService *services.ClientService) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, clients)
	})
}
