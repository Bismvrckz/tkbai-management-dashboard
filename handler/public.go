package handler

import (
	"strings"
	"tkbai/databases"
	webtemplate "tkbai/webTemplate"

	"github.com/labstack/echo/v4"
)

func PublicDashboardView(ctx echo.Context) (err error) {
	return webtemplate.PublicDashboard().Render(ctx.Request().Context(), ctx.Response().Writer)
}

func PublicCertificateDetail(ctx echo.Context) (err error) {
	credential := ctx.FormValue("credential")

	result, err := databases.DbTkbaiInterface.ViewStudentDataByIdOrName(strings.ToUpper(credential))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return webtemplate.DetailNotFound().Render(ctx.Request().Context(), ctx.Response().Writer)
		}
		return err
	}

	return webtemplate.StudentDetail(result).Render(ctx.Request().Context(), ctx.Response().Writer)
}
