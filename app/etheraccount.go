package app

import "github.com/labstack/echo/v5"

func (a App) EtherAddress() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return nil
	}
}
