package geojson

import (
	"encoding/json"
	"fmt"
)

const (
	// FeatureCollectionType string.
	FeatureCollectionType = "FeatureCollection"
)

// FeatureCollection is a list of Features.
type FeatureCollection []Feature

// NewFeatureCollection returns a FeatureCollection consisting of the supplied Features.
func NewFeatureCollection(features ...*Feature) *FeatureCollection {
	coll := make(FeatureCollection, len(features))
	for i, f := range features {
		coll[i] = *f
	}
	return &coll
}

// MarshalJSON returns the JSON encoding of the FeatureCollection.
func (c *FeatureCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&collection{
		Type:     FeatureCollectionType,
		Features: *c,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (c *FeatureCollection) UnmarshalJSON(data []byte) error {
	col := collection{}
	if err := json.Unmarshal(data, &col); err != nil {
		return err
	}

	if col.Type != FeatureCollectionType {
		return fmt.Errorf("type is '%s', expecting '%s'", col.Type, FeatureCollectionType)
	}

	*c = col.Features
	return nil
}

type collection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
