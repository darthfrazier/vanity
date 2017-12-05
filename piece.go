package vanity

import "time"

type Piece struct {
	ID          int64
	Name        string
	Brand       string
	Description string
	// TODO:: I want these to be an ENUM. Figure out how to do that
	PieceType    string
	PieceSubType string

	// Piece qualities
	Price   int
	Color   string
	Pattern string
	Picture string

	// Piece state tracking
	TimesWornSinceLastWash int
	LastWornDate           time.Time
	// TODO:: I want these to be an ENUM. Figure out how to do that
	LaundryState string
	WearState    string
	State        string
}
