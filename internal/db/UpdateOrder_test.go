package db

import (
	"api/internal/data"
	"context"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateOrder(t *testing.T) {
	tests := []struct {
		Description             string
		PrepareDB               func(*testing.T) (postgres, data.Order, data.Where)
		ColumnsToUpdate         []string
		ExpectedErr             error
		CheckOutputExpectations func(*testing.T, data.Order)
	}{
		{
			Description:     "All ok, update from unpaid to paid",
			ColumnsToUpdate: []string{"order_status"},
			PrepareDB: func(t *testing.T) (postgres, data.Order, data.Where) {
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
					PaymentID:     uuid.New().String(),
					OrderStatus:   "UNPAID",
					ItemID:        item.ID,
				}
				_, err = DB.DB.NamedExec("INSERT INTO orders(id, item_id, email, payment_method, payment_url, order_status, price, payment_id) VALUES (:id, :item_id, :email, :payment_method, :payment_url, :order_status, :price, :payment_id)", &o)
				require.Nil(t, err)
				o.OrderStatus = "PAID"
				return DB, o, data.Where{Column: "payment_id", EqualsTo: o.PaymentID}
			},
			ExpectedErr: nil,
			CheckOutputExpectations: func(t *testing.T, o data.Order) {
				assert.Equal(t, data.PAID, o.OrderStatus)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres, order, where := tt.PrepareDB(t)
			ctx := context.TODO()
			var err error
			err = postgres.UpdateOrder(ctx, order, where, tt.ColumnsToUpdate...)
			assert.Equal(t, tt.ExpectedErr, err)
			dbOrder, err := postgres.GetOrderByID(ctx, order.ID.String())
			require.Nil(t, err)
			tt.CheckOutputExpectations(t, dbOrder)
		})
	}
}
