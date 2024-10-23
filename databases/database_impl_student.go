package databases

import (
	"database/sql"
	"tkbai/config"
)

type StudentData struct {
	ID                   sql.NullInt64  `json:"id" db:"id"`
	StudentID            sql.NullString `json:"studentID" db:"student_id"`
	Name                 sql.NullString `json:"name" db:"name"`
	StudentNumber        sql.NullString `json:"studentNumber" db:"student_number" `
	Major                sql.NullString `json:"major" db:"major"`
	DateOfAdministration sql.NullString `json:"dateOfTest" db:"date_of_administration"`
	InsertDate           sql.NullTime   `json:"insertDate" db:"insert_date"`
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateStudentData(data StudentData) (err error) {
	query := "INSERT INTO tkbai_data (student_id, name, student_number, major, date_of_administration) VALUES (?,?,?,?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.StudentID, data.Name, data.StudentNumber, data.Major, data.DateOfAdministration)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *TkbaiDbImplement) CountAllStudentData() (result int64, err error) {
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) DeleteALlStudentData() (err error) {
	query := "DELETE FROM tkbai_data "
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewAllStudentData(start, length string) (result []StudentData, err error) {
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, length, start)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewStudentDataByIDAndName(studentID, studentName string) (result StudentData, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_id = ? AND name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, studentID, studentName)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewStudentDataByIdOrName(credential string) (result StudentData, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_number = ? OR name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, credential, credential)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewStudentDataBulk() (result []StudentData, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Select(&result, "SELECT * FROM tkbai_data")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}
