package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	e.GET("/annotations:*", annotation)

	e.Start(":6701")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello World")
}

func annotation(c echo.Context) error {
	path := c.Request().URL.Path
	path = strings.TrimPrefix(path, "/")
	return c.File(fmt.Sprintf("data/%s", path))
}

func manifest(c echo.Context) error {
	symlink := "data/the-used-manifest.json"
	target, err := os.Readlink(symlink)
	if err != nil {
		log.Fatalf("Cannot read symbolic link: %s", err)
	}
	fmt.Println(target)
	c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	return c.File(symlink)
}
