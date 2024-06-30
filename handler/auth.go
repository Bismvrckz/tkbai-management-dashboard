package handler

import (
	"errors"
	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
	"time"
	"tkbai/config"
	"tkbai/models"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     config.IAMClientID,
		ClientSecret: config.IAMClientSecret,
		RedirectURL:  config.IAMLoginRedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
)

func LoginOIDC(ctx echo.Context) (err error) {
	ctxCore := ctx.Request().Context()
	provider, err := oidc.NewProvider(ctxCore, config.IAMDockerConfigURL)
	if err != nil {
		config.LogErr(err, "NewProvider Error")
		return err
	}
	oauth2Config.Endpoint = provider.Endpoint()

	oidcConfig := &oidc.Config{
		ClientID: config.IAMClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	rawAccessToken := ctx.Request().Header.Get("Authorization")

	if rawAccessToken == "" {
		return ctx.Redirect(http.StatusFound, strings.Replace(oauth2Config.AuthCodeURL(config.IAMState), "keycloak", "localhost", 1))
	}
	parts := strings.Split(rawAccessToken, " ")

	if len(parts) != 2 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var _, errVerifier = verifier.Verify(ctxCore, parts[1])
	if errVerifier != nil {
		return ctx.Redirect(http.StatusFound, strings.Replace(oauth2Config.AuthCodeURL(config.IAMState), "keycloak", "localhost", 1))
	}

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  "",
		ResponseMessage: "success",
	})
}

func LoginCallbackOIDC(ctx echo.Context) (err error) {
	//loggers := logging.Log

	if ctx.Request().URL.Query().Get("state") != config.IAMState {
		return ctx.HTML(http.StatusBadRequest, "state did not match")
	}

	ctxCore := ctx.Request().Context()
	provider, err := oidc.NewProvider(ctxCore, config.IAMDockerConfigURL)
	if err != nil {
		return err
	}

	oauth2Config.Endpoint = provider.Endpoint()

	oauth2Token, err := oauth2Config.Exchange(ctxCore, ctx.Request().URL.Query().Get("code"))
	if err != nil {
		return errors.New("Failed to exchange token: " + err.Error())
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("no id_token field in oauth2 token")
	}

	//oidcConfig := &oidc.Config{
	//	ClientID: config.IAMClientID,
	//}
	//verifier := provider.Verifier(oidcConfig)
	//idToken, err := verifier.Verify(ctxCore, rawIDToken)
	//if err != nil {
	//	fmt.Printf("idToken %v\n", idToken)
	//	return errors.New("Failed to verify ID Token: " + err.Error())
	//}
	//
	//resp := struct {
	//	OAuth2Token   *oauth2.Token
	//	IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	//}{
	//	oauth2Token,
	//	new(json.RawMessage),
	//}
	//if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
	//	return ctx.HTML(http.StatusInternalServerError, err.Error())
	//}

	tokenExp := GenerateJwtString(jwt.MapClaims{
		"tokenExp": oauth2Token.Expiry.Format("2006-01-02T15:04:05.999999-07:00"),
	})

	return ctx.HTML(http.StatusOK,
		`<!DOCTYPE html>
              <html>
              	<head>
              		<meta http-equiv="refresh" content="0; url='`+config.IAMLoginRedirect302URL+`'" />
              	</head>
              	<body>
              		<script type="text/javascript">
              			function setCookie(name, value, days, path) {
              				var expires = "";
              				if (days) {
              					var date = new Date();
              					date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
              					expires = "; expires=" + date.toUTCString();
							}
              				document.cookie = name + "=" + value + expires + "; path=" + path;
						}

              			setCookie("accessToken", "`+oauth2Token.AccessToken+`", 1, "`+config.AppPrefix+`");
              			setCookie("refreshToken", "`+oauth2Token.RefreshToken+`", 1, "`+config.AppPrefix+`");
              			setCookie("expiry", "`+tokenExp+`", 1, "`+config.AppPrefix+`");
              			setCookie("idToken", '`+rawIDToken+`', 1, "`+config.AppPrefix+`");
              		</script>
              	</body>
              </html>	
              `)
}

