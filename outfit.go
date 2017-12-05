package vanity

import (
	"time"

	"cloud.google.com/go/datastore"
)

type Outfit struct {
	ID          int64
	name        string
	description string
	class       string

	// Outfit components
	// TODO:: properly define enconding for these fields
	Jacket      *datastore.Key
	Tops        string
	Bottom      *datastore.Key
	Shoes       *datastore.Key
	Accessories string

	// Outfit qualities
	Style       string
	Inspiration string
	TotalCost   int
	Rating      int
	Pictures    string

	// Outfit state tracking
	LastWornDate      time.Time
	ScheduledWearDate time.Time
	LastModified      time.Time
	State             string
}
