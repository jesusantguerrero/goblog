package controllers

import (
	"net/http"

	"github.com/jesusantguerrero/beeblog/post/models"
	"github.com/labstack/echo"
)

var index = 1
var posts []*models.Post

func Routes(api *echo.Echo) *echo.Echo {

	api.Add("GET", "/api/v1/posts", func(c echo.Context) error {
		return c.JSON(http.StatusOK, posts)
	})

	api.Add("POST", "/api/v1/posts", func(c echo.Context) error {
		post := &models.Post{}
		if err := c.Bind(post); err != nil {
			return err
		}
		post.ID = index
		index++
		posts = append(posts, post)
		return c.JSON(http.StatusCreated, post)
	})

	api.Add("DELETE", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "post deleted")
	})

	api.Add("PUT", "/api/v1/posts", func(c echo.Context) error {
		return c.String(http.StatusOK, "post updated")
	})

	return api
}
