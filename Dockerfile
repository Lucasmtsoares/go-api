# Usa uma imagem oficial do Go para build
FROM golang:1.22.2-alpine AS builder

ARG COMMIT_HASH
LABEL commit-hash=$COMMIT_HASH

WORKDIR /app  

COPY go.mod go.sum .env.prod ./  
COPY cmd ./cmd

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Imagem final
FROM alpine:latest

WORKDIR /app 

RUN apk add --no-cache bash

# Correções importantes aqui:
COPY --from=builder /app/main .  
COPY --from=builder /app/.env.prod ./.env  

RUN chmod +x ./main

EXPOSE 8000

# Forma recomendada de passar variáveis de ambiente:
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB

ENV POSTGRES_USER=$POSTGRES_USER
ENV POSTGRES_PASSWORD=$POSTGRES_PASSWORD
ENV POSTGRES_DB=$POSTGRES_DB

ENV DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@go_db:5432/${POSTGRES_DB}?sslmode=disable

CMD ["./main"]  # Executa o programa diretamente