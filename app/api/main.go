package main

import (
	"gtwn/go-echo-swagger/app/api/server"
	_ "gtwn/go-echo-swagger/app/api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1234
// @BasePath /api

func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Configure()(e.Group("/api"))

	e.Logger.Fatal(e.Start(":1234"))
}
