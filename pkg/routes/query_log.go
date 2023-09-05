package routes

import (
	"github.com/compy/finder.sh/pkg/controller"

	"github.com/labstack/echo/v4"
)

type (
	querylog struct {
		controller.Controller
	}
)

func (c *querylog) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "querylog"

	return c.RenderPage(ctx, page)
}
