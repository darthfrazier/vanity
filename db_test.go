package vanity

import (
	"testing"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"
)

func TestDB(t *testing.T) {
	// Initialize DB connection
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "vanity-prototype")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	db, err := newDatastoreDB(client)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	o := &Outfit{
		State: "WEARING",
		name:  "test outfit",
		class: "B",
	}

	id, err := db.AddOutfit(o)
	if err != nil {
		t.Fatal(err)
	}

	o.ID = id
	o.Style = "goth"
	if err := db.UpdateOutfit(o); err != nil {
		t.Error(err)
	}

	retOutfit, err := db.GetOutfit(id)
	if err != nil {
		t.Error(err)
	}
	if got, want := retOutfit.Style, o.Style; got != want {
		t.Errorf("Update style: got %q, want %q", got, want)
	}

	if err := db.DeleteOutfit(id); err != nil {
		t.Error(err)
	}

	if _, err := db.GetOutfit(id); err == nil {
		t.Error("want non-nil err")
	}
}
