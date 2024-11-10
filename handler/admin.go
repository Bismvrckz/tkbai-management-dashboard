package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/csv"
	"net/http"
	"strings"
	"tkbai/config"
	"tkbai/databases"
	"tkbai/models"
	webtemplate "tkbai/webTemplate"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminLoginView(ctx echo.Context) (err error) {
	return webtemplate.AdminLogin(ctx.Get("alertMessage")).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func AdminLogin(ctx echo.Context) (err error) {
	var body models.Login
	err = ctx.Bind(&body)
	if err != nil {
		return err
	}

	result, err := databases.DbTkbaiInterface.GetUserByEmail(body.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.Set("alertMessage", "Email tidak terdaftar")
			return AdminLoginView(ctx)
		}
		return err
	}

	hasher := sha256.New()
	hasher.Write([]byte(body.Password))
	shaPassword := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	if result.Password.String != shaPassword {
		ctx.Set("alertMessage", "Password tidak sesuai")
		return AdminLoginView(ctx)
	}

	config.LogTrc("success login")

	sess, err := session.Get(config.SessionCookieName, ctx)
	if err != nil {
		return err
	}
	sess.Values["userEmail"] = result.Email.String

	sess.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   900,
	}

	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/dashboard")
}

func AdminLogout(ctx echo.Context) (err error) {
	sess, _ := session.Get(config.SessionCookieName, ctx)

	sess.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1,
	}

	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/login")
}

func AdminDashboardView(ctx echo.Context) (err error) {
	result, err := databases.DbTkbaiInterface.ViewStudentDataBulk()
	if err != nil {
		return err
	}

	return webtemplate.AdminDashboard("", result).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func AdminInputView(ctx echo.Context) (err error) {
	return webtemplate.AddCSV("").Render(ctx.Request().Context(), ctx.Response().Writer)
}

func AdminUploadCSVCertificate(ctx echo.Context) (err error) {
	file, err := ctx.FormFile("csv")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	err = src.Close()
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(src)
	csvReader.Comma = ','

	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	err = databases.DbTkbaiInterface.DeleteALlStudentData()
	if err != nil {
		return err
	}

	for i, csvRecord := range csvRecords {
		if i == 0 {
			continue
		}

		err = databases.DbTkbaiInterface.CreateStudentData(databases.StudentData{
			StudentID:            sql.NullString{String: strings.ToUpper(csvRecord[0]), Valid: true},
			Name:                 sql.NullString{String: strings.ToUpper(csvRecord[1]), Valid: true},
			StudentNumber:        sql.NullString{String: csvRecord[2], Valid: true},
			Major:                sql.NullString{String: csvRecord[3], Valid: true},
			DateOfAdministration: sql.NullString{String: csvRecord[4], Valid: true},
		})

		if err != nil {
			return err
		}
	}

	return ctx.Redirect(http.StatusSeeOther, config.AppPrefix+"/admin/dashboard")
}
