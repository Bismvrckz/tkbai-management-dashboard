package handler

import (
	"database/sql"
	"encoding/csv"
	"net/http"
	"strings"
	"tkbai/databases"
	"tkbai/models"

	"github.com/labstack/echo/v4"
)

func GetAllToeflCertificate(ctx echo.Context) (err error) {
	start := ctx.QueryParam("start")
	length := ctx.QueryParam("length")

	result, err := databases.DbTkbaiInterface.ViewToeflDataAll(start, length)
	if err != nil {
		return err
	}

	resultCount, err := databases.DbTkbaiInterface.CountToeflDataAll()
	if err != nil {
		return err

	}

	for i, each := range result {
		dateOfTestSplit := strings.Split(each.DateOfTest.String, " ")
		insertDateSplit := strings.Split(each.InsertDate.String, " ")

		result[i].DateOfTest.String = dateOfTestSplit[0]
		result[i].InsertDate.String = insertDateSplit[0]
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":            result,
		"draw":            ctx.QueryParam("draw"),
		"recordsTotal":    resultCount,
		"recordsFiltered": resultCount,
	})
}

func GetToeflCertificateByID(ctx echo.Context) (err error) {
	certificateId := ctx.Param("id")
	certificateHolder := ctx.Param("certHolder")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByIDAndName(certificateId, certificateHolder)
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}

func UploadCSVCertificate(ctx echo.Context) (err error) {
	file, err := ctx.FormFile("toefl_csv")
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

	//opt := option.WithCredentials(&google.Credentials{
	//	ProjectID:              "tkbai-management-dashboard",
	//	TokenSource:            nil,
	//	JSON:                   nil,
	//	UniverseDomainProvider: nil,
	//})
	//client, err := firestore.NewClient(ctxbg, "tkbai-management-dashboard", opt)
	//if err != nil {
	//	config.LogErr(err, "firestore error")
	//	return echo.ErrInternalServerError
	//}
	//wr, err := client.Doc("tkbai").Create(ctxbg, map[string]interface{}{
	//	"TestID":        csvRecords[0][0],
	//	"Name":          csvRecords[0][1],
	//	"StudentNumber": csvRecords[0][2],
	//	"Major":         csvRecords[0][3],
	//	"DateOfTest":    csvRecords[0][4],
	//	"ToeflScore":    csvRecords[0][5],
	//})
	//if err != nil {
	//	config.LogErr(err, "firestore Doc Create error")
	//	return echo.ErrInternalServerError
	//}
	//fmt.Println(wr.UpdateTime)
	//err = client.Close()
	//if err != nil {
	//	config.LogErr(err, "firestore close error")
	//}

	for i, csvRecord := range csvRecords {
		if i == 0 {
			continue
		}

		_, err := databases.DbTkbaiInterface.CreateToeflCertificate(databases.ToeflCertificate{
			TestID:        sql.NullString{String: csvRecord[0], Valid: true},
			Name:          sql.NullString{String: csvRecord[1], Valid: true},
			StudentNumber: sql.NullString{String: csvRecord[2], Valid: true},
			Major:         sql.NullString{String: csvRecord[3], Valid: true},
			DateOfTest:    sql.NullString{String: csvRecord[4], Valid: true},
			ToeflScore:    sql.NullString{String: csvRecord[5], Valid: true},
		})
		if err != nil {
			return err
		}

	}

	return ctx.JSON(http.StatusOK, "success")
}

func ValidateCertificateByID(ctx echo.Context) error {
	certificateId := ctx.Param("id")
	certificateHolder := ctx.Param("certHolder")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByIDAndName(certificateId, certificateHolder)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}
