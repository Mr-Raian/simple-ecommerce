package db

import (
	"api/internal/data"
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetItemByID(t *testing.T) {
	tests := []struct {
		Description             string
		PrepareDB               func(*testing.T) (postgres, data.Item)
		ColumnsToSelect         []string
		ExpectedErr             error
		CheckOutputExpectations func(*testing.T, data.Item)
	}{
		{
			Description: "All ok",
			PrepareDB: func(t *testing.T) (postgres, data.Item) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				_, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata)", item)
				require.Nil(t, err)
				return DB, item
			},
			ExpectedErr: nil,
			CheckOutputExpectations: func(t *testing.T, i data.Item) {
				assert.NotZero(t, i.ID)
				assert.NotZero(t, i.Price)
				assert.NotZero(t, i.Metadata)
			},
		},
		{
			Description: "Get price only",
			PrepareDB: func(t *testing.T) (postgres, data.Item) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				_, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata)", item)
				require.Nil(t, err)
				return DB, item
			},
			ColumnsToSelect: []string{"price"},
			ExpectedErr:     nil,
			CheckOutputExpectations: func(t *testing.T, i data.Item) {
				assert.Zero(t, i.ID)
				assert.NotZero(t, i.Price)
				assert.Zero(t, i.Metadata)
			},
		},
		{
			Description: "All ok",
			PrepareDB: func(t *testing.T) (postgres, data.Item) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				// _, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata)", item)
				// require.Nil(t, err)
				return DB, item
			},
			ExpectedErr: sql.ErrNoRows,
			CheckOutputExpectations: func(t *testing.T, i data.Item) {
				assert.Zero(t, i.ID)
				assert.Zero(t, i.Price)
				assert.Zero(t, i.Metadata)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres, item := tt.PrepareDB(t)
			ctx := context.TODO()
			var err error
			dbItem := data.Item{}
			if len(tt.ColumnsToSelect) == 0 {
				dbItem, err = postgres.GetItemByID(ctx, item.ID.String())
			} else {
				dbItem, err = postgres.GetItemByID(ctx, item.ID.String(), tt.ColumnsToSelect...)
			}
			t.Log(dbItem)
			assert.Equal(t, tt.ExpectedErr, err)
			tt.CheckOutputExpectations(t, dbItem)
		})
	}
}
