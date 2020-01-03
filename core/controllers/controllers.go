package controllers

import (
	"net/http"
	"strconv"

	"github.com/jesusantguerrero/goblog/core/models"
	"github.com/labstack/echo"
)

type Controller struct {
	api          *echo.Echo
	model        *models.Model
	resourceName string
}

func (c *Controller) Boot(api *echo.Echo, model *models.Model, resourceName string) {
	c.api = api
	c.model = model
	c.resourceName = resourceName
	c.Routes()
}

// Routes - set routes
func (ctrl Controller) Routes() *echo.Echo {

	ctrl.api.Add("GET", "/api/v1/"+ctrl.resourceName, ctrl.GetResources)

	ctrl.api.Add("GET", "/api/v1/"+ctrl.resourceName+"/:id", ctrl.GetResource)

	ctrl.api.Add("POST", "/api/v1/"+ctrl.resourceName, ctrl.CreateResource)

	ctrl.api.Add("DELETE", "/api/v1/"+ctrl.resourceName+"/:id", ctrl.DeleteResource)

	ctrl.api.Add("PUT", "/api/v1/"+ctrl.resourceName+"/:id", ctrl.UpdateResource)

	return ctrl.api
}

func (ctrl Controller) GetResources(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.model.Get())
}

func (ctrl Controller) GetResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, ctrl.model.GetOne(id))
}

func (ctrl Controller) CreateResource(c echo.Context) error {
	post := ctrl.model.Resource
	if err := c.Bind(post); err != nil {
		return err
	}

	ctrl.model.Save(post)
	return c.JSON(http.StatusCreated, post)
}

func (ctrl Controller) DeleteResource(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctrl.model.Delete(id)
	return c.String(http.StatusOK, "post deleted")
}

func (ctrl Controller) UpdateResource(c echo.Context) error {
	post := ctrl.model.Resource
	if err := c.Bind(post); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, ctrl.model.Update(post, id))
}
