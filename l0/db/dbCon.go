package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
}

func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("postgres://habrpguser:pgpwd4habr@localhost:5432/postgres?sslmode=disable"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
