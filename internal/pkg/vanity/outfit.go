package vanity

import (
	"time"

	"fmt"
)

// Outfit state constants
const (
	// Outfit states
	WEARABLE          = "WEARABLE"
	PARTIALY_WEARABLE = "PARTIALY_WEARABLE"
	NOT_WEARABLE      = "NOT_WEARABLE"
	INCOMPLETE        = "INCOMPLETE"

	// Outfit classes
	WORK          = "WORK"
	WORK_EVERYDAY = "WORK_EVERYDAY"
	EVERYDAY      = "EVERYDAY"
	GRAILED       = "GRAILED"
	POWERBLACK    = "POWERBLACK"

	// Outfit weather ratings
	HOT      = "HOT"
	STANDARD = "STANDARD"
	COLD     = "COLD"
)

// TODO:: Shiity solution for validating state changes. Find a way to implement ENUMS
var outfitStates = []string{WEARABLE, PARTIALY_WEARABLE, NOT_WEARABLE, INCOMPLETE}
var outfitClasses = []string{WORK, WORK_EVERYDAY, EVERYDAY, GRAILED, POWERBLACK}
var outfitWeather = []string{HOT, STANDARD, COLD}

// This is the struct that we commit to the datastore
type Outfit struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Class       string `json:"class"`

	// Outfit components
	Jacket      int64   `json:"jacket"`
	Tops        []int64 `json:"tops"`
	Bottom      int64   `json:"bottom"`
	Shoes       int64   `json:"shoes"`
	Accessories []int64 `json:"accessories"`

	// Outfit qualities
	Style       string `json:"style"`
	Inspiration string `json:"inspiration"`
	TotalCost   int    `json:"total_cost"`
	Pictures    string `json:"pictures"`
	Weather     string `json:"weather"`

	// Ranking data
	Rating         int `json:"rating"`
	TimesSuggested int `json:"times_suggested"`
	TimesAccepted  int `json:"times_accepted"`

	// Outfit state tracking
	MissingPieces     []int64   `json:"missing_pieces"`
	LimboPieces       []int64   `json:"limbo_pieces"`
	LastWornDate      time.Time `json:"last_worn_date"`
	ScheduledWearDate time.Time `json:"schedule_wear_date"`
	LastModified      time.Time `json:"last_modified"`
	State             string    `json:"state"`
}

func (o *Outfit) InitializeOutfit(db WardrobeDatabase) error {
	// TODO:: Lets make this a concurrent operation once it is working
	// TODO:: Implement batch read write
	// Validate Tops
	if o.Tops != nil {
		for _, pieceID := range o.Tops {
			piece, err := db.GetPiece(pieceID)
			if err != nil {
				return err
			}
			// TODO:: add onstant for this
			if piece.Type != "TOP" {
				return fmt.Errorf("Piece id: %v is not a Top", pieceID)
			}
		}
	}

	// Validate Accessories
	if o.Accessories != nil {
		for _, pieceID := range o.Accessories {
			piece, err := db.GetPiece(pieceID)
			if err != nil {
				return err
			}
			// TODO:: add onstant for this
			if piece.Type != "ACCESSORY" {
				return fmt.Errorf("Piece id: %v is not an ACCESSORY", pieceID)
			}
		}
	}

	// Validate Jacket
	if o.Jacket != 0 {
		piece, err := db.GetPiece(o.Jacket)
		if err != nil {
			return err
		}
		// TODO:: add onstant for this
		if piece.Type != "JACKET" {
			return fmt.Errorf("Piece id: %v is not a JACKET", o.Jacket)
		}
	}

	// Validate Bottom
	if o.Bottom != 0 {
		piece, err := db.GetPiece(o.Bottom)
		if err != nil {
			return err
		}
		// TODO:: add onstant for this
		if piece.Type != "BOTTOM" {
			return fmt.Errorf("Piece id: %v is not a BOTTOM", o.Bottom)
		}
	}

	// Validate Shoes
	if o.Shoes != 0 {
		piece, err := db.GetPiece(o.Shoes)
		if err != nil {
			return err
		}
		// TODO:: add onstant for this
		if piece.Type != "SHOES" {
			return fmt.Errorf("Piece id: %v is not a SHOES", o.Shoes)
		}
	}

	// Validate Class
	if o.Class != "" {
		if !validateEnum(o.Class, outfitClasses) {
			return fmt.Errorf("Invalid class: %v for outfit", o.Class)
		}
	}

	// Validate Weather
	if o.Weather != "" {
		if !validateEnum(o.Weather, outfitWeather) {
			return fmt.Errorf("Invalid weather rating: %v for outfit", o.Weather)
		}
	}

	if (o.Tops != nil) && (o.Bottom != 0) && (o.Shoes != 0) {
		o.State = WEARABLE
	} else {
		o.State = INCOMPLETE
	}
	o.LastModified = time.Now()

	return nil
}

func (o *Outfit) SetState(state string) error {
	if !validateEnum(state, outfitStates) {
		return fmt.Errorf("Invalid state: %v for outfit", state)
	}
	o.State = state
	return nil
}
