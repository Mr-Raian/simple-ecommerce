package handler

import (
	"api/internal/data"
	"context"
)

type Handler struct {
	DB                   data.DataAccesor
	Mailer               Mailer
	CardPaymentProccesor CardPaymentProccesor
}

//go:generate mockgen -typed -destination=../mocks/CardPaymentProccesor.go -package=mocks . CardPaymentProccesor
type CardPaymentProccesor interface {
	CreateCheckout(context.Context, uint) (data.Checkout, error)
	ParseWebhook(body []byte, header string) (status string, paymentID string, err error)
}

//go:generate mockgen -typed -destination=../mocks/Mailer.go -package=mocks . Mailer
type Mailer interface {
	Send(context.Context, *data.Email) error
}

type config interface {
	GetDataAccesor() data.DataAccesor
	GetMailer() Mailer
	GetCardPaymentProccesor() CardPaymentProccesor
}

func New(cfg config) *Handler {
	return &Handler{
		DB:     cfg.GetDataAccesor(),
		Mailer: cfg.GetMailer(),
	}
}
