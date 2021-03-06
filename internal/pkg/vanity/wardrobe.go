package vanity

type WardrobeDatabase interface {
	Close()

	// Outfit CRUD
	AddPiece(p *Piece) (id int64, err error)

	GetPiece(id int64) (*Piece, error)

	// TODO:: Define some sort of struct that I can build query from
	// GetPieces() ([]*Piece, error)

	UpdatePiece(p *Piece) error

	DeletePiece(id int64) error

	// Outfit CRUD
	AddOutfit(o *Outfit) (id int64, err error)

	GetOutfit(id int64) (*Outfit, error)

	GetOutfitsByID(ids []int64) ([]*Outfit, error)

	// TODO:: Define some sort of struct that I can build query from
	GetOutfits(state string, class string) ([]*Outfit, error)

	UpdateOutfit(o *Outfit) error

	UpdateOutfits(outfits []*Outfit) error

	DeleteOutfit(id int64) error
}
