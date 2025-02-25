package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		panic("X ERRO: A variável de ambiente DATABASE_URL não está definida!")
	}

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