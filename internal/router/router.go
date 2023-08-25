package router

import (
	"api/internal/handler"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(handlers handler.Handler) *echo.Echo {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.CORS())
	r.Use(middleware.ContextTimeout(time.Second * 60))

	r.POST("/contactForm", handlers.ContactForm)
	r.GET("/items/get_data", handlers.GetItemData)
	r.GET("/orders/new", handlers.CreateOrder)
	r.POST("/orders/info", handlers.OrderInfo)
	r.GET("/health", handlers.HealthCheck)
	r.POST("/stripeWebhook", handlers.StripeWebhook)
	return r
}
