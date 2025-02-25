# Usa uma imagem oficial do Go para build
FROM golang:1.22.2 AS builder

ARG COMMIT_HASH
LABEL commit-hash=$COMMIT_HASH

WORKDIR /app

COPY go.mod go.sum .env.prod ./
COPY cmd ./cmd

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /app/main .
COPY --from=builder /app/.env.prod ./.env

RUN chmod +x ./main

EXPOSE 8000

CMD sh -c "export $(grep -v '^#' .env | xargs) && ./main"
