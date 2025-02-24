# Usa uma imagem oficial do Go para build
FROM golang:1.22.2 AS builder

ARG COMMIT_HASH
LABEL commit-hash=$COMMIT_HASH
# Define o diretório de trabalho no contêiner
WORKDIR /app

# Copia os arquivos do projeto para dentro do contêiner
COPY . .

# Compila o binário para Linux sem dependências dinâmicas
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Usa uma imagem mínima para rodar a aplicação
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /root/

# Copia o binário do estágio de build para a imagem final
COPY --from=builder /app/main .

# Dá permissão de execução para o binário
RUN chmod +x ./main

# Expõe a porta da aplicação
EXPOSE 8000

# Comando padrão para rodar o app
CMD ["./main"]
