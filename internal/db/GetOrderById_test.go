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

func TestGetOrderById(t *testing.T) {
	tests := []struct {
		Description             string
		PrepareDB               func(*testing.T) (postgres, data.Order)
		ColumnsToSelect         []string
		ExpectedErr             error
		CheckOutputExpectations func(*testing.T, data.Order)
	}{
		{
			Description: "All ok",
			PrepareDB: func(t *testing.T) (postgres, data.Order) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				_, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata)", item)
				require.Nil(t, err)
				o := data.Order{
					ID:            uuid.New(),
					Price:         uint(rand.Uint32()),
					Email:         "mail@mail.com",
					PaymentURL:    "http://order.com",
					PaymentMethod: "STRIPE",
					OrderStatus:   "PAID",
					ItemID:        item.ID,
				}
				_, err = DB.DB.NamedExec("INSERT INTO orders(id, item_id, email, payment_method, payment_url, order_status, price, payment_id) VALUES (:id, :item_id, :email, :payment_method, :payment_url, :order_status, :price, :payment_id)", &o)
				require.Nil(t, err)
				return DB, o
			},
			ExpectedErr: nil,
			CheckOutputExpectations: func(t *testing.T, o data.Order) {
				assert.NotZero(t, o)
				assert.NotZero(t, o.Price)
				assert.NotZero(t, o.Email)
				assert.NotZero(t, o.ID)
				assert.NotZero(t, o.OrderStatus)
				assert.NotZero(t, o.PaymentURL)
			},
		},
		{
			Description:     "Get only one column",
			ColumnsToSelect: []string{"price"},
			PrepareDB: func(t *testing.T) (postgres, data.Order) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				_, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata)", item)
				require.Nil(t, err)
				o := data.Order{
					ID:            uuid.New(),
					Price:         uint(rand.Uint32()),
					Email:         "mail@mail.com",
					PaymentURL:    "http://order.com",
					PaymentMethod: "STRIPE",
					OrderStatus:   "PAID",
					ItemID:        item.ID,
				}
				_, err = DB.DB.NamedExec("INSERT INTO orders(id, item_id, email, payment_method, payment_url, order_status, price, payment_id) VALUES (:id, :item_id, :email, :payment_method, :payment_url, :order_status, :price, :payment_id)", &o)
				require.Nil(t, err)
				return DB, o
			},
			ExpectedErr: nil,
			CheckOutputExpectations: func(t *testing.T, o data.Order) {
				assert.NotZero(t, o)
				assert.NotZero(t, o.Price)
				assert.Zero(t, o.Email)
				assert.Zero(t, o.ID)
				assert.Zero(t, o.OrderStatus)
				assert.Zero(t, o.PaymentURL)
			},
		},
		{
			Description: "Not found",
			PrepareDB: func(t *testing.T) (postgres, data.Order) {
				DB := newTestingDataAccesor(t)
				return DB, data.Order{ID: uuid.New()}
			},
			ExpectedErr: sql.ErrNoRows,
			CheckOutputExpectations: func(t *testing.T, o data.Order) {
				assert.Zero(t, o)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres, item := tt.PrepareDB(t)
			ctx := context.TODO()
			var err error
			dbItem := data.Order{}
			if len(tt.ColumnsToSelect) == 0 {
				dbItem, err = postgres.GetOrderByID(ctx, item.ID.String())
			} else {
				dbItem, err = postgres.GetOrderByID(ctx, item.ID.String(), tt.ColumnsToSelect...)
			}
			assert.Equal(t, tt.ExpectedErr, err)
			tt.CheckOutputExpectations(t, dbItem)
		})
	}
}
