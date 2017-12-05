package vanity

import (
	"errors"
	"log"
	"os"

	"cloud.google.com/go/datastore"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/sessions"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	DB          WardrobeDatabase
	OAuthConfig *oauth2.Config

	StorageBucket     *storage.BucketHandle
	StorageBucketName string

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