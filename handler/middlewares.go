package handler

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai/config"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		sess, err := session.Get(config.SessionCookieName, ctx)
		if err != nil {
			return err
		}

		if sess.Values["userEmail"] == nil && ctx.Path() != config.AppPrefix+"/admin/login" {
			sess.Options.MaxAge = -1
			sess.Save(ctx.Request(), ctx.Response())
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/login")
		}

		if strings.Contains(ctx.Path(), "*") {
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/dashboard")
		}

		return next(ctx)
	}
}
