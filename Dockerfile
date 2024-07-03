# Usa uma imagem base com Go na versão do projeto
FROM golang:1.22.4-alpine

WORKDIR /app
COPY go.mod go.sum ./

# Baixa as dependências Go
RUN go mod download
COPY . .

# Compila a aplicação Go
RUN go build -o main .
EXPOSE 80

# Comando para rodar a aplicação
CMD ["./main"]
