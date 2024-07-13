package handler

import (
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai/databases"
)

func PublicDashboardView(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})

	return ctx.Render(http.StatusOK, "public.dashboard", data)
}

func PublicCertificateDetail(ctx echo.Context) (err error) {
	data := ctx.Get("data").(map[string]interface{})
	certificateId := ctx.Param("id")
	decodedId, err := base64.StdEncoding.DecodeString(certificateId)
	if err != nil {
		return err
	}

	result, err := databases.DbTkbaiInterface.ViewToeflDataByID(string(decodedId))
	if err != nil {
		return err
	}

	data["listData"] = result

	return ctx.Render(http.StatusOK, "public.detail.toefl", data)
}
