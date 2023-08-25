package db

import (
	"api/internal/data"

	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgres struct {
	DB         *sqlx.DB
	Validaator interface {
		Struct(interface{}) error
	}
}

func NewDataAccesor(DSN string) data.DataAccesor {
	DB, err := ConnectToDB(DSN)
	if err != nil {
		panic(err)
	}
	return postgres{
		DB:         DB,
		Validaator: validator.New(),
	}
}

func ConnectToDB(DSN string) (*sqlx.DB, error) {
	DB, err := sqlx.Open("pgx", DSN)
	if err != nil {
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		return nil, err
	}
	return DB, nil
}
