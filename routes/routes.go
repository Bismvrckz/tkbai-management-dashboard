package routes

import (
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	web := ein.Tkbai.Group(config.AppPrefix, handler.WebMiddleware)
	web.GET("/", handler.PublicDashboardView)
	web.GET("/certificate/:id/name/:certHolder", handler.PublicCertificateDetail)
	web.GET("/admin/login", handler.AdminLoginView)
	web.POST("/admin/login", handler.AdminLogin)

	//ADMIN
	//admin := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminGetCookieMid, handler.AdminValidateTokenMid)
	web.GET("/admin/dashboard", handler.AdminDashboardView)
	web.GET("/admin/add/csv", handler.AdminInputView)
	web.POST("/admin/add/csv", handler.AdminUploadCSVCertificate)
	//
	//api := ein.Tkbai.Group(config.AppPrefix + config.ApiPrefix)
	//
	//api.GET("/entry/login", handler.LoginOIDC)
	//api.GET("/auth/loginCallback", handler.LoginCallbackOIDC)
	//api.GET("/auth/logout", handler.LogoutOIDC)
	//api.GET("/auth/logoutCallback", handler.LogoutCallbackOIDC)
	//api.POST("/entry/validate", handler.ValidateOIDC)
	//
	//// Admin
	web.GET(config.AppPrefix+"/admin/data/toefl/id/:id/name/:certHolder", handler.GetToeflCertificateByID)
	//api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	//api.POST("/admin/data/toefl/csv", handler.AdminUploadCSVCertificate)
	//
	//// Certificate
	web.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
