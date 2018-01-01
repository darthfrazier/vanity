package vanity

import (
	"fmt"
	"time"
)

const (
	// Piece states
	CLOSET          = "CLOSET"
	LAUNDRY         = "LAUNDRY"
	LAUNDRY_PENDING = "LAUNDRY_PENDING"
	WEARING         = "WEARING"
	LIMBO           = "LIMBO"

	// Piece types

	// Piece subtypes
)

// TODO:: Shiity solution for validating state changes
var pieceStates = []string{CLOSET, LAUNDRY, LAUNDRY_PENDING, WEARING, LIMBO}
var pieceTypes = []string{}
var pieceSubTypes = []string{}

type Piece struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	// TODO:: I want these to be an ENUM. Figure out how to do that
	Type    string `json:"piece_type"`
	SubType string `json:"piece_subtype"`

	// Piece qualities
	Price   int    `json:"price"`
	Color   string `json:"color"`
	Pattern string `json:"pattern"`
	Picture string `json:"picture"`

	// Outfits this piece belongs top
	Outfits []int64

	// Piece state tracking
	TimesWornSinceLastWash int       `json:"times_worn_since_last_wash"`
	LastWornDate           time.Time `json:"last_worn_date"`
	// TODO:: I want these to be an ENUM. Figure out how to do that
	State string `json:"state"`
}

func (p *Piece) InitializeOufit() error {
	// TODO
}

// TODO:: pull these constants out to a config file
func (p *Piece) CheckIn(checkinSource string, db WardrobeDatabase) error {
	if checkinSource == "HAMPER-2" {
		p.SetState(LAUNDRY_PENDING, db)
	} else if checkinSource == "MASTER-1" {
		switch p.State {
		case CLOSET:
			p.SetState(WEARING, db)
			// TODO:: create or update active state
		case LAUNDRY, LAUNDRY_PENDING, LIMBO:
			p.SetState(CLOSET, db)
			// TODO:: Trigger inventory service for sorting
		case WEARING:
			p.SetState(CLOSET, db)
			// TODO:: Trigger inventory service for sorting
			// TODO:: create or update active state
		}
	}

	err := db.UpdatePiece(p)
	if err != nil {
		return fmt.Errorf("could not save piece: %v", p)
	}
	return nil
}

func (p *Piece) SetState(state string, db WardrobeDatabase) error {
	if !validateEnum(state, pieceStates) {
		return fmt.Errorf("Invalid state: %v for piece", state)
	}
	p.State = state

	// Propogate state change to outfits
	outfits := db.GetOutfitsByID(p.Outfits)
	for _, outfit := range outfits {
		switch state {
		case LAUNDRY_PENDING, WEARING:
			outfit.State = NOT_WEARABLE
			outfit.MissingPieces = append(outfit.MissingPieces, p.ID)
		case CLOSET:
			outfit.State = WEARABLE
			// TODO:: remove id from missing pieces
		}
	}

	return nil
}
