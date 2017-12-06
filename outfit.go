package vanity

import (
	"time"

	"cloud.google.com/go/datastore"
	"encoding/json"
	"fmt"
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

// JSON encode array of keys and save to outfit Tops field as string
func (o *Outfit) EncodeTopsKeys(keys []*datastore.Key) error {
	json_keys, err := json.Marshal(keys)
	if err != nil {
		return fmt.Errorf("json: could not marshal keys: %v", err)
	}
	o.Tops = string(json_keys)
	return nil
}

// JSON decode array of keys from outfit Tops field and return as slice of Keys
func (o *Outfit) DecodeTopsKeys() ([]datastore.Key, error) {
	var keys []datastore.Key
	err := json.Unmarshal([]byte(o.Tops), &keys)
	if err != nil {
		return nil, fmt.Errorf("json: could not unmarshal keys: %v", err)
	}
	return keys, nil
}

// JSON encode array of keys and save to outfit Accessories field as string
func (o *Outfit) EncodeAccsKeys(keys []*datastore.Key) error {
	// JSON encode array of keys and save to outfit
	json_keys, err := json.Marshal(keys)
	if err != nil {
		return fmt.Errorf("json: could not marshal keys: %v", err)
	}
	o.Accessories = string(json_keys)
	return nil
}

// JSON decode array of keys from outfit Accessories field and return as slice of Keys
func (o *Outfit) DecodeAccsKeys() ([]datastore.Key, error) {
	// JSON decode array of keys
	var keys []datastore.Key
	err := json.Unmarshal([]byte(o.Accessories), &keys)
	if err != nil {
		return nil, fmt.Errorf("json: could not unmarshal keys: %v", err)
	}
	return keys, nil
}
