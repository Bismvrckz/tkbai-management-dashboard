package handler

import (
	"encoding/json"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai/config"
)

var (
	loggers = config.Log
)

type TokenStruct struct {
	AccessToken  string
	Expiry       string
	Message      string
	RefreshToken string
	IdToken      string
}

func AdminGetCookieMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		loggers.Debug().Msg("==========================> Start		| GetCookieMiddleware")
		htmlData := make(map[string]interface{})

		loggers.Debug().Any("config.WebPrefix", config.AppPrefix).Msg("PATH")
		loggers.Debug().Any("ctx.Path()", ctx.Path()).Msg("PATH")
		if !strings.Contains(ctx.Path(), config.AppPrefix) {
			return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/")
		}

		accessToken, err := ReadCookie(ctx, "accessToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		refreshToken, err := ReadCookie(ctx, "refreshToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		idToken, err := ReadCookie(ctx, "idToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		expiry, err := ReadCookie(ctx, "expiry")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		loggers.Debug().Any("accessToken", accessToken != nil).Msg("COOKIE")
		loggers.Debug().Any("refreshToken", refreshToken != nil).Msg("COOKIE")
		loggers.Debug().Any("idToken", idToken != nil).Msg("COOKIE")
		loggers.Debug().Any("expiry", expiry != nil).Msg("COOKIE")

		htmlData["refreshToken"] = refreshToken.Value
		htmlData["accessToken"] = accessToken.Value
		htmlData["idToken"] = idToken.Value
		htmlData["expiry"] = expiry.Value

		htmlData["baseURL"] = config.AppPrefix
		htmlData["apiPrefix"] = config.ApiPrefix
		htmlData["apiHost"] = config.APIHost
		htmlData["prefix"] = config.AppPrefix

		ctx.Set("htmlData", htmlData)

		loggers.Debug().Msg("==========================> Done		| GetCookieMiddleware")
		return next(ctx)
	}
}

func AdminValidateTokenMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		loggers.Debug().Msg("==========================> Start		| ValidateTokenMiddleware")

		htmlData := ctx.Get("htmlData").(map[string]interface{})

		url := config.APIHost + config.ApiPrefix + "/entry/validate"

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			loggers.Err(err).Msg("Error creating request")
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", htmlData["accessToken"].(string))
		req.AddCookie(&http.Cookie{Name: "refreshToken", Value: htmlData["refreshToken"].(string)})
		req.AddCookie(&http.Cookie{Name: "expiry", Value: htmlData["expiry"].(string)})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			loggers.Err(err).Msg("Error sending request")
			return err
		}

		loggers.Debug().Str("Response Status", resp.Status).Msg("")
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			DeleteCookie(ctx, "accessToken")
			DeleteCookie(ctx, "refreshToken")
			DeleteCookie(ctx, "idToken")
			DeleteCookie(ctx, "expiry")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		var result struct {
			ResponseCode    string
			ResponseMessage string
			AdditionalInfo  TokenStruct
		}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}

		loggers.Debug().Any("AdditionalInfo.Message", result.AdditionalInfo.Message).Msg("")

		err = resp.Body.Close()
		if err != nil {
			loggers.Err(err).Msg("Error closing body")
			return err
		}

		WriteCookie(ctx, "accessToken", result.AdditionalInfo.AccessToken, config.AppPrefix, 24)
		WriteCookie(ctx, "refreshToken", result.AdditionalInfo.RefreshToken, config.AppPrefix, 24)
		WriteCookie(ctx, "expiry", result.AdditionalInfo.Expiry, config.AppPrefix, 24)

		htmlData["refreshToken"] = result.AdditionalInfo.RefreshToken
		htmlData["accessToken"] = result.AdditionalInfo.AccessToken
		htmlData["expiry"] = result.AdditionalInfo.Expiry

		ctx.Set("htmlData", htmlData)

		loggers.Debug().Msg("==========================> Done		| ValidateTokenMiddleware")
		return next(ctx)
	}
}

func WebMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {

		data := map[string]interface{}{}

		if ctx.Path() != config.AppPrefix+"/login/admin" && strings.Contains(ctx.Path(), "admin") {
			sess, err := session.Get(config.SessionCookieName, ctx)
			if err != nil {
				return err
			}

			if sess.Values["UserEmail"] == nil {
				sess.Options.MaxAge = -1
				sess.Save(ctx.Request(), ctx.Response())
				return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/login/admin")
			}

			data["UserEmail"] = sess.Values["UserEmail"].(string)
		}

		data["webPublicPrefix"] = config.WebPublicPrefix
		data["appPrefix"] = config.AppPrefix
		data["apiHost"] = config.APIHost
		data["apiPrefix"] = config.ApiPrefix

		ctx.Set("data", data)
		return next(ctx)
	}
}
