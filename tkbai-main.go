package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
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

	// recover
	a.Tkbai.Use(middleware.Recover())

	// set cors
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

	//Use session middleware
	a.Tkbai.Use(session.Middleware(sessionStore))

	//logging
	initLoggingMiddleware(a)

	//init handler
	handler.InitErrHandler(a)

	//add routes
	routes.BuildRoutes(a)

	initTemplate(a)

	err := databases.ConnectTkbaiDatabase()
	if err != nil {
		log.Fatal(err)
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
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
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
