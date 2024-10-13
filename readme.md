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
## Geração da documentação Swagger:

```bash
    swag init -g src/main.go
```

# Remover a documentação:
```bash
rm -rf docs
```

# Estrutura do Projeto
## O projeto está organizado da seguinte forma:
```
banking/
├── Dockerfile
├── LICENSE
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── readme.md
├── src
│   ├── bank.db
│   ├── controllers
│   │   ├── client_controller.go
│   │   └── transfer_controller.go
│   ├── database
│   │   └── database.go
│   ├── main.go
│   ├── models
│   │   ├── client.go
│   │   └── transfer.go
│   ├── repositories
│   │   ├── client_repository.go
│   │   └── transfer_repository.go
│   └── services
│       ├── client_service.go
│       └── transfer_service.go
└── tests
    ├── controllers
    │   ├── client_controller_integration_test.go
    │   └── transfer_controller_integration_test.go
    ├── database
    │   └── database_integration_test.go
    ├── models
    │   ├── client_test.go
    │   └── transfer_test.go
    ├── repositories
    │   ├── client_repository_integration_test.go
    │   ├── test_helpers.go
    │   └── transfer_repository_integration_test.go
    └── services
        ├── client_service_test.go
        ├── mock_repositories.go
        └── transfer_service_test.go
               # Serviços de Lógica de Negócio
```
```
Explicação da Estrutura de Diretórios
Arquivos e Diretórios na Raiz
Dockerfile: Arquivo para construir a imagem Docker da aplicação. Contém as instruções para configurar o ambiente necessário para rodar a API no contêiner.

LICENSE: Arquivo de licença, onde são descritos os direitos de uso do projeto.

bank.db: Arquivo de banco de dados SQLite. Este é o arquivo onde as tabelas e os dados da aplicação são armazenados. Pode ser recriado automaticamente ao iniciar o projeto.

docker-compose.yml: Arquivo de configuração do Docker Compose, que permite orquestrar o contêiner da aplicação e quaisquer serviços adicionais (como banco de dados).

docs: Diretório que contém os arquivos de documentação do Swagger.

docs.go: Gerado automaticamente pelo swag e serve como entrada para a documentação Swagger.
swagger.json e swagger.yaml: Arquivos de especificação OpenAPI para a API, em formatos JSON e YAML.
go.mod e go.sum: Arquivos de configuração do módulo Go.

go.mod: Especifica o módulo e suas dependências.
go.sum: Contém hashes de verificação de integridade para as dependências do projeto.
readme.md: Documentação principal do projeto. Geralmente contém instruções de instalação, execução e detalhes sobre a API.

Diretório src
Contém o código-fonte principal do projeto.

src/bank.db: Backup ou versão inicial do banco de dados SQLite para configuração rápida.

controllers: Diretório onde estão os controladores que definem os endpoints da API.

client_controller.go: Controlador para endpoints relacionados a clientes.
transfer_controller.go: Controlador para endpoints de transferência entre contas.
database: Diretório para configuração e inicialização do banco de dados.

database.go: Arquivo responsável por conectar-se ao SQLite, executar migrações e garantir que as tabelas estejam prontas para uso.
main.go: Ponto de entrada da aplicação. Este arquivo inicializa o banco de dados, configura as rotas e inicia o servidor.

models: Diretório com a definição dos modelos de dados.

client.go: Define o modelo Client, representando a tabela de clientes.
transfer.go: Define o modelo Transfer, representando a tabela de transferências.
repositories: Diretório que contém os repositórios, que lidam com o acesso ao banco de dados.

client_repository.go: Funções de acesso ao banco para operações relacionadas a clientes.
transfer_repository.go: Funções de acesso ao banco para operações relacionadas a transferências.
services: Diretório que contém a lógica de negócios da aplicação.

client_service.go: Contém a lógica de negócios para operações relacionadas a clientes.
transfer_service.go: Contém a lógica de negócios para operações de transferência.
Diretório tests
Contém testes unitários e de integração para todas as camadas da aplicação. Os testes estão organizados por funcionalidade, refletindo a estrutura do código principal.

controllers: Contém testes de integração para os controladores da API.

client_controller_integration_test.go: Testa endpoints relacionados a clientes.
transfer_controller_integration_test.go: Testa endpoints de transferência.
database: Contém testes de integração para o banco de dados.

database_integration_test.go: Testes de integração para validar a inicialização e migração do banco de dados.
models: Testes para os modelos de dados.

client_test.go: Testes unitários para o modelo Client.
transfer_test.go: Testes unitários para o modelo Transfer.
repositories: Contém testes para as funções de acesso ao banco de dados.

client_repository_integration_test.go: Testa as operações de banco de dados relacionadas a clientes.
transfer_repository_integration_test.go: Testa as operações de banco de dados relacionadas a transferências.
test_helpers.go: Funções auxiliares usadas nos testes dos repositórios, como criação de dados de teste.
services: Contém os testes de unidade para os serviços da aplicação.

client_service_test.go: Testa as regras de negócios para operações de clientes.
transfer_service_test.go: Testa as regras de negócios para operações de transferência.
mock_repositories.go: Mock dos repositórios, usado nos testes para simular comportamentos e evitar o acesso ao banco de dados real.
```

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

