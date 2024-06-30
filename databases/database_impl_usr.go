package databases

import (
	"database/sql"
	"tkbai/config"
)

type TkbaiUser struct {
	Id         sql.NullInt64  `db:"id"`
	Email      sql.NullString `db:"email"`
	Password   sql.NullString `db:"password"`
	InsertDate sql.NullTime   `db:"insert_date"`
}

func (tkbaiDbImpl *TkbaiDbImplement) GetUserByEmail(email string) (user TkbaiUser, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&user, "SELECT * FROM tkbai_user where email = ?", email)
	if err != nil {
		config.LogErr(err, "Query Error")
		return user, err
	}

	return user, err
}
