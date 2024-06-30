package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai/config"

	"time"
)

func WriteCookie(c echo.Context, key, value, path string, hoursExpire int) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(time.Duration(hoursExpire) * time.Hour)
	cookie.Path = path
	c.SetCookie(cookie)
}

func ReadCookie(c echo.Context, key string) (*http.Cookie, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func DeleteCookie(c echo.Context, key string) {
	cookie, err := ReadCookie(c, key)
	if err != nil {
		config.LogErr(err, "Error delete cookie")
	}
	cookie.Path = config.AppPrefix
	cookie.MaxAge = -1
	c.SetCookie(cookie)
}
