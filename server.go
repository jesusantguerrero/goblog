package main

import (
	"net/http"

	commentcontroller "github.com/jesusantguerrero/goblog/comment/controller"
	"github.com/jesusantguerrero/goblog/post/controllers"
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

	// api.Use(middleware.Logger())
	// api.Use(middleware.Recover())
	controllers.Routes(api)
	commentcontroller.Routes(api)
	api.Logger.Fatal(api.Start(":1323"))
}
