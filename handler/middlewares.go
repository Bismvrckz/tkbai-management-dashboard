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

func WebMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		data := map[string]interface{}{}

		if ctx.Path() != config.AppPrefix+"/admin/login" && strings.Contains(ctx.Path(), "admin") {
			sess, err := session.Get(config.SessionCookieName, ctx)
			if err != nil {
				return err
			}

			if sess.Values["UserEmail"] == nil {
				sess.Options.MaxAge = -1
				sess.Save(ctx.Request(), ctx.Response())
				return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/login")
			}

			data["UserEmail"] = sess.Values["UserEmail"].(string)
		}

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
