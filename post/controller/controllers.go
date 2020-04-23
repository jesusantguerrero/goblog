package controller

import (
	BaseController "github.com/jesusantguerrero/goblog/core/controllers"
	"github.com/jesusantguerrero/goblog/post/models"
	"github.com/labstack/echo"
)

var model = models.NewModel()

type Controller struct {
	BaseController.Controller
}

// LocalBoot boot the controller
func (c *Controller) LocalBoot(api *echo.Echo) {
	c.Boot(api, model, "posts")
}
