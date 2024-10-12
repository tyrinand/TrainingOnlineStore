package repository

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}

func New(dataBase *sqlx.DB) *repository {
	return &repository{db: dataBase}
}
