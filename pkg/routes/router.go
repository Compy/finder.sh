package routes

import (
	"net/http"

	"github.com/compy/finder.sh/config"
	"github.com/compy/finder.sh/pkg/controller"
	"github.com/compy/finder.sh/pkg/middleware"
	"github.com/compy/finder.sh/pkg/services"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non static file route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.Logger(),
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.Cache),

		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf,header:X-CSRF-Token",
		}),
	)

	// Base controller
	ctr := controller.NewController(c)

	// Error handler
	err := errorHandler{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Example routes
	navRoutes(c, g, ctr)
	userRoutes(c, g, ctr)
}

func navRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	home := home{Controller: ctr}
	g.GET("/", home.Get, middleware.RequireAuthentication()).Name = "home"

	search := search{Controller: ctr}
	g.GET("/search", search.Get, middleware.RequireAuthentication()).Name = "search"

	datasources := datasources{Controller: ctr}
	g.GET("/datasources", datasources.Get, middleware.RequireAuthentication()).Name = "datasources"
	g.GET("/datasources/add", datasources.Add, middleware.RequireAuthentication()).Name = "datasources.add"
	g.GET("/datasources/add/:datasource", datasources.Configure, middleware.RequireAuthentication()).Name = "datasources.configure"
	g.POST("/datasources/save", datasources.New, middleware.RequireAuthentication()).Name = "datasources.new"
	g.GET("/datasources/reindex/:datasource", datasources.Reindex, middleware.RequireAuthentication()).Name = "datasources.reindex"

	querylog := querylog{Controller: ctr}
	g.GET("/queries", querylog.Get, middleware.RequireGroup("administrators")).Name = "querylog"

	account := account{Controller: ctr}
	g.GET("/account", account.Get, middleware.RequireAuthentication()).Name = "account"
	g.POST("/account", account.Post, middleware.RequireAuthentication()).Name = "account.post"

	contact := contact{Controller: ctr}
	g.GET("/appearance", contact.Get, middleware.RequireAuthentication()).Name = "appearance"
	g.POST("/contact", contact.Post, middleware.RequireAuthentication()).Name = "contact.post"

	users := users{Controller: ctr}
	g.GET("/users", users.ListUsers, middleware.RequireGroup("administrators")).Name = "users.list"
	g.POST("/users", users.NewUser, middleware.RequireGroup("administrators")).Name = "users.new"

	groups := groups{Controller: ctr}
	g.GET("/groups", groups.ListGroups, middleware.RequireGroup("administrators")).Name = "groups.list"
	g.POST("/groups", groups.NewGroup, middleware.RequireGroup("administrators")).Name = "groups.new"
	g.GET("/groups/:group", groups.GetGroup, middleware.RequireGroup("administrators")).Name = "groups.get"
	g.DELETE("/groups/:group", groups.DeleteGroup, middleware.RequireGroup("administrators")).Name = "groups.delete"
	g.POST("/groups/:group/member", groups.AddMember, middleware.RequireGroup("administrators")).Name = "groups.addMember"
	g.DELETE("/groups/:group/member/:userid", groups.RemoveMember, middleware.RequireGroup("administrators")).Name = "groups.removeMember"
}

func userRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	logout := logout{Controller: ctr}
	g.GET("/logout", logout.Get, middleware.RequireAuthentication()).Name = "logout"

	verifyEmail := verifyEmail{Controller: ctr}
	g.GET("/email/verify/:token", verifyEmail.Get).Name = "verify_email"

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	login := login{Controller: ctr}
	noAuth.GET("/login", login.Get).Name = "login"
	noAuth.POST("/login", login.Post).Name = "login.post"

	register := register{Controller: ctr}
	noAuth.GET("/register", register.Get).Name = "register"
	noAuth.POST("/register", register.Post).Name = "register.post"

	forgot := forgotPassword{Controller: ctr}
	noAuth.GET("/password", forgot.Get).Name = "forgot_password"
	noAuth.POST("/password", forgot.Post).Name = "forgot_password.post"

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(c.ORM),
		middleware.LoadValidPasswordToken(c.Auth),
	)
	reset := resetPassword{Controller: ctr}
	resetGroup.GET("/token/:user/:password_token/:token", reset.Get).Name = "reset_password"
	resetGroup.POST("/token/:user/:password_token/:token", reset.Post).Name = "reset_password.post"
}
