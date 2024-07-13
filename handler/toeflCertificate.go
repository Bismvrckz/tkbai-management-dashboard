package handler

import (
	"fmt"
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
		//insertDateSplit := strings.Split(each.InsertDate.Time.String(), " ")

		result[i].DateOfTest.String = dateOfTestSplit[0]
		//result[i].InsertDate.Time. = insertDateSplit[0]
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

func ValidateCertificateByID(ctx echo.Context) error {
	certificateId := ctx.Param("id")
	certificateHolder := ctx.Param("certHolder")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByIDAndName(certificateId, certificateHolder)
	if err != nil {
		return err
	}

	fmt.Printf("result: %v\n", result.ID.Int64)

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}
