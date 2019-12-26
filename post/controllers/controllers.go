package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Routes(api *echo.Echo) *echo.Echo {

	api.Add("GET", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "all the posts")
	})

	api.Add("POST", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusCreated, "post created")
	})

	api.Add("DELETE", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "post deleted")
	})

	api.Add("PUT", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "post updated")
	})

	return api
}
