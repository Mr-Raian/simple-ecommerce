package db

import (
	"api/internal/data"
	"context"
	"time"

	"github.com/google/uuid"
)

func (rcv postgres) SaveContactForm(ctx context.Context, form *data.ContactForm) error {
	if err := rcv.Validaator.Struct(form); err != nil {
		return data.ErrContactFormValidationFailed
	}
	form.Timestamp = time.Now()
	form.ID = uuid.New()
	if _, err := rcv.DB.NamedExecContext(ctx,
		"INSERT INTO contact_forms(id, first_name, email, message, timestamp) VALUES (:id, :first_name, :email, :message, :timestamp);", form); err != nil {
		return err
	}
	return nil
}
