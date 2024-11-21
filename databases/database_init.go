package databases

import (
	"tkbai/config"
)

func ConnectTkbaiDatabase() (err error) {
	cmsDB, err := config.TkbaiDbConnection()
	if err != nil {
		return err
	}

	err = cmsDB.Ping()
	if err != nil {
		return err
	}

	DbTkbaiInterface = &TkbaiDbImplement{ConnectTkbaiDB: cmsDB}

	return err
}
