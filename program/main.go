package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}

	e := echo.New()
	e.GET("/", hello)
	e.Start(port)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
