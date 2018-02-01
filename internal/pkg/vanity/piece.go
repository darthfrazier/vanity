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
	JACKET      = "JACKET"
	TOP         = "TOP"
	BOTTOM      = "BOTTOM"
	SHOES       = "SHOES"
	ACCESSORIES = "ACCESSORIES"

	// Piece subtypes
	// Jacket subtupes
	RAIN_JACKET  = "RAIN_JACKET"
	HEAVY_JACKET = "HEAVY_JACKET"
	LIGHT_JACKET = "LIGHT_JACKET"
)

// TODO:: Shiity solution for validating state changes
var pieceStates = []string{CLOSET, LAUNDRY, LAUNDRY_PENDING, WEARING, LIMBO}
var pieceTypes = []string{JACKET, TOP, BOTTOM, SHOES, ACCESSORIES}
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
	return nil
}

// TODO:: pull these constants out to a config file
func (p *Piece) CheckIn(checkinSource string, db WardrobeDatabase) error {
	var newState string
	if checkinSource == "HAMPER-2" {
		newState = LAUNDRY_PENDING
	} else if checkinSource == "MASTER-1" {
		switch p.State {
		case CLOSET:
			newState = WEARING
			// TODO:: create or update active state
		case LAUNDRY, LAUNDRY_PENDING, LIMBO:
			newState = CLOSET
		case WEARING:
			newState = CLOSET
			// TODO:: create or update active state
		}
	}
	fmt.Println(newState)

	err := p.SetState(newState, db)
	if err != nil {
		return fmt.Errorf("Could not change piece state: %v", err)
	}

	fmt.Println(p)
	err = db.UpdatePiece(p)
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
	if p.Outfits != nil {
		outfits, err := db.GetOutfitsByID(p.Outfits)
		fmt.Print(err)
		if err != nil {
			return fmt.Errorf("Could not retrieve outfits: %v", err)
		}

		for _, outfit := range outfits {
			switch state {
			case LAUNDRY_PENDING, WEARING:
				outfit.State = NOT_WEARABLE
				outfit.MissingPieces = append(outfit.MissingPieces, p.ID)
			case CLOSET:
				outfit.State = WEARABLE
				missingPieces := outfit.MissingPieces
				for i, id := range missingPieces {
					if id == p.ID {
						missingPieces = append(missingPieces[:i], missingPieces[i+1:]...)
						break
					}
				}
			}
		}
		if err := db.UpdateOutfits(outfits); err != nil {
			return fmt.Errorf("Could not updated outfits: %v", err)
		}
	}

	return nil
}
