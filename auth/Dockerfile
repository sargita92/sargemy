FROM golang:1.25-alpine

WORKDIR /app

# Dependências básicas
RUN apk add --no-cache git

# Instala air
RUN go install github.com/air-verse/air@latest

# Copia apenas arquivos de dependência primeiro (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
