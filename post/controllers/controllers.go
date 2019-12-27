package controllers

import (
	"net/http"
	"strconv"

	"github.com/jesusantguerrero/beeblog/post/models"
	"github.com/labstack/echo"
)

var index = 1
var posts []*models.Post

func Routes(api *echo.Echo) *echo.Echo {

	api.Add("GET", "/api/v1/posts", func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.Get())
	})

	api.Add("GET", "/api/v1/posts/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetOne(id))
	})

	api.Add("POST", "/api/v1/posts", func(c echo.Context) error {
		post := &models.Post{}
		if err := c.Bind(post); err != nil {
			return err
		}

		models.Save(post)
		return c.JSON(http.StatusCreated, post)
	})

	api.Add("DELETE", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "post deleted")
	})

	api.Add("PUT", "/api/v1/posts/:id", func(c echo.Context) error {
		// id, _ := strconv.Atoi(c.Param("id"))
		return c.String(http.StatusOK, "post updated")
	})

	return api
}
