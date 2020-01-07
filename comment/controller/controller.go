package controller

import (
	"github.com/jesusantguerrero/goblog/comment/models"
	BaseController "github.com/jesusantguerrero/goblog/core/controllers"
	"github.com/labstack/echo"
)

var model = models.NewModel()

type Controller struct {
	BaseController.Controller
}

func (c *Controller) SuperBoot(api *echo.Echo) {
	c.Boot(api, model, "comments")
}
