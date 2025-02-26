package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Iniciando conexão...")
	dbURL := "postgres://postgres:12345@db:5432/postgres?sslmode=disable"
	fmt.Println("Etapa 1")

	if dbURL == "" {
		panic("X ERRO: A variável de ambiente DATABASE_URL não está definida!")
	}
	fmt.Println("Etapa 2")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connectado com sucesso")

	return db, nil
}