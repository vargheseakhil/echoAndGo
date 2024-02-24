package main

import (
	"echoGoApp/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1213"))

}
