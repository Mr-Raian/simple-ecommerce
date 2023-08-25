package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigByKey(t *testing.T) {
	tests := []struct {
		Description string
		PrepareDB   func(*testing.T) (postgres, map[string]string)
		ExpectedErr error
	}{
		{
			Description: "All ok",
			PrepareDB: func(t *testing.T) (postgres, map[string]string) {
				DB := newTestingDataAccesor(t)
				config := map[string]string{
					"admin_email": "admin@admin.com",
				}
				for k, v := range config {
					_, err := DB.DB.Exec("INSERT INTO config(key, value) VALUES ($1, $2)", k, v)
					require.Nil(t, err)
				}
				return DB, config
			},
			ExpectedErr: nil,
		},
		{
			Description: "Not found",
			PrepareDB: func(t *testing.T) (postgres, map[string]string) {
				DB := newTestingDataAccesor(t)
				config := map[string]string{
					"admin_email": "admin@admin.com",
				}
				return DB, config
			},
			ExpectedErr: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres, config := tt.PrepareDB(t)
			ctx := context.TODO()
			for k, v := range config {
				value, err := postgres.GetConfigByKey(ctx, k)
				if tt.ExpectedErr == nil {
					assert.Equal(t, v, value)
				}
				require.Equal(t, tt.ExpectedErr, err)
			}
		})
	}
}