func LogoutOIDC(ctx echo.Context) (err error) {
	idToken := ctx.QueryParam("idToken")

	logoutUrl := config.IAMConfigURL + "/protocol/openid-connect/logout"
	logoutRedirectUrl := config.IAMLogoutRedirectURL

	return ctx.Redirect(http.StatusSeeOther, logoutUrl+"?post_logout_redirect_uri="+logoutRedirectUrl+"&client_id="+config.IAMClientID+"&id_token_hint="+idToken)
}

func LogoutCallbackOIDC(ctx echo.Context) (err error) {
	return ctx.HTML(http.StatusOK,
		`
			<!DOCTYPE html>
			<html>
				<head>
					<meta http-equiv="refresh" content="0; url='`+config.AdminLoginURL+`'" />
				</head>
				<body></body>
			</html>	
			`)
}

func ValidateOIDC(ctx echo.Context) (err error) {
	funcName := "ValidateOIDC"
	accessToken := ctx.Request().Header.Get("Authorization")
	refreshToken, err := ctx.Request().Cookie("refreshToken")
	if err != nil {
		return err
	}
	expiry, err := ctx.Request().Cookie("expiry")
	if err != nil {
		return err
	}

	if accessToken == "" || refreshToken.Value == "" || expiry.Value == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			ResponseCode:    "02",
			AdditionalInfo:  nil,
			ResponseMessage: "failed",
		})
	}

	tokenExp, err := ParseJwtString(expiry.Value, "tokenExp")
	if err != nil {
		if strings.Contains(err.Error(), "token signature is invalid: signature is invalid") {
			return ctx.JSON(http.StatusUnauthorized, "Invalid Token")
		}
		return err
	}

	parsedTime, err := time.Parse("2006-01-02T15:04:05.999999-07:00", tokenExp.(string))
	if err != nil {
		return err
	}

	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Value,
		Expiry:       parsedTime,
		TokenType:    "Bearer",
	}

	if !token.Valid() {
		ctxCore := ctx.Request().Context()

		config.LogDbg(funcName, "AccessToken not valid, refreshing")

		provider, err := oidc.NewProvider(ctxCore, config.IAMDockerConfigURL)
		if err != nil {
			config.LogErr(err, "Failed to create oidc provider")
			return err
		}

		oauth2Config.Endpoint = provider.Endpoint()
		tokenSource := oauth2Config.TokenSource(ctxCore, token)
		newToken, err := tokenSource.Token()

		if err != nil {
			config.LogErr(err, "Failed to get token")

			if strings.Contains(err.Error(), "Token is not active") {
				return ctx.NoContent(http.StatusForbidden)
			} else if strings.Contains(err.Error(), "Session not active") {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			return ctx.JSON(http.StatusBadRequest, "Failed to refresh token | "+err.Error())
		}

		expiry := GenerateJwtString(jwt.MapClaims{
			"tokenExp": newToken.Expiry.Format("2006-01-02T15:04:05.999999-07:00"),
		})

		return ctx.JSON(http.StatusOK, models.Response{
			ResponseCode: "00",
			AdditionalInfo: map[string]interface{}{
				"accessToken":  newToken.AccessToken,
				"refreshToken": newToken.RefreshToken,
				"message":      "Token refreshed using refresh token.",
				"expiry":       expiry,
			},
			ResponseMessage: "success",
		})
	} else {
		config.LogDbg(funcName, "AccessToken still valid")

		return ctx.JSON(http.StatusOK, models.Response{
			ResponseCode: "00",
			AdditionalInfo: map[string]interface{}{
				"accessToken":  accessToken,
				"refreshToken": refreshToken.Value,
				"message":      "Token is still valid.",
				"expiry":       expiry.Value,
			},
			ResponseMessage: "success",
		})
	}
}
