package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func (rcv *Handler) OrderInfo(c echo.Context) error {
	var input struct {
		ID string `json:"id"`
	}
	if err := c.Bind(&input); err != nil {
		return err
	}
	order, err := rcv.DB.GetOrderByID(c.Request().Context(), input.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, "not found")
		}
		return err
	}
	return c.JSON(200, order)
}
