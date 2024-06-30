package config

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	dbMaxIdleCons = 10
	dbMaxCons     = 100
)

func TkbaiDbConnection() (db *sqlx.DB, err error) {
	funcName := "TkbaiDbConnection"
	db, err = sqlx.Open("sqlite3", "db/sqlite/tkbai.db")
	if err != nil {
		Log.Err(err).Str("FUNC", funcName).Msg("")
		return db, err
	}

	err = db.Ping()
	if err != nil {
		Log.Err(err).Str("FUNC", funcName).Msg("")
		return db, err
	}

	db.SetMaxOpenConns(dbMaxCons)
	db.SetMaxIdleConns(dbMaxIdleCons)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)
	return db, err
}
