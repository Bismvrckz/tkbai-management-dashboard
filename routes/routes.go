package routes

import (
	"github.com/gorilla/csrf"
	"github.com/labstack/echo/v4"
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	// CSRF
	csrfMid := csrf.Protect([]byte(config.CsrfToken))

	// Public
	public := ein.Tkbai.Group(config.AppPrefix, handler.PublicMiddleware)
	public.Use(echo.WrapMiddleware(csrfMid))
	public.GET("/dashboard", handler.PublicDashboardView)
	public.GET("/certificate/:id", handler.PublicCertificateDetail)
	public.POST("/detail", handler.PublicCertificateDetail)

	// Admin
	admin := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminMiddleware)
	admin.Use(echo.WrapMiddleware(csrfMid))

	admin.GET("/data/toefl/id/:id/name/:certHolder", handler.GetToeflCertificateByID)
	admin.GET("/dashboard", handler.AdminDashboardView)
	admin.GET("/add/csv", handler.AdminInputView)
	admin.POST("/add/csv", handler.AdminUploadCSVCertificate)
	admin.GET("/login", handler.AdminLoginView)
	admin.GET("/logout", handler.AdminLogout)
	admin.POST("/login", handler.AdminLogin)
	//api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	//api.POST("/admin/data/toefl/csv", handler.AdminUploadCSVCertificate)

	// Certificate
	public.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
