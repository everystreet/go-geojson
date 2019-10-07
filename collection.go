package geojson

import (
	"encoding/json"
	"fmt"
)

// TypePropFeatureCollection is the value of the "type" property for Feature Collections.
const TypePropFeatureCollection = "FeatureCollection"

// FeatureCollection is a list of Features.
type FeatureCollection struct {
	Features []Feature
	BBox     *BoundingBox
}

// NewFeatureCollection returns a FeatureCollection consisting of the supplied Features.
func NewFeatureCollection(features ...*Feature) *FeatureCollection {
	c := FeatureCollection{
		Features: make([]Feature, len(features)),
	}

	for i, f := range features {
		c.Features[i] = *f
	}
	return &c
}

// MarshalJSON returns the JSON encoding of the FeatureCollection.
func (c *FeatureCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&collection{
		Type:     TypePropFeatureCollection,
		BBox:     c.BBox,
		Features: c.Features,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (c *FeatureCollection) UnmarshalJSON(data []byte) error {
	col := collection{}
	if err := json.Unmarshal(data, &col); err != nil {
		return err
	}

	if col.Type != TypePropFeatureCollection {
		return fmt.Errorf("type is '%s', expecting '%s'", col.Type, TypePropFeatureCollection)
	}

	c.BBox = col.BBox
	c.Features = col.Features
	return nil
}

// WithBoundingBox sets the optional bounding box.
func (c *FeatureCollection) WithBoundingBox(bottomLeft, topRight Position) *FeatureCollection {
	c.BBox = &BoundingBox{
		BottomLeft: bottomLeft,
		TopRight:   topRight,
	}
	return c
}

type collection struct {
	Type     string       `json:"type"`
	BBox     *BoundingBox `json:"bbox,omitempty"`
	Features []Feature    `json:"features"`
}
