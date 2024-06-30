package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"tkbai/config"
	"tkbai/databases"
	"tkbai/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminLoginView(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})
	return ctx.Render(http.StatusOK, "admin.login", data)
}

func AdminLogin(ctx echo.Context) (err error) {
	//data := ctx.Get("data").(map[string]interface{})
	var body models.Login
	err = ctx.Bind(&body)
	if err != nil {
		return err
	}

	result, err := databases.DbTkbaiInterface.GetUserByEmail(body.Email)
	if err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write([]byte(body.Password))
	shaPassword := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	if result.Password.String != shaPassword {
		return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/login/admin")
	}

	sess, err := session.Get(config.SessionCookieName, ctx)
	if err != nil {
		return err
	}
	sess.Values["UserEmail"] = result.Email.String

	sess.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/dash")
}

func AdminDashboardView(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})

	result, err := databases.DbTkbaiInterface.ViewToeflDataBulk()
	if err != nil {
		return err
	}

	data["listData"] = result

	return ctx.Render(http.StatusOK, "admin.dashboard", data)
}
