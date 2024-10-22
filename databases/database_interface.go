package databases

type (
	TkbaiInterface interface {
		ViewAllStudentData(start, length string) (result []StudentData, err error)
		ViewStudentDataBulk() (result []StudentData, err error)
		CountAllStudentData() (result int64, err error)
		ViewStudentDataByIDAndName(certificateId, certificateHolder string) (result StudentData, err error)
		ViewStudentDataByIdOrName(credential string) (result StudentData, err error)
		DeleteALlStudentData() (err error)
		CreateStudentData(certificate StudentData) (err error)

		GetUserByEmail(email string) (user TkbaiUser, err error)
	}
)
