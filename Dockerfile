FROM golang:1.22.2 AS builder

ARG COMMIT_HASH
LABEL commit-hash=$COMMIT_HASH

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env.prod ./.env

RUN chmod +x ./main

EXPOSE 8000

CMD export $(grep -v '^#' .env | xargs) && ./main
