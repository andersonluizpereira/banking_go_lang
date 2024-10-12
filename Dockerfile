# Usa uma imagem base do Go
FROM golang:1.18

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código para o diretório de trabalho
COPY . .

# Compila a aplicação
RUN go build -o main .

# Define o comando para executar a aplicação
CMD ["./main"]

# Expõe a porta que o servidor usa
EXPOSE 8080
