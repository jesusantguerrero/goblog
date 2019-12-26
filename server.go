package main

import (
	"net/http"

	"github.com/jesusantguerrero/beeblog/post/controllers"
	"github.com/labstack/echo"
)

func main() {
	api := echo.New()
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api.GET("/api/v1/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is the status page")
	})

	controllers.Routes(api)

	api.Logger.Fatal(api.Start(":1323"))
}
