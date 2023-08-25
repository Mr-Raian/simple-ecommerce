package handler_test

import (
	"api/internal/data"
	"api/internal/handler"
	"api/internal/mocks"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetItemData(t *testing.T) {
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
				mockDB.EXPECT().GetItemByID(gomock.Any(), id).Return(data.Item{}, nil).MaxTimes(1)
				return handler.Handler{DB: mockDB}
			},
			ID:               uuid.New().String(),
			ExpectedHTTPCode: 200,
			ExpectedError:    nil,
		},
		{
			Description: "404, Not found", Handlers: func(id string) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				mockDB.EXPECT().GetItemByID(gomock.Any(), id).Return(data.Item{}, sql.ErrNoRows).MaxTimes(1)
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
			req := httptest.NewRequest(http.MethodGet, "/?product_id="+tt.ID, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			err := h.GetItemData(e.NewContext(req, rec))
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedHTTPCode, rec.Code)
		})
	}
}
