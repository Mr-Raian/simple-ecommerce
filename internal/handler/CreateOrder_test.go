package handler_test

import (
	"api/internal/data"
	"api/internal/handler"
	"api/internal/mocks"
	"bytes"
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		Description      string
		Handlers         func(o *data.Order) handler.Handler
		Order            data.Order
		ExpectedHTTPCode int
		ExpectedError    error
	}{
		{
			Description: "200, No errors", Handlers: func(o *data.Order) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				o.ID = uuid.New()
				o.PaymentMethod = data.STRIPE
				o.OrderStatus = data.UNPAID
				item := data.Item{
					Price: uint(rand.Uint32()),
					ID:    uuid.New(),
				}
				o.ItemID = item.ID
				mockDB.EXPECT().GetItemByID(gomock.Any(), o.ItemID.String(), "price", "id").Return(item, nil).MaxTimes(1)
				mockCardP := mocks.NewMockCardPaymentProccesor(ctrl)
				mockCardP.EXPECT().CreateCheckout(gomock.Any(), item.Price)
				mockDB.EXPECT().CreateOrder(gomock.Any(), item.ID, item.Price, data.Checkout{}, o.PaymentMethod, o.Email).Return(*o, nil).MaxTimes(1)
				return handler.Handler{DB: mockDB, CardPaymentProccesor: mockCardP}
			},
			Order:            data.Order{Email: "m@m.c", PaymentMethod: "MONERO"},
			ExpectedHTTPCode: 200,
			ExpectedError:    nil,
		},
		{
			Description: "400, not found", Handlers: func(o *data.Order) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				o.ID = uuid.New()
				item := data.Item{
					Price: uint(rand.Uint32()),
					ID:    uuid.New(),
				}
				o.ItemID = item.ID
				mockDB.EXPECT().GetItemByID(gomock.Any(), o.ItemID.String(), "price", "id").Return(data.Item{}, sql.ErrNoRows).MaxTimes(1)
				return handler.Handler{DB: mockDB}
			},
			Order:            data.Order{Email: "m@m.c", PaymentMethod: "MONERO"},
			ExpectedHTTPCode: 400,
			ExpectedError:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			h := tt.Handlers(&tt.Order)
			form, err := json.Marshal(struct {
				ProductID     uuid.UUID `json:"item_id"`
				Email         string    `json:"email"`
				PaymentMethod string    `json:"payment_method"`
			}{
				ProductID:     tt.Order.ItemID,
				Email:         tt.Order.Email,
				PaymentMethod: string(tt.Order.PaymentMethod),
			})
			t.Log(string(form))
			require.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(form))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			err = h.CreateOrder(e.NewContext(req, rec))
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedHTTPCode, rec.Code)
		})
	}
}
