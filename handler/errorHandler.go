package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai/config"
)

func InitErrHandler(ein *config.Apps) {
	loggers := config.Log

	// return route response
	//var code int
	//var response models.Status

	ein.Tkbai.HTTPErrorHandler = func(err error, ctx echo.Context) {
		loggers.Debug().Msg(err.Error())

		var report *echo.HTTPError
		ok := errors.As(err, &report)

		if ok {
			err = ctx.JSON(report.Code, map[string]interface{}{
				"message": report.Message,
			})
			if err != nil {
				loggers.Error().Err(err).Msg("")
			}
			return
		}

		err = ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": err.Error(),
		})
		if err != nil {
			loggers.Error().Err(err).Msg("")
		}
	}
}
