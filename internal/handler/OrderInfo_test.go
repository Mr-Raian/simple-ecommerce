package handler_test

import (
	"api/internal/data"
	"api/internal/handler"
	"api/internal/mocks"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestOrderInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		Description      string
		Handlers         func(string) handler.Handler
		ID               string
		ExpectedHTTPCode int
		ExpectedError    error
	}{
		{
			Description: "200, No errors", Handlers: func(id string) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				mockDB.EXPECT().GetOrderByID(gomock.Any(), id).Return(data.Order{}, nil).MaxTimes(1)
				return handler.Handler{DB: mockDB}
			},
			ID:               uuid.New().String(),
			ExpectedHTTPCode: 200,
			ExpectedError:    nil,
		},
		{
			Description: "404", Handlers: func(id string) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				mockDB.EXPECT().GetOrderByID(gomock.Any(), id).Return(data.Order{}, sql.ErrNoRows).MaxTimes(1)
				return handler.Handler{DB: mockDB}
			},
			ID:               uuid.New().String(),
			ExpectedHTTPCode: 404,
			ExpectedError:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			h := tt.Handlers(tt.ID)
			form, err := json.Marshal(struct {
				ID string `json:"id"`
			}{tt.ID})
			require.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(form))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			err = h.OrderInfo(e.NewContext(req, rec))
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedHTTPCode, rec.Code)
		})
	}
}
