# Usa uma imagem base com Go
FROM golang:1.22.4-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o go.mod e o go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências Go
RUN go mod download

# Copia o código-fonte para o diretório de trabalho
COPY . .

# Compila a aplicação Go
RUN go build -o main .

# Expõe a porta que sua aplicação usará
EXPOSE 80

# Comando para rodar a aplicação
CMD ["./main"]
