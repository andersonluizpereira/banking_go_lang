# Usa uma imagem base do Go
FROM golang:1.18 as builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o go.mod e o go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte do projeto para o diretório de trabalho
COPY . .

# Compila o binário da aplicação
RUN go build -o bankingapp main.go

# Segunda etapa para criar uma imagem menor
FROM golang:1.18

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o binário da etapa de compilação para a nova imagem
COPY --from=builder /app/bankingapp .

# Copia arquivos necessários, como configurações e dependências
COPY ./src /app/src

# Expõe a porta 8080 para o servidor
EXPOSE 8080

# Define o comando padrão para executar o contêiner com o CLI
CMD ["./bankingapp", "run"]
