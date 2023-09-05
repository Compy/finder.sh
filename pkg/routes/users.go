package routes

import (
	"net/http"
	"strings"

	"github.com/compy/finder.sh/ent"
	"github.com/compy/finder.sh/pkg/context"
	"github.com/compy/finder.sh/pkg/controller"
	"github.com/compy/finder.sh/pkg/msg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	users struct {
		controller.Controller
	}

	newUserForm struct {
		Name       string `form:"name" validate:"required"`
		Email      string `form:"email" validate:"required,email"`
		Password   string `form:"password" validate:"required"`
		Submission controller.FormSubmission
	}

	addUserToGroupForm struct {
		UserID  int `form:"userId" validate:"required"`
		GroupID int `form:"groupId" validate:"required"`
	}

	userPageData struct {
		Users []*ent.User
	}
)

func (c *users) ListUsers(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "users"
	page.Form = newUserForm{}
	data := userPageData{}

	if form := ctx.Get(context.FormKey); form != nil {
		log.Info("Got form")
		page.Form = form.(*newUserForm)
	}

	if userList, err := c.Container.ORM.User.Query().Order(ent.Asc("name")).All(ctx.Request().Context()); err == nil {
		data.Users = userList
	} else {
		return c.Fail(err, "Could not list users")
	}

	page.Data = data

	return c.RenderPage(ctx, page)
}

func (c *users) NewUser(ctx echo.Context) error {
	var form newUserForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		log.Error("Form submission had errors")
		return c.ListUsers(ctx)
	}

	// Hash the password
	pwHash, err := c.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return c.Fail(err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := c.Container.ORM.User.
		Create().
		SetName(form.Name).
		SetEmail(strings.ToLower(form.Email)).
		SetPassword(pwHash).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("user created: %s", u.Name)
	case *ent.ConstraintError:
		msg.Danger(ctx, "A user with that email address already exists.")
		ctx.Response().Header().Set("HX-Redirect", "/users")
		return ctx.String(http.StatusBadRequest, "")
	default:
		return c.Fail(err, "unable to create user")
	}

	//return c.ListUsers(ctx)
	msg.Success(ctx, "User account created successfully")
	ctx.Response().Header().Set("HX-Redirect", "/users")
	return ctx.String(http.StatusOK, "")
}
