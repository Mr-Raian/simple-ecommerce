package payment

import (
	"api/internal/data"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/webhook"
)

type Stripe struct {
	EndpointSecret string
}

func New(key string) Stripe {
	stripe.Key = key
	return Stripe{}
}

func (rcv Stripe) CreateCheckout(ctx context.Context, price uint) (data.Checkout, error) {
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("https://freedomof.tech/"),
					},
					UnitAmount: stripe.Int64(int64(price)),
				},
				Quantity: stripe.Int64(int64(1)),
			},
		},
	}
	s, err := session.New(params)
	return data.Checkout{PaymentURL: s.URL, PaymentID: s.ID}, err
}

var ErrUnimplementedStripeEvent = errors.New("ErrUnimplementedStripeEvent")

func (rcv Stripe) ParseWebhook(body []byte, header string) (status string, paymentID string, err error) {
	event, err := webhook.ConstructEvent(body, header, rcv.EndpointSecret)
	if err != nil {
		return "", "", err
	}

	switch event.Type {
	case "checkout.session.completed":
		var s stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &s); err != nil {
			return "", "", err
		}
		return "PAID", s.ID, nil
	default:
		return "", "", ErrUnimplementedStripeEvent
	}
}
