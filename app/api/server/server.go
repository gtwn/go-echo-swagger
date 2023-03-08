package server

import (
	"github.com/gtwn/go-echo-swagger/app/api/route"
	"github.com/labstack/echo/v4"
)

func Configure() func(*echo.Group) error {
	return func(g *echo.Group) error {

		g.GET("/healthcheck", route.HealthCheck())

		return nil
	}
}
