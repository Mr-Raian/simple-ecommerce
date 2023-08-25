package db

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
)

func newTestingDataAccesor(t *testing.T) postgres {
	DB, err := ConnectToDB("postgresql://postgres:example@localhost/api")
	require.Nil(t, err)
	driver, err := pgmigrate.WithInstance(DB.DB, &pgmigrate.Config{})
	require.Nil(t, err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./../../migrations/",
		"postgres", driver)
	require.Nil(t, err)
	err = m.Down()
	if err != nil {
		if err.Error() != "no change" {
			t.Fatal(err)
		}
	}
	err = m.Up()
	if err != nil {
		t.Log(err)
	}
	require.Nil(t, err)
	return postgres{
		DB:         DB,
		Validaator: validator.New(),
	}
}
