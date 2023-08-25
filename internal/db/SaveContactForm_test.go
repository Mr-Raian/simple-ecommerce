package db

import (
	"api/internal/data"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestSaveContactForm(t *testing.T) {
	tests := []struct {
		Description         string
		Form                data.ContactForm
		PrepareDB           func(*testing.T) postgres
		CheckDBExpectations func(*testing.T, postgres, data.ContactForm)
		ExpectedErr         error
	}{
		{
			Description: "All ok",
			Form:        data.ContactForm{FirstName: "Joe", Email: "joe@joe.com", Message: "hey joe"},
			PrepareDB: func(t *testing.T) postgres {
				return newTestingDataAccesor(t)
			},
			ExpectedErr: nil,
			CheckDBExpectations: func(t *testing.T, p postgres, cf data.ContactForm) {
				var found int
				err := p.DB.Get(&found, "SELECT count(id) FROM contact_forms WHERE first_name = $1 AND email = $2 AND message = $3;", cf.FirstName, cf.Email, cf.Message)
				assert.Nil(t, err)
				assert.Equal(t, 1, found)
			},
		},
		{
			Description: "Invalid email",
			Form:        data.ContactForm{FirstName: "Joe", Email: "joe#joe.com", Message: "hey joe"},
			PrepareDB: func(t *testing.T) postgres {
				return newTestingDataAccesor(t)
			},
			ExpectedErr: data.ErrContactFormValidationFailed,
			CheckDBExpectations: func(t *testing.T, p postgres, cf data.ContactForm) {
				var found int
				err := p.DB.Get(&found, "SELECT count(id) FROM contact_forms WHERE first_name = $1 AND email = $2 AND message = $3;", cf.FirstName, cf.Email, cf.Message)
				assert.Nil(t, err)
				assert.Equal(t, 0, found)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			postgres := tt.PrepareDB(t)
			ctx := context.TODO()
			err := postgres.SaveContactForm(ctx, &tt.Form)
			assert.Equal(t, tt.ExpectedErr, err)
			tt.CheckDBExpectations(t, postgres, tt.Form)
		})
	}
}
