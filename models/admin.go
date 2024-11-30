package models

type AddStudentPayload struct {
	StudentName    string `form:"studentName"`
	StudentAddress string `form:"studentAddress"`
	StudentNumber  string `form:"studentNumber"`
	StudentMajor   string `form:"major"`
}

type DeleteStudentPayload struct {
	ID string `form:"id"`
}
