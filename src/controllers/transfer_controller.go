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
# Banking API

Esta é uma API de exemplo para um sistema bancário, desenvolvida em Go utilizando o framework Gin. A API permite a criação de clientes, realização de transferências entre contas e consulta de histórico de transferências. A documentação foi gerada automaticamente com Swagger.

## Sumário
- [Tecnologias](#tecnologias)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Configuração](#instalação-e-configuração)
- [Endpoints da API](#endpoints-da-api)
- [Documentação Swagger](#documentação-swagger)
- [Comandos Úteis](#comandos-úteis)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Exemplo de Uso](#exemplo-de-uso)

## Tecnologias

- Go
- Gin
- SQLite
- Swagger para documentação

## Pré-requisitos

- [Go](https://golang.org/doc/install) versão 1.16 ou superior
- [Docker](https://docs.docker.com/get-docker/) (opcional, para rodar com Docker)
- [Swag CLI](https://github.com/swaggo/swag) para gerar a documentação Swagger
    ```bash
    go get -u github.com/swaggo/swag/cmd/swag
    ```

## Instalação e Configuração

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/banking-api.git
    cd banking-api
    ```

2. Inicialize o módulo Go (se ainda não estiver inicializado):

    ```bash
    go mod init banking
    go mod tidy
    ```

3. Gere a documentação Swagger:

    ```bash
    swag init -g src/main.go
    ```

4. Construa o banco de dados SQLite. A aplicação criará as tabelas `clients` e `transfers` automaticamente ao iniciar. O caminho padrão do banco de dados é `./bank.db`.

## Executando a Aplicação

Você pode rodar a aplicação localmente com Go ou usando Docker.

### Executando Localmente

1. Compile e execute a aplicação:

    ```bash
    go run src/main.go
    ```

2. Acesse a API em `http://localhost:8080`.

### Executando com Docker

1. Construa a imagem Docker:

    ```bash
    docker build -t banking-api .
    ```

2. Execute o contêiner:

    ```bash
    docker run -p 8080:8080 banking-api
    ```

## Endpoints da API

### Clientes

- **POST** `/v1/clients`: Cria um novo cliente.
- **GET** `/v1/clients`: Lista todos os clientes.
- **GET** `/v1/clients/{accountNum}`: Busca um cliente pelo número da conta.

### Transferências

- **POST** `/v1/transfer`: Realiza uma transferência entre duas contas.
- **GET** `/v1/transfers/{accountNum}`: Obtém o histórico de transferências associado a uma conta específica.

## Documentação Swagger

A documentação Swagger está disponível em `http://localhost:8080/swagger/index.html` após iniciar a aplicação.

### Configuração no Código

No arquivo `main.go`, a configuração Swagger está especificada para gerar a documentação baseada em anotações nos handlers.

Exemplo de anotação:

```go
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com as informações fornecidas
// @Tags clients
// @Accept json
// @Produce json
// @Param client body models.Client true "Cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Mensagem de erro"
// @Router /v1/clients [post]
```

# Comandos Úteis
Geração da documentação Swagger:

```bash
swag init -g src/main.go
```

# Remover a documentação:
```bash
rm -rf docs
```
# Estrutura do Projeto
## O projeto está organizado da seguinte forma:

banking/
├── go.mod
├── go.sum
└── src/
    ├── main.go                  # Ponto de entrada da aplicação
    ├── controllers/             # Controladores de rota
    │   ├── client_controller.go # Controller de Clientes
    │   ├── transfer_controller.go # Controller de Transferências
    │   └── routes.go            # Iniciação de Rotas
    ├── database/
    │   └── database.go          # Inicialização do banco de dados SQLite
    ├── models/                  # Definições de Modelos (Client, Transfer)
    ├── repositories/            # Repositórios de Acesso ao Banco
    └── services/                # Serviços de Lógica de Negócio

/

# Exemplo de Uso

## Criar um Cliente:

```bash
curl -X POST http://localhost:8080/v1/clients \
-H "Content-Type: application/json" \
-d '{
      "name": "John Doe",
      "account_num": "123456",
      "balance": 1000.0
    }'
```

## Listar Clientes:
```bash
curl -X GET http://localhost:8080/v1/clients
```

## Buscar um Cliente:
```bash
curl -X GET http://localhost:8080/v1/clients/123456
```

## Realizar uma Transferência:
```bash
curl -X POST http://localhost:8080/v1/transfer \
-H "Content-Type: application/json" \
-d '{
      "from_account": "123456",
      "to_account": "654321",
      "amount": 100.0
    }'
```

## Consultar Histórico de Transferências:
```bash
curl -X GET http://localhost:8080/v1/transfers/123456
```




