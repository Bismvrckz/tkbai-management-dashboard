package routes

import (
	"github.com/gorilla/csrf"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	// CSRF
	csrfMid := csrf.Protect([]byte(config.CsrfToken))

	// Public
	public := ein.Tkbai.Group(config.AppPrefix,
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) (err error) {
				if strings.Contains(ctx.Path(), "*") {
					return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/dashboard")
				}
				return next(ctx)
			}
		})
	public.Use(echo.WrapMiddleware(csrfMid))
	public.GET("/dashboard", handler.PublicDashboardView)
	public.GET("/certificate/:id", handler.PublicStudentDetailView)
	public.POST("/detail", handler.PublicStudentDetailView)

	// Admin
	admin := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminMiddleware)
	admin.Use(echo.WrapMiddleware(csrfMid))

	admin.GET("/dashboard", handler.AdminDashboardView)
	admin.GET("/add/csv", handler.AdminInputView)
	admin.POST("/add/csv", handler.AdminAddStudentBulk)
	admin.POST("/add/student", handler.AdminAddStudent)
	admin.GET("/login", handler.AdminLoginView)
	admin.GET("/logout", handler.AdminLogout)
	admin.POST("/login", handler.AdminLogin)
	admin.POST("/delete/student", handler.AdminDeleteStudent)
}
