package l0

import (
	"github.com/jmoiron/sqlx"
)

type DataBase struct {
	db *sqlx.DB
}

func NewDataBase(db *sqlx.DB) *DataBase {
	return &DataBase{db: db}
}
