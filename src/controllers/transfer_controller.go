package controllers

import (
	"banking/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TransferController define o controlador para as operações de transferência
type TransferController struct {
	TransferService services.TransferServiceInterface
}

// NewTransferController cria uma nova instância de TransferController
func NewTransferController(transferService services.TransferServiceInterface) *TransferController {
	return &TransferController{TransferService: transferService}
}

// TransferFunds realiza uma transferência entre contas
// @Summary Realiza uma transferência
// @Description Realiza uma transferência entre duas contas fornecidas
// @Tags transfers
// @Accept json
// @Produce json
// @Param transferRequest body TransferRequest true "Dados da Transferência"
// @Success 200 {object} map[string]interface{} "Transferência realizada com sucesso"
// @Failure 400 {object} map[string]interface{} "Mensagem de erro"
// @Router /v1/transfer [post]
func (tc *TransferController) TransferFunds(c *gin.Context) {
	var transferRequest TransferRequest
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TransferService.TransferFunds(transferRequest.FromAccount, transferRequest.ToAccount, transferRequest.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "transfer successful"})
}

// GetTransferHistory obtém o histórico de transferências de uma conta
// @Summary Obtém histórico de transferências
// @Description Retorna o histórico de transferências associado a uma conta fornecida
// @Tags transfers
// @Produce json
// @Param accountNum path string true "Número da conta"
// @Success 200 {array} models.Transfer
// @Failure 500 {object} map[string]interface{} "Mensagem de erro"
// @Router /v1/transfers/{accountNum} [get]
func (tc *TransferController) GetTransferHistory(c *gin.Context) {
	accountNum := c.Param("accountNum")
	transfers, err := tc.TransferService.GetTransferHistory(accountNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transfers)
}

// TransferRequest representa o corpo da requisição de transferência
type TransferRequest struct {
	FromAccount string  `json:"from_account" example:"123456"`
	ToAccount   string  `json:"to_account" example:"654321"`
	Amount      float64 `json:"amount" example:"100.50"`
}

// InitTransferRoutes inicializa as rotas de transferência
func InitTransferRoutes(r *gin.Engine, transferService services.TransferServiceInterface) {
	transferController := NewTransferController(transferService)

	v1 := r.Group("/v1")
	{
		v1.POST("/transfer", transferController.TransferFunds)
		v1.GET("/transfers/:accountNum", transferController.GetTransferHistory)
	}
}
