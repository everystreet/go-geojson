package geojson

import (
	"encoding/json"
	"fmt"
)

// "type" properties.
const (
	TypePropFeature           = "Feature"
	TypePropFeatureCollection = "FeatureCollection"
)

// Feature consists of a specific geometry type and a list of properties.
type Feature struct {
	Geometry   Geometry
	BBox       *BoundingBox
	Properties PropertyList
}

// Geometry contains the points represented by a particular geometry type.
type Geometry interface {
	json.Marshaler
	json.Unmarshaler
	Type() GeometryType
	Validate() error
}

// MarshalJSON returns the JSON encoding of the Feature.
func (f Feature) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string       `json:"type"`
		BBox       *BoundingBox `json:"bbox,omitempty"`
		Geometry   Geometry     `json:"geometry"`
		Properties PropertyList `json:"properties,omitempty"`
	}{
		Type:       TypePropFeature,
		BBox:       f.BBox,
		Geometry:   f.Geometry,
		Properties: f.Properties,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (f *Feature) UnmarshalJSON(data []byte) error {
	var feature struct {
		Type       string          `json:"type"`
		BBox       *BoundingBox    `json:"bbox,omitempty"`
		Geometry   json.RawMessage `json:"geometry"`
		Properties PropertyList    `json:"properties,omitempty"`
	}

	if err := json.Unmarshal(data, &feature); err != nil {
		return err
	} else if feature.Type != TypePropFeature {
		return fmt.Errorf("type is '%s', expecting '%s'", feature.Type, TypePropFeature)
	}

	f.BBox = feature.BBox
	f.Properties = feature.Properties

	geo, err := unmarshalGeometry(feature.Geometry)
	if err != nil {
		return err
	}
	f.Geometry = geo
	return nil
}

// WithBoundingBox sets the optional bounding box.
func (f *Feature) WithBoundingBox(bottomLeft, topRight Position) *Feature {
	f.BBox = &BoundingBox{
		BottomLeft: bottomLeft,
		TopRight:   topRight,
	}
	return f
}

// WithProperties sets the optional properties, removing all existing properties.
func (f *Feature) WithProperties(props ...Property) *Feature {
	f.Properties = PropertyList(props)
	return f
}

// AddProperty appends a new property.
func (f *Feature) AddProperty(name string, value interface{}) *Feature {
	f.Properties = append(f.Properties, Property{
		Name:  name,
		Value: value,
	})
	return f
}

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
func (c FeatureCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&featureCollection{
		Type:     TypePropFeatureCollection,
		BBox:     c.BBox,
		Features: c.Features,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (c *FeatureCollection) UnmarshalJSON(data []byte) error {
	col := featureCollection{}
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

type featureCollection struct {
	Type     string       `json:"type"`
	BBox     *BoundingBox `json:"bbox,omitempty"`
	Features []Feature    `json:"features"`
}
