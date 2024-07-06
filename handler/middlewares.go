package handler

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai/config"
)

type TokenStruct struct {
	AccessToken  string
	Expiry       string
	Message      string
	RefreshToken string
	IdToken      string
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if strings.Contains(ctx.Path(), "*") {
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/dashboard")
		}

		data := map[string]interface{}{}

		sess, err := session.Get(config.SessionCookieName, ctx)
		if err != nil {
			return err
		}

		if sess.Values["UserEmail"] == nil && ctx.Path() != config.AppPrefix+"/admin/login" {
			sess.Options.MaxAge = -1
			sess.Save(ctx.Request(), ctx.Response())
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/login")
		}

		//TODO: possible nil pointer
		data["UserEmail"] = sess.Values["UserEmail"].(string)
		data["webPublicPrefix"] = config.WebPublicPrefix
		data["appPrefix"] = config.AppPrefix
		data["apiHost"] = config.APIHost
		data["apiPrefix"] = config.ApiPrefix
		data["styleNonce"] = config.StyleSrcNonce
		data["scriptNonce"] = config.ScriptSrcNonce

		ctx.Set("data", data)
		return next(ctx)
	}
}

func PublicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if strings.Contains(ctx.Path(), "*") {
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/dashboard")
		}
		data := map[string]interface{}{}

		data["webPublicPrefix"] = config.WebPublicPrefix
		data["appPrefix"] = config.AppPrefix
		data["apiHost"] = config.APIHost
		data["apiPrefix"] = config.ApiPrefix
		data["styleNonce"] = config.StyleSrcNonce
		data["scriptNonce"] = config.ScriptSrcNonce

		ctx.Set("data", data)
		return next(ctx)
	}
}
