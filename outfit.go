package vanity

import (
	"time"

	"cloud.google.com/go/datastore"
	"encoding/json"
	"fmt"
)

type Outfit struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Class       string `json:"class"`

	// Outfit components
	Jacket      *datastore.Key `json:"jacket"`
	Tops        string         `json:"tops"`
	Bottom      *datastore.Key `json:"bottom"`
	Shoes       *datastore.Key `json:"shoes"`
	Accessories string         `json:"accessories"`

	// Outfit qualities
	Style       string `json:"style"`
	Inspiration string `json:"inspiration"`
	TotalCost   int    `json:"total_cost"`
	Rating      int    `json:"rating"`
	Pictures    string `json:"pictures"`

	// Outfit state tracking
	LastWornDate      time.Time `json:"last_worn_date"`
	ScheduledWearDate time.Time `json:"schedule_wear_date"`
	LastModified      time.Time `json:"last_modified"`
	State             string    `json:"state"`
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
