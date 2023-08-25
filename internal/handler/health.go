package handler

import (
	"github.com/labstack/echo/v4"
)

func (rcv *Handler) HealthCheck(c echo.Context) error {
	if err := rcv.DB.GetConn().Ping(); err != nil {
		return &echo.HTTPError{Code: 500, Message: "Internal error", Internal: err}
	}
	return c.String(200, "OK")
}
