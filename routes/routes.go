package routes

import (
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	web := ein.Tkbai.Group(config.AppPrefix, handler.WebMiddleware)
	web.GET("/", handler.PublicDashboardView)
	web.GET("/certificate/:id/name/:certHolder", handler.PublicCertificateDetail)

	//ADMIN
	//admin := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminGetCookieMid, handler.AdminValidateTokenMid)
	//
	//api := ein.Tkbai.Group(config.AppPrefix + config.ApiPrefix)
	//
	//api.GET("/entry/login", handler.LoginOIDC)
	//api.GET("/auth/loginCallback", handler.LoginCallbackOIDC)
	//api.GET("/auth/logout", handler.LogoutOIDC)
	//api.GET("/auth/logoutCallback", handler.LogoutCallbackOIDC)
	//api.POST("/entry/validate", handler.ValidateOIDC)
	//

	// Admin
	adminDash := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.WebMiddleware)
	adminDash.GET(config.AppPrefix+"/data/toefl/id/:id/name/:certHolder", handler.GetToeflCertificateByID)
	adminDash.GET("/dashboard", handler.AdminDashboardView)
	adminDash.GET("/add/csv", handler.AdminInputView)
	adminDash.POST("/add/csv", handler.AdminUploadCSVCertificate)
	adminDash.GET("/login", handler.AdminLoginView)
	adminDash.POST("/login", handler.AdminLogin)
	//api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	//api.POST("/admin/data/toefl/csv", handler.AdminUploadCSVCertificate)
	//
	//// Certificate
	web.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
