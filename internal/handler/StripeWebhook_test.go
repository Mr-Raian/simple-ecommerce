package handler

import (
	"api/internal/data"
	"api/internal/mocks"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStripeWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		Description      string
		Prepare          func(t *testing.T) (Handler, []byte, string)
		ExpectedHTTPCode int
		ExpectedError    error
	}{
		{
			Description: "200, no errors",
			Prepare: func(t *testing.T) (Handler, []byte, string) {
				mockCardP := mocks.NewMockCardPaymentProccesor(ctrl)
				header := "123123"
				body := []byte{0, 1, 2, 3}
				payment_id := "12323123123"
				mockCardP.EXPECT().ParseWebhook(body, header).Return(string(data.PAID), payment_id, nil).MaxTimes(1)
				mockDB := mocks.NewMockDataAccesor(ctrl)
				mockDB.EXPECT().UpdateOrder(gomock.Any(), gomock.Any(), data.Where{Column: "payment_id", EqualsTo: payment_id}, "order_status").Return(nil).MaxTimes(1)
				return Handler{DB: mockDB, CardPaymentProccesor: mockCardP}, body, header
			},
			ExpectedHTTPCode: 200,
			ExpectedError:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			h, body, header := tt.Prepare(t)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
			req.Header.Set("Stripe-Signature", header)
			rec := httptest.NewRecorder()
			err := h.StripeWebhook(e.NewContext(req, rec))
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedHTTPCode, rec.Code)
		})
	}
}
