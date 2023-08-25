package handler_test

import (
	"api/internal/data"
	"api/internal/handler"
	"api/internal/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestContactForm(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		Description      string
		Handlers         func(f *data.ContactForm) handler.Handler
		form             data.ContactForm
		ExpectedHTTPCode int
		ExpectedError    error
	}{
		{
			Description: "200, No errors", Handlers: func(f *data.ContactForm) handler.Handler {
				mockDB := mocks.NewMockDataAccesor(ctrl)
				mockDB.EXPECT().SaveContactForm(gomock.Any(), f).Return(nil).MaxTimes(1)
				mockDB.EXPECT().GetConfigByKey(gomock.Any(), "admin_email").Return("admin@admin.com", nil).MaxTimes(1)
				mockMailer := mocks.NewMockMailer(ctrl)
				mockMailer.EXPECT().Send(gomock.Any(), gomock.Any()).MaxTimes(2)
				return handler.Handler{DB: mockDB, Mailer: mockMailer}
			},
			form:             data.ContactForm{FirstName: "Joe", Email: "m@m.com", Message: "hey"},
			ExpectedHTTPCode: 200,
			ExpectedError:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			h := tt.Handlers(&tt.form)
			form, err := json.Marshal(tt.form)
			require.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(form))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			err = h.ContactForm(e.NewContext(req, rec))
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedHTTPCode, rec.Code)
		})
	}
}
