package server

import "github.com/labstack/echo/v4"

func HandleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "login")
	}
}
