package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"tkbai/config"
)

type ToeflCertificate struct {
	ID            sql.NullInt64  `json:"id"`
	TestID        sql.NullString `json:"testID"`
	Name          sql.NullString `json:"name"`
	StudentNumber sql.NullString `json:"studentNumber"`
	Major         sql.NullString `json:"major"`
	DateOfTest    sql.NullString `json:"dateOfTest"`
	ToeflScore    sql.NullString `json:"toeflScore"`
	InsertDate    sql.NullString `json:"insertDate"`
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataAll(start, length string) (result []ToeflCertificate, err error) {
	funcName := "ViewToeflDataAll"
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.Query(query, length, start)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	for rows.Next() {
		var each ToeflCertificate
		err = rows.Scan(&each.ID, &each.TestID, &each.Name, &each.StudentNumber, &each.Major, &each.DateOfTest, &each.ToeflScore, &each.InsertDate)
		if err != nil {
			break
		}

		result = append(result, each)
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err != nil {
		config.LogErr(err, "Scan Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataBulk() (result []ToeflCertificate, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, "SELECT * FROM tkbai_data")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CountToeflDataAll() (result int64, err error) {
	funcName := "CountToeflDataAll"
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.Query(query)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataByIDAndName(certificateId, certificateHolder string) (result ToeflCertificate, err error) {
	funcName := "ViewToeflDataByIDAndName"
	query := `SELECT * FROM tkbai_data WHERE test_id = ? AND name = ?`
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.Query(query, certificateId, certificateHolder)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	if !rows.Next() {
		err = errors.New("not found")
		config.LogErr(err, fmt.Sprintf("Test ID %v not found", certificateId))
		return result, echo.ErrNotFound
	}

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.TestID, &result.Name, &result.StudentNumber, &result.Major, &result.DateOfTest, &result.ToeflScore, &result.InsertDate)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateCertificateBulk(certificates []ToeflCertificate) (rowsAffected int64, err error) {
	var args []any
	var parameterString string
	for _, each := range certificates {
		parameterString += "(?,?,?,?,STR_TO_DATE(?, '%d-%b-%y'),?),"
		args = append(args, each.TestID, each.Name, each.StudentNumber, each.Major, each.DateOfTest, each.ToeflScore)
	}

	if len(parameterString) > 1 {
		parameterString = parameterString[:len(parameterString)-1]
	}

	funcName := "CreateCertificateBulk"
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES " + parameterString
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.Exec(query, args...)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	if rowsAffected != 1 {
		err = errors.New(fmt.Sprintf("expected single row affected, got %d rows affected", rows))
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return rowsAffected, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateToeflCertificate(certificate ToeflCertificate) (rowsAffected int64, err error) {
	funcName := "CreateToeflCertificate"
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES (?,?,?,?,STR_TO_DATE(?, '%d-%b-%y'),?)"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.Exec(query, certificate.TestID, certificate.Name, certificate.StudentNumber, certificate.Major, certificate.DateOfTest, certificate.ToeflScore)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	if rowsAffected != 1 {
		err = errors.New(fmt.Sprintf("expected single row affected, got %d rows affected", rows))
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return rowsAffected, err
}
