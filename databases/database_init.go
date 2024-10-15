package databases

import (
	"context"
	"tkbai/config"

	"cloud.google.com/go/firestore"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
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

func ConnectTkbaiFirestore() (err error) {
	ctx := context.Background()
	opt := option.WithCredentials(&google.Credentials{
		ProjectID:              "tkbai-management-dashboard",
		TokenSource:            nil,
		JSON:                   nil,
		UniverseDomainProvider: nil,
	})
	client, err := firestore.NewClient(ctx, "tiebing-test1", opt)
	if err != nil {
		config.LogErr(err, "firestore error")
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			config.LogErr(err, "firestore close error")
		}
	}(client)

	return err
}
