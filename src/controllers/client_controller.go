package controllers

import (
	"banking/src/models"
	"banking/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ClientController gerencia as rotas de cliente
type ClientController struct {
	ClientService services.ClientServiceInterface
}

// NewClientController cria uma nova instância de ClientController
func NewClientController(clientService services.ClientServiceInterface) *ClientController {
	return &ClientController{ClientService: clientService}
}

// CreateClient cria um novo cliente
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com as informações fornecidas
// @Tags clients
// @Accept json
// @Produce json
// @Param client body models.Client true "Cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Mensagem de erro"
// @Router /v1/clients [post]
func (cc *ClientController) CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.ClientService.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// GetClients lista todos os clientes
// @Summary Lista todos os clientes
// @Description Retorna uma lista de todos os clientes cadastrados
// @Tags clients
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]interface{} "Mensagem de erro"
// @Router /v1/clients [get]
func (cc *ClientController) GetClients(c *gin.Context) {
	clients, err := cc.ClientService.GetClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// GetClientByAccountNum busca cliente por número da conta
// @Summary Busca cliente por número da conta
// @Description Busca um cliente pelo número da conta fornecido
// @Tags clients
// @Produce json
// @Param accountNum path string true "Número da conta"
// @Success 200 {object} models.Client
// @Failure 404 {object} map[string]interface{} "client not found"
// @Router /v1/clients/{accountNum} [get]
func (cc *ClientController) GetClientByAccountNum(c *gin.Context) {
	accountNum := c.Param("accountNum")
	client, err := cc.ClientService.GetClientByAccountNum(accountNum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// InitRoutes inicializa as rotas para o controlador de clientes
func InitRoutes(r *gin.Engine, clientService services.ClientServiceInterface) {
	clientController := NewClientController(clientService)
	v1 := r.Group("/v1")
	{
		v1.POST("/clients", clientController.CreateClient)
		v1.GET("/clients", clientController.GetClients)
		v1.GET("/clients/:accountNum", clientController.GetClientByAccountNum)
	}
}
