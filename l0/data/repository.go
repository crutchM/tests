package data

import (
	"database/sql"
	"tests/l0/data/db"
	domain "tests/l0/domain"
)

type Repository struct {
	domain.PgRepo
}

func NewRepository(bd *sql.DB) *Repository {
	return &Repository{
		PgRepo: db.NewPgRepository(bd),
	}
}
