package db

import (
	"api/internal/data"
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrCantInsert = errors.New("can't insert")

func (rcv postgres) CreateOrder(ctx context.Context, itemID uuid.UUID, itemPrice uint, checkout data.Checkout, paymentMethod data.PaymentMethod, email string) (data.Order, error) {
	order := data.Order{
		ID:            uuid.New(),
		ItemID:        itemID,
		Price:         itemPrice,
		Email:         email,
		PaymentMethod: paymentMethod,
		OrderStatus:   data.UNPAID,
		PaymentID:     checkout.PaymentID,
		PaymentURL:    checkout.PaymentURL,
	}
	if err := rcv.Validaator.Struct(&order); err != nil {
		return order, err
	}
	if _, err := rcv.DB.NamedExecContext(ctx, "INSERT INTO orders(id, item_id, email, payment_method, payment_url, order_status, price, payment_id) VALUES (:id, :item_id, :email, :payment_method, :payment_url, :order_status, :price, :payment_id)", &order); err != nil {
		return order, errors.Join(ErrCantInsert, err)
	}
	return order, nil
}
