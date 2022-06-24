package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/obegendi/go-ddd/api/handlers"
	"github.com/obegendi/go-ddd/pkg/customer"
)

func CustomerRouter(app *echo.Echo, service customer.Service) {
	customerGroup := app.Group("/api/v1/customers")
	customerGroup.POST("/register", handlers.Register(service))
}
