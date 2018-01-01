package vanity_test

import (
	"testing"

	"fmt"
	"github.com/darthfrazier/vanity/internal/pkg/vanity"
)

var getPieceReturn *vanity.Piece

type TestWardrobeDatabase struct{}

func (db *TestWardrobeDatabase) Close() {
	// No op
}

func (db *TestWardrobeDatabase) AddPiece(p *vanity.Piece) (id int64, err error) {
	return 0, nil
}

func (db *TestWardrobeDatabase) GetPiece(id int64) (*vanity.Piece, error) {
	return getPieceReturn, nil
}

func (db *TestWardrobeDatabase) UpdatePiece(p *vanity.Piece) error {
	return nil
}

func (db *TestWardrobeDatabase) DeletePiece(id int64) error {
	return nil
}

func (db *TestWardrobeDatabase) AddOutfit(o *vanity.Outfit) (id int64, err error) {
	return 0, nil
}

func (db *TestWardrobeDatabase) GetOutfit(id int64) (*vanity.Outfit, error) {
	return nil, nil
}

func (db *TestWardrobeDatabase) GetOutfits(state string, class string) ([]*vanity.Outfit, error) {
	return nil, nil
}

func (db *TestWardrobeDatabase) UpdateOutfit(o *vanity.Outfit) error {
	return nil
}

func (db *TestWardrobeDatabase) DeleteOutfit(id int64) error {
	return nil
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertNotNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertError(t *testing.T, err error, message string) {
	assertEqual(t, err.Error(), message, "")
}

// func TestInitiializeOutfitNoPiecesSuccess(t *testing.T) {
// 	outfit := &vanity.Outfit{
// 		State: "WEARING",
// 		Name:  "test outfit",
// 		Class: "POWERBLACK",
// 	}
// 	err := outfit.InitializeOutfit()
// 	assertNotNil(t, err)
// }

// func TestInitiializeOutfitNoPiecesInvalidClassFailure(t *testing.T) {
// 	outfit := &vanity.Outfit{
// 		State: "WEARING",
// 		Name:  "test outfit",
// 		Class: "INVALID",
// 	}
// 	err := outfit.InitializeOutfit()
// 	assertEqual(t, err.Error(), "Invalid class: INVALID for outfit", "")
// }

// func TestInitiializeOutfitInvalidPieceFailure(t *testing.T) {
// 	outfit := &vanity.Outfit{
// 		State:  "WEARING",
// 		Name:   "test outfit",
// 		Class:  "INVALID",
// 		Bottom: 123,
// 	}
// 	err := outfit.InitializeOutfit()
// 	assertError(t, err, "datastoreDB: could not get Piece: datastore: no such entity")
// }

func TestInitiializeOutfitInvalidTopFailure(t *testing.T) {
	var db *TestWardrobeDatabase
	getPieceReturn = &vanity.Piece{
		ID: 123,
	}
	outfit := &vanity.Outfit{
		State: "WEARING",
		Name:  "test outfit",
		Class: "INVALID",
		Tops:  []int64{123, 234},
	}
	err := outfit.InitializeOutfit(db)
	assertError(t, err, "datastoreDB: could not get Piece: datastore: no such entity")
}

func TestInitiializeOutfitInvalidTopFailure2(t *testing.T) {
	var db *TestWardrobeDatabase
	getPieceReturn = &vanity.Piece{
		ID:   223,
		Type: "TOP",
	}
	outfit := &vanity.Outfit{
		State: "WEARING",
		Name:  "test outfit",
		Class: "INVALID",
		Tops:  []int64{123, 234},
	}
	err := outfit.InitializeOutfit(db)
	assertError(t, err, "datastoreDB: could not get Piece: datastore: no such entity")
}
