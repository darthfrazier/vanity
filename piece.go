package vanity

import (
	"time"
)

type Piece struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	// TODO:: I want these to be an ENUM. Figure out how to do that
	PieceType    string `json:"piece_type"`
	PieceSubType string `json:"piece_subtype"`

	// Piece qualities
	Price   int    `json:"price"`
	Color   string `json:"color"`
	Pattern string `json:"pattern"`
	Picture string `json:"picture"`

	// Piece state tracking
	TimesWornSinceLastWash int       `json:"times_worn_since_last_wash"`
	LastWornDate           time.Time `json:"last_worn_date"`
	// TODO:: I want these to be an ENUM. Figure out how to do that
	LaundryState string `json:"laundry_state"`
	WearState    string `json:"wear_state"`
	State        string `json:"state"`
}
