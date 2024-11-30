package databases

import (
	"github.com/jmoiron/sqlx"
)

type (
	TkbaiDbImplement struct {
		ConnectTkbaiDB *sqlx.DB
		Err            error
	}
)

var DbTkbaiInterface *TkbaiDbImplement
