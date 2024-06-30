package routes

import (
	"tkbai/config"
	"tkbai/handler"
)

func BuildRoutes(ein *config.Apps) {
	web := ein.Tkbai.Group(config.AppPrefix, handler.WebMiddleware)
	web.GET("/", handler.PublicDashboardView)
	web.GET("/certificate/:id/name/:certHolder", handler.PublicCertificateDetail)
	web.GET("/login/admin", handler.AdminLoginView)
	web.POST("/login/admin", handler.AdminLogin)

	//ADMIN
	//admin := ein.Tkbai.Group(config.AppPrefix+"/admin", handler.AdminGetCookieMid, handler.AdminValidateTokenMid)
	web.GET("/dash", handler.AdminDashboardView)
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
	//api.GET("/admin/data/toefl/id/:id/name/:certHolder", handler.GetToeflCertificateByID)
	//api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	//api.POST("/admin/data/toefl/csv", handler.UploadCSVCertificate)
	//
	//// Certificate
	//api.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
