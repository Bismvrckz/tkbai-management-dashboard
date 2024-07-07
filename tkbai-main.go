package main

import (
	"embed"
	"html/template"
	"net/http"
	"strings"
	"tkbai/config"
	"tkbai/databases"
	"tkbai/handler"
	"tkbai/models"
	"tkbai/routes"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed public
	embedPublic embed.FS

	pages = []string{
		"public/views/*.html",
	}
)

func main() {
	a := new(config.Apps)

	a.Tkbai = echo.New()

	a.Tkbai.Use(middleware.Recover())

	a.Tkbai.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.WebHost, config.APIHost},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))

	sessionStore := sessions.NewCookieStore([]byte(config.AppSessionSecret))
	sessionStore.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	a.Tkbai.Use(session.Middleware(sessionStore))

	a.Tkbai.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))
	a.Tkbai.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_csrf",
		CookieSameSite: http.SameSiteStrictMode,
		CookieHTTPOnly: true,
		CookieSecure:   true,
	}))
	a.Tkbai.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            2592000,
		ContentSecurityPolicy: "default-src 'self' ;font-src 'self' fonts.googleapis.com fonts.gstatic.com; style-src 'nonce-" + config.StyleSrcNonce + "' 'self' fonts.googleapis.com fonts.gstatic.com; script-src 'self' 'nonce-" + config.ScriptSrcNonce + "' ; img-src data://* 'self' www.w3.org ",
	}))

	//'nonce-" + config.StyleSrcNonce + "' 'self' fonts.googleapis.com fonts.gstatic.com

	//logging
	initLoggingMiddleware(a)

	//init handler
	handler.InitErrHandler(a)

	//add routes
	routes.BuildRoutes(a)

	initTemplate(a)

	err := databases.ConnectTkbaiDatabase()
	if err != nil {
		config.LogErr(err, "Error connecting to database")
	}

	a.Tkbai.Logger.Fatal(a.Tkbai.Start(config.SERVERPort))
}

func initLoggingMiddleware(ein *config.Apps) {
	logger := config.Log

	ein.Tkbai.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogURIPath:   true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogRequestID: true,
		LogError:     true,
		LogMethod:    true,
		LogValuesFunc: func(ctx echo.Context, values middleware.RequestLoggerValues) (err error) {
			if strings.Contains(values.URI, "public") {
				return err
			}

			if values.Error != nil {
				logger.Error().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Stack().Err(values.Error).Msg("")
			} else {
				logger.Info().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Msg("Request")
			}
			return nil
		},
	}))
}

// init template
func initTemplate(srv *config.Apps) {
	t := &models.Template{
		Templates: template.Must(template.ParseFS(embedPublic, pages...)),
	}

	srv.Tkbai.Renderer = t
	srv.Tkbai.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(embedPublic),
	}))
}
