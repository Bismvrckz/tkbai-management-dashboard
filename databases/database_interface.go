package databases

type (
	TkbaiInterface interface {
		ViewToeflDataAll(start, length string) (result []ToeflCertificate, err error)
		ViewToeflDataBulk() (result []ToeflCertificate, err error)
		CountToeflDataAll() (result int64, err error)
		ViewToeflDataByIDAndName(certificateId, certificateHolder string) (result ToeflCertificate, err error)
		ViewToeflDataByIdOrName(credential string) (result ToeflCertificate, err error)
		CreateCertificateBulk(certificates []ToeflCertificate) (rowsAffected int64, err error)
		DeleteALlCertificate() (err error)
		CreateToeflCertificate(certificate ToeflCertificate) (err error)

		GetUserByEmail(email string) (user TkbaiUser, err error)
	}
)
