package l0

import "database/sql"

type Repository struct {
	cacheRepo
	pgRepo
}

func NewRepository(bd *sql.DB) *Repository {
	return &Repository{cacheRepo: NewCacheRepository(), pgRepo: NewPgRepository(bd)}
}
