package handler

import (
	"api/internal/data"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (rcv *Handler) ContactForm(c echo.Context) error {
	form := data.ContactForm{}
	if err := c.Bind(&form); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := rcv.DB.SaveContactForm(ctx, &form); err != nil {
		return err
	}
	// Could be optimized mroe in sake's of Single Responsability
	adminEmail, err := rcv.DB.GetConfigByKey(ctx, "admin_email")
	if err != nil {
		return err
	}
	if err := rcv.Mailer.Send(ctx, &data.Email{
		Recipients: []string{adminEmail},
		Subject:    "New form filled",
		BodyText:   fmt.Sprintf("message: \"%s\" from: %s %s", form.Message, form.FirstName, form.Email),
	}); err != nil {
		return err
	}
	if err := rcv.Mailer.Send(ctx, &data.Email{
		Recipients: []string{form.Email},
		Subject:    "We've recevied your form",
		BodyText:   fmt.Sprintf("Your message \"%s\" has been received", form.Message),
	}); err != nil {
		return err
	}
	return c.JSON(200, form)
}
