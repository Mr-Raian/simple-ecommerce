package data

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

//go:generate mockgen -typed -destination=../mocks/data.go -package=mocks . DataAccesor

type DataAccesor interface {
	GetConn() *sqlx.DB
	GetItemByID(ctx context.Context, id string, columns ...string) (Item, error)
	GetOrderByID(ctx context.Context, id string, columns ...string) (Order, error)
	UpdateOrder(ctx context.Context, order Order, where Where, columns ...string) error
	SaveContactForm(ctx context.Context, form *ContactForm) error
	GetConfigByKey(ctx context.Context, key string) (string, error)
	CreateOrder(ctx context.Context, itemID uuid.UUID, itemPrice uint, checkout Checkout, paymentMethod PaymentMethod, email string) (Order, error)
}

type Where struct {
	Column   string
	EqualsTo string
}
