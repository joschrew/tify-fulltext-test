package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{http.MethodGet},
	}))

	e.GET("/", hello)
	e.GET("/manifest", manifest)

	e.Start(":6701")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello World")
}

func manifest(c echo.Context) error {
	return c.File("manifest2.json")
}
