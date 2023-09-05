package routes

import (
	"github.com/compy/finder.sh/pkg/context"
	"github.com/compy/finder.sh/pkg/controller"
	"github.com/compy/finder.sh/pkg/msg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	account struct {
		controller.Controller
	}

	accountSettingsForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password1"`
		ConfirmPassword string `form:"password2" validate:"eqfield=Password"`
		Submission      controller.FormSubmission
	}
)

func (c *account) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "account"
	page.Form = accountSettingsForm{
		Name:  page.AuthUser.Name,
		Email: page.AuthUser.Email,
	}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*accountSettingsForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *account) Post(ctx echo.Context) error {
	var form accountSettingsForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse account settings form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		msg.Danger(ctx, "There were errors with your submission. Please check the fields below for specific errors.")
		return c.Get(ctx)
	}
	user, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err != nil {
		msg.Danger(ctx, "Could not fetch user information")
		return c.Get(ctx)
	}
	err = user.Update().SetName(form.Name).SetEmail(form.Email).Exec(ctx.Request().Context())
	if err != nil {
		msg.Danger(ctx, "Could not update your account information.")
		log.Errorf("Could not update account information: %v", err)
		return c.Get(ctx)
	}
	if form.Password != "" && form.ConfirmPassword != "" {
		pwHash, err := c.Container.Auth.HashPassword(form.Password)
		if err != nil {
			return c.Fail(err, "unable to hash password")
		}
		err = user.Update().SetPassword(pwHash).Exec(ctx.Request().Context())
		if err != nil {
			return c.Fail(err, "could not update password")
		}
	}
	/*
		// Hash the password
		pwHash, err := c.Container.Auth.HashPassword(form.Password)
		if err != nil {
			return c.Fail(err, "unable to hash password")
		}

		// Attempt creating the user
		u, err := c.Container.ORM.User.
			Create().
			SetName(form.Name).
			SetEmail(form.Email).
			SetPassword(pwHash).
			Save(ctx.Request().Context())

		switch err.(type) {
		case nil:
			ctx.Logger().Infof("user created: %s", u.Name)
		case *ent.ConstraintError:
			msg.Warning(ctx, "A user with this email address already exists. Please log in.")
			return c.Redirect(ctx, "login")
		default:
			return c.Fail(err, "unable to create user")
		}

		// Log the user in
		err = c.Container.Auth.Login(ctx, u.ID)
		if err != nil {
			ctx.Logger().Errorf("unable to log in: %v", err)
			msg.Info(ctx, "Your account has been created.")
			return c.Redirect(ctx, "login")
		}

		msg.Success(ctx, "Your account has been created. You are now logged in.")

		// Send the verification email
		c.sendVerificationEmail(ctx, u)
	*/

	msg.Success(ctx, "Account settings updated")
	return c.Redirect(ctx, "account")
}
