package databases

import (
	"database/sql"
	"tkbai/config"
)

type ToeflCertificate struct {
	ID            sql.NullInt64  `json:"id" db:"id"`
	TestID        sql.NullString `json:"testID" db:"test_id"`
	Name          sql.NullString `json:"name" db:"name"`
	StudentNumber sql.NullString `json:"studentNumber" db:"student_number" `
	Major         sql.NullString `json:"major" db:"major"`
	DateOfTest    sql.NullString `json:"dateOfTest" db:"date_of_test"`
	ToeflScore    sql.NullString `json:"toeflScore" db:"toefl_score"`
	InsertDate    sql.NullTime   `json:"insertDate" db:"insert_date"`
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateToeflCertificate(certificate ToeflCertificate) (err error) {
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES (?,?,?,?,?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, certificate.TestID, certificate.Name, certificate.StudentNumber, certificate.Major, certificate.DateOfTest, certificate.ToeflScore)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *TkbaiDbImplement) CountToeflDataAll() (result int64, err error) {
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) DeleteALlCertificate() (err error) {
	query := "DELETE FROM tkbai_data "
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataAll(start, length string) (result []ToeflCertificate, err error) {
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, length, start)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataByIDAndName(certificateId, certificateHolder string) (result ToeflCertificate, err error) {
	query := `SELECT * FROM tkbai_data WHERE test_id = ? AND name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, certificateId, certificateHolder)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataByIdOrName(credential string) (result ToeflCertificate, err error) {
	query := `SELECT * FROM tkbai_data WHERE test_id = ? OR name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, credential, credential)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataBulk() (result []ToeflCertificate, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Select(&result, "SELECT * FROM tkbai_data")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}
