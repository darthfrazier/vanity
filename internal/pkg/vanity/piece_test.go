package vanity_test

import (
	"github.com/darthfrazier/vanity/internal/pkg/vanity"
	"testing"
)

func TestCheckInPiece(t *testing.T) {
	var db *TestWardrobeDatabase
	piece := &vanity.Piece{
		ID:      123,
		Name:    "MNML Tee",
		Brand:   "MNML",
		Price:   9000,
		Type:    "TOP",
		SubType: "TEE",
		Color:   "Salmon",
		State:   "CLOSET",
	}

	err := piece.CheckIn("MASTER-1", db)
	assertNil(t, err)
	assertEqual(t, piece.State, "WEARING", "Piece state not changed")
}
