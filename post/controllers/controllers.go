package controllers

import (
	"net/http"
	"strconv"

	"github.com/jesusantguerrero/goblog/post/models"
	"github.com/labstack/echo"
)

func Routes(api *echo.Echo) *echo.Echo {

	api.Add("GET", "/api/v1/posts", GetResources)

	api.Add("GET", "/api/v1/posts/:id", GetResource)

	api.Add("POST", "/api/v1/posts", CreateResource)

	api.Add("DELETE", "/api/v1/posts/:id", DeleteResource)

	api.Add("PUT", "/api/v1/posts/:id", UpdateResource)

	return api
}

func GetResources(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Get())
}

func GetResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.GetOne(id))
}

func CreateResource(c echo.Context) error {
	post := &models.Post{}
	if err := c.Bind(post); err != nil {
		return err
	}

	models.Save(post)
	return c.JSON(http.StatusCreated, post)
}

func DeleteResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	models.Delete(id)
	return c.String(http.StatusOK, "post deleted")
}

func UpdateResource(c echo.Context) error {
	post := &models.Post{}
	if err := c.Bind(post); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.Update(post, id))
}
