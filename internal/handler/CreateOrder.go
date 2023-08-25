package handler

import (
	"api/internal/data"
	"database/sql"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (rcv *Handler) CreateOrder(c echo.Context) error {
	var input struct {
		ItemID        uuid.UUID          `json:"item_id"`
		Email         string             `json:"email"`
		PaymentMethod data.PaymentMethod `json:"payment_method"`
	}
	if err := c.Bind(&input); err != nil {
		return err
	}
	ctx := c.Request().Context()
	item, err := rcv.DB.GetItemByID(ctx, input.ItemID.String(), "price", "id")
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, echo.NewHTTPError(400, "item not found"))
		}
		return err
	}
	checkout, err := rcv.CardPaymentProccesor.CreateCheckout(ctx, item.Price)
	if err != nil {
		return err
	}
	order, err := rcv.DB.CreateOrder(ctx, input.ItemID, item.Price, checkout, input.PaymentMethod, input.Email)
	if err != nil {
		return err
	}
	return c.JSON(200, order)
}
