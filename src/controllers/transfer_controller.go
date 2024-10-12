package controllers

import (
	"banking/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitTransferRoutes(r *gin.Engine, transferService *services.TransferService) {
	r.POST("/v1/transfer", func(c *gin.Context) {
		var transferRequest struct {
			FromAccount string  `json:"from_account"`
			ToAccount   string  `json:"to_account"`
			Amount      float64 `json:"amount"`
		}

		if err := c.ShouldBindJSON(&transferRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := transferService.TransferFunds(transferRequest.FromAccount, transferRequest.ToAccount, transferRequest.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "transfer successful"})
	})

	r.GET("/v1/transfers/:accountNum", func(c *gin.Context) {
		accountNum := c.Param("accountNum")
		transfers, err := transferService.GetTransferHistory(accountNum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, transfers)
	})
}
