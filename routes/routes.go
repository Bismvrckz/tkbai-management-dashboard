package routes

import (
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	// Public
	web := ein.Tkbai.Group(config.AppPrefix, handler.PublicMiddleware)
	web.GET("/dashboard", handler.PublicDashboardView)
	web.GET("/certificate/:id", handler.PublicCertificateDetail)

	// Admin
	adminDash := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminMiddleware)
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
