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
