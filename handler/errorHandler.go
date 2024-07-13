package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai/config"
)

func InitErrHandler(ein *config.Apps) {
	ein.Tkbai.HTTPErrorHandler = func(err error, ctx echo.Context) {
		config.Log.Debug().Msg(err.Error())

		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
}
