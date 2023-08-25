package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func (rcv *Handler) GetItemData(c echo.Context) error {
	item, err := rcv.DB.GetItemByID(c.Request().Context(), c.QueryParam("product_id"))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, echo.NewHTTPError(404))
		}
		return err
	}

	return c.JSON(200, item)
}
