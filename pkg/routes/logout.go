package routes

import (
	"github.com/compy/finder.sh/pkg/controller"
	"github.com/compy/finder.sh/pkg/msg"

	"github.com/labstack/echo/v4"
)

type logout struct {
	controller.Controller
}

func (l *logout) Get(c echo.Context) error {
	if err := l.Container.Auth.Logout(c); err == nil {
		msg.Success(c, "You have been logged out successfully.")
	} else {
		msg.Danger(c, "An error occurred. Please try again.")
	}
	return l.Redirect(c, "home")
}
