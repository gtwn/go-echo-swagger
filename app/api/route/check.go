package route

import "github.com/labstack/echo/v4"

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK JA")
	}
}
