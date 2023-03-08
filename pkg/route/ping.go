package roue

import "github.com/labstack/echo/v4"

// PingHandler godoc
// @summary Ping
// @tags ping
// @description Ping checking for the service
// @id PingHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /ping [get]
func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}
