package l0

import "database/sql"

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}
