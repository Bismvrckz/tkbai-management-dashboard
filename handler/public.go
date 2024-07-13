package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"tkbai/databases"
)

func PublicDashboardView(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})
	data["csrf"] = ctx.Get(middleware.DefaultCSRFConfig.ContextKey).(string)

	return ctx.Render(http.StatusOK, "public.dashboard", data)
}

func PublicCertificateDetail(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})
	certificateId := ctx.FormValue("certificateId")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByID(certificateId)
	if err != nil {
		if err.Error() == "not found" {
			return ctx.Render(http.StatusOK, "public.detail.notfound", data)
		}
		return err
	}

	data["listData"] = result

	return ctx.Render(http.StatusOK, "public.detail.toefl", data)
}
