package vanity

import (
	"log"

	"cloud.google.com/go/datastore"

	"gopkg.in/mgo.v2"

	"golang.org/x/net/context"
)

var (
	DB WardrobeDatabase

	// Force import of mgo library.
	_ mgo.Session
)

func init() {
	var err error

	DB, err = configureDatastoreDB("vanity-prototype")

	if err != nil {
		log.Fatal(err)
	}
}

func configureDatastoreDB(projectID string) (WardrobeDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return newDatastoreDB(client)
}
