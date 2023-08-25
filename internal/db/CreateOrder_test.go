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

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		Description         string
		PrepareDB           func(*testing.T) (postgres, data.Order)
		CheckDBExpectations func(*testing.T, postgres, data.Order)
		ExpectedErr         error
	}{
		{
			Description: "All ok",
			ExpectedErr: nil,
			PrepareDB: func(t *testing.T) (postgres, data.Order) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID:       uuid.New(),
					Price:    uint(rand.Uint32()),
					Metadata: `{"meta":"data"}`,
				}
				_, err := DB.DB.NamedExec("INSERT INTO items(id, price, metadata) VALUES (:id, :price, :metadata);", item)
				require.Nil(t, err)
				return DB, data.Order{ItemID: item.ID, Email: "random@email.com", PaymentMethod: "STRIPE", Price: uint(rand.Uint32()), PaymentID: "123", PaymentURL: "http"}
			},
			CheckDBExpectations: func(t *testing.T, p postgres, o data.Order) {
				var found int
				rows, err := p.DB.NamedQuery("SELECT count(id) FROM orders WHERE id = :id AND email = :email AND item_id = :item_id AND price = :price AND payment_method = :payment_method AND payment_url = :payment_url AND payment_id = :payment_id AND order_status = :order_status;", o)
				require.Nil(t, err)
				defer rows.Close()
				t.Log(rows)
				require.True(t, rows.Next())
				rows.Scan(&found)
				assert.Nil(t, err)
				assert.Equal(t, 1, found)
			},
		},
		{
			Description: "item not found",
			ExpectedErr: ErrCantInsert,
			PrepareDB: func(t *testing.T) (postgres, data.Order) {
				DB := newTestingDataAccesor(t)
				item := data.Item{
					ID: uuid.New(),
				}
				return DB, data.Order{ItemID: item.ID, Email: "random@email.com", PaymentMethod: "STRIPE"}
			},
			CheckDBExpectations: func(t *testing.T, p postgres, o data.Order) {
				var found int
				rows, err := p.DB.NamedQuery("SELECT count(id) FROM orders WHERE id = :id AND email = :email AND item_id = :item_id AND price = :price AND payment_method = :payment_method AND payment_url = :payment_url AND payment_id = :payment_id AND order_status = :order_status;", o)
				require.Nil(t, err)
				defer rows.Close()
				t.Log(rows)
				require.True(t, rows.Next())
				rows.Scan(&found)
				assert.Nil(t, err)
				assert.Equal(t, 0, found)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres, o := tt.PrepareDB(t)
			ctx := context.TODO()
			order, err := postgres.CreateOrder(ctx, o.ItemID, o.Price, data.Checkout{PaymentID: o.PaymentID, PaymentURL: o.PaymentURL}, o.PaymentMethod, o.Email)
			assert.ErrorIs(t, err, tt.ExpectedErr)
			tt.CheckDBExpectations(t, postgres, order)
		})
	}
}
