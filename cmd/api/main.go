package main

import (
	"api/internal/amz"
	"api/internal/data"
	"api/internal/db"
	"api/internal/handler"
	"api/internal/payment"
	"api/internal/router"
	"context"
)

type config struct {
	ENV       string
	ListenOn  string
	DSN       string
	StripeKey string
}

func (rcv config) GetDataAccesor() data.DataAccesor {
	return db.NewDataAccesor(rcv.DSN)
}

func (rcv config) GetMailer() handler.Mailer {
	return amz.NewSesClient(context.TODO())
}

func (rcv config) GetCardPaymentProccesor() handler.CardPaymentProccesor {
	return payment.New(rcv.StripeKey)
}
func main() {
	cfg := config{
		ListenOn: "127.0.0.1:8080",
		ENV:      "dev",
		DSN:      "postgresql://postgres:example@localhost",
	}
	r := router.New(*handler.New(cfg))
	r.Debug = true
	r.Logger.Fatal(r.Start(cfg.ListenOn))
}
