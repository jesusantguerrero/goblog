package main

import (
	"net/http"
	"os"

	"github.com/jesusantguerrero/goblog/post/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	api := echo.New()
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api.GET("/api/v1/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is the status page")
	})

	api.POST("/api/v1/testpost", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is the post")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	controllers.Routes(api)
	api.Logger.Fatal(api.Start(":" + port))
}
