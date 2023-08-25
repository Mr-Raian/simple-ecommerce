package handler

import (
	"api/internal/data"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func (rcv Handler) StripeWebhook(c echo.Context) error {
	body := c.Request().Body
	header := c.Request().Header.Get("Stripe-Signature")
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	status, paymentID, err := rcv.CardPaymentProccesor.ParseWebhook(bodyBytes, header)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if status == "PAID" {
		if err := rcv.DB.UpdateOrder(ctx, data.Order{OrderStatus: "PAID"}, data.Where{Column: "payment_id", EqualsTo: paymentID}, "order_status"); err != nil {
			return err
		}
	}

	return c.JSON(200, echo.NewHTTPError(200))
}
