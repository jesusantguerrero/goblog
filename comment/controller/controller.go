package controller

import (
	"net/http"
	"strconv"

	"github.com/jesusantguerrero/goblog/comment/models"
	"github.com/labstack/echo"
)

var model = models.NewModel()

// Routes - routes
func Routes(api *echo.Echo) *echo.Echo {

	api.Add("GET", "/api/v1/comments", GetResources)

	api.Add("GET", "/api/v1/comments/:id", GetResource)

	api.Add("POST", "/api/v1/comments", CreateResource)

	api.Add("DELETE", "/api/v1/comments/:id", DeleteResource)

	api.Add("PUT", "/api/v1/comments/:id", UpdateResource)

	return api
}

// GetResources : this is a get resource
func GetResources(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Get())
}

func GetResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, model.GetOne(id))
}

func CreateResource(c echo.Context) error {
	post := &models.Comment{}
	if err := c.Bind(post); err != nil {
		return err
	}

	newComment := model.Save(post)
	return c.JSON(http.StatusCreated, newComment)
}

func DeleteResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	model.Delete(id)
	return c.String(http.StatusOK, "post deleted")
}

func UpdateResource(c echo.Context) error {
	post := &models.Comment{}
	if err := c.Bind(post); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, model.Update(post, id))
}
