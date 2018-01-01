package vanity

import (
	"fmt"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"
)

const (
	PieceKind  = "Piece"
	OutfitKind = "Oufit"
)

// Persists wardrobe data to cloud datastore
type datastoreDB struct {
	client *datastore.Client
}

// Ensure datastoreDB conforms to WardrobeDatabase
var _ WardrobeDatabase = &datastoreDB{}

func NewDatastoreDB(client *datastore.Client) (WardrobeDatabase, error) {
	ctx := context.Background()

	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoreDB: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoreDB: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}

func (db *datastoreDB) Close() {
	// No op
}

func (db *datastoreDB) datastoreKey(id int64, kind string) *datastore.Key {
	return datastore.IDKey(kind, id, nil)
}

func (db *datastoreDB) AddPiece(p *Piece) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey(PieceKind, nil)
	k, err = db.client.Put(ctx, k, p)
	if err != nil {
		return 0, fmt.Errorf("datastoreDB: could not put Piece: %v", err)
	}
	return k.ID, nil
}

func (db *datastoreDB) GetPiece(id int64) (*Piece, error) {
	ctx := context.Background()
	k := db.datastoreKey(id, PieceKind)
	piece := &Piece{}
	if err := db.client.Get(ctx, k, piece); err != nil {
		return nil, fmt.Errorf("datastoreDB: could not get Piece: %v", err)
	}
	return piece, nil
}

func (db *datastoreDB) GetPiecesByID(ids []int64) ([]*Piece, error) {
	ctx := context.Background()
	keys := make([]*datastore.Key, len(ids))
	pieces := make([]*Piece, len(ids))
	for _, id := range ids {
		keys = append(keys, db.datastoreKey(id, PieceKind))
	}

	if err := db.client.GetMulti(ctx, k, pieces); err != nil {
		return nil, fmt.Errorf("datastoreDB: could not get Piece: %v", err)
	}

	return pieces, nil
}

// func (db *datastoreDB) GetPieces() ([]*Outfit, error) {

// }

func (db *datastoreDB) UpdatePiece(p *Piece) error {
	ctx := context.Background()
	k := db.datastoreKey(p.ID, PieceKind)
	if _, err := db.client.Put(ctx, k, p); err != nil {
		return fmt.Errorf("datastoreDB: could not update Piece: %v", err)
	}
	return nil
}

func (db *datastoreDB) DeletePiece(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id, PieceKind)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoreDB: could not delete Piece: %v", err)
	}
	return nil
}

func (db *datastoreDB) AddOutfit(o *Outfit) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey(OutfitKind, nil)
	o.ID = k.ID
	k, err = db.client.Put(ctx, k, o)
	if err != nil {
		return 0, fmt.Errorf("datastoreDB: could not put Outfit: %v", err)
	}
	return k.ID, nil
}

func (db *datastoreDB) GetOutfit(id int64) (*Outfit, error) {
	ctx := context.Background()
	k := db.datastoreKey(id, OutfitKind)
	outfit := &Outfit{}
	if err := db.client.Get(ctx, k, outfit); err != nil {
		return nil, fmt.Errorf("datastoreDB: could not get Outfit: %v", err)
	}
	return outfit, nil
}

func (db *datastoreDB) GetOutfitsByID(ids []int64) ([]*Outfit, error) {
	ctx := context.Background()
	keys := make([]*datastore.Key, len(ids))
	outfits := make([]*Outfit, len(ids))
	for _, id := range ids {
		keys = append(keys, db.datastoreKey(id, OutfitKind))
	}

	if err := db.client.GetMulti(ctx, k, outfits); err != nil {
		return nil, fmt.Errorf("datastoreDB: could not get Outfit: %v", err)
	}

	return outfits, nil
}

func (db *datastoreDB) GetOutfits(state string, class string) ([]*Outfit, error) {
	ctx := context.Background()
	outfits := make([]*Outfit, 0)
	q := datastore.NewQuery(OutfitKind)
	if state != "" {
		q = q.Filter("State =", state)
	}
	if class != "" {
		q = q.Filter("Class =", class)
	}
	q = q.Order("LastWornDate")

	keys, err := db.client.GetAll(ctx, q, &outfits)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list outfits: %v", err)
	}

	for i, k := range keys {
		outfits[i].ID = k.ID
	}

	return outfits, nil
}

func (db *datastoreDB) UpdateOutfit(o *Outfit) error {
	ctx := context.Background()
	k := db.datastoreKey(o.ID, OutfitKind)
	if _, err := db.client.Put(ctx, k, o); err != nil {
		return fmt.Errorf("datastoreDB: could not update Oufit: %v", err)
	}
	return nil
}

func (db *datastoreDB) DeleteOutfit(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id, OutfitKind)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoreDB: could not delete Outfit: %v", err)
	}
	return nil
}
