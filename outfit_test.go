package vanity

import (
	"testing"

	"cloud.google.com/go/datastore"
)

func TestOufitEncodeDecodeTopsKeys(t *testing.T) {
	outfit := &Outfit{
		State: "WEARING",
		name:  "test outfit",
		class: "B",
	}
	jacketKey := &datastore.Key{
		ID:   123,
		Kind: "Test",
	}

	topsKeys := []*datastore.Key{
		&datastore.Key{
			ID:   1,
			Kind: "Test 1",
		},
		&datastore.Key{
			ID:   2,
			Kind: "Test 2",
		},
		&datastore.Key{
			ID:   3,
			Kind: "Test 3",
		},
	}
	outfit.Jacket = jacketKey
	outfit.EncodeTopsKeys(topsKeys)
	decodedTopsKeys, _ := outfit.DecodeTopsKeys()
	for i, el := range decodedTopsKeys {
		if el != *topsKeys[i] {
			t.Errorf("got %q: wanted %q", &el, topsKeys[i])
		}
	}
}
