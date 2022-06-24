package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/obegendi/go-ddd/pkg/customer"
)

func Register(service customer.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}
