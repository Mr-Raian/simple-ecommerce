package data

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ContactForm struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
	FirstName string    `db:"first_name" json:"first_name" validate:"max=64"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
	Message   string    `db:"message" json:"message" validate:"required"`
}

var ErrContactFormValidationFailed = errors.New("ContactFormValidationFailed")
