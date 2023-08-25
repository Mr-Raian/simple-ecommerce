package data

import "github.com/google/uuid"

// TODO: Add validation tags
type Order struct {
	ID            uuid.UUID     `json:"id" db:"id"`
	ItemID        uuid.UUID     `json:"item_id" db:"item_id"`
	Price         uint          `json:"price" db:"price"`
	Email         string        `json:"email" db:"email"`
	PaymentMethod PaymentMethod `json:"payment_methos" db:"payment_method"`
	PaymentURL    string        `json:"paymnet_url" db:"payment_url"`
	PaymentID     string        `json:"payment_id" db:"payment_id"`
	OrderStatus   OrderStatus   `json:"order_status" db:"order_status"`
}

type PaymentMethod string

const (
	MONERO PaymentMethod = "MONERO"
	STRIPE PaymentMethod = "STRIPE"
)

type OrderStatus string

const (
	PAID    OrderStatus = "PAID"
	UNPAID  OrderStatus = "UNPAID"
	EXPIRED OrderStatus = "EXPIRED"
)
