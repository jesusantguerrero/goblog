package main

import (
	"net/http"
	"os"

	Comment "github.com/jesusantguerrero/goblog/comment/controller"
	Post "github.com/jesusantguerrero/goblog/post/controllers"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// api.Use(middleware.Logger())
	// api.Use(middleware.Recover())
	PostConstroller := Post.Controller{}
	CommentController := Comment.Controller{}

	CommentController.LocalBoot(api)
	PostConstroller.LocalBoot(api)
	api.Logger.Fatal(api.Start(":" + port))
}
