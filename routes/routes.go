package routes

import (
	"tkbai/config"
	"tkbai/handler"

	"github.com/gorilla/csrf"
	"github.com/labstack/echo/v4"
)

func BuildRoutes(ein *config.Apps) {
	// Public
	web := ein.Tkbai.Group(config.AppPrefix, handler.PublicMiddleware)
	web.GET("/dashboard", handler.PublicDashboardView)
	web.GET("/certificate/:id", handler.PublicCertificateDetail)
	web.POST("/detail", handler.PublicCertificateDetail)

	// Admin
	adminDash := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminMiddleware)
	csrfMid := csrf.Protect([]byte(config.CsrfToken))
	web.Use(echo.WrapMiddleware(csrfMid))

	adminDash.GET("/data/toefl/id/:id/name/:certHolder", handler.GetToeflCertificateByID)
	adminDash.GET("/dashboard", handler.AdminDashboardView)
	adminDash.GET("/add/csv", handler.AdminInputView)
	adminDash.POST("/add/csv", handler.AdminUploadCSVCertificate)
	adminDash.GET("/login", handler.AdminLoginView)
	adminDash.GET("/logout", handler.AdminLogout)
	adminDash.POST("/login", handler.AdminLogin)
	//api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	//api.POST("/admin/data/toefl/csv", handler.AdminUploadCSVCertificate)

	// Certificate
	web.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
