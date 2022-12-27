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

// NewFeature creates a new feature.
func NewFeature[G Geometry](geometry G, properties ...Property) Feature[G] {
	return Feature[G]{
		geometry:   geometry,
		properties: PropertyList(properties),
	}
}

// NewFeatureWithBoundingBox creates a new feature with the supplied bounding box.
func NewFeatureWithBoundingBox[G Geometry](geometry G, box BoundingBox, properties ...Property) Feature[G] {
	return Feature[G]{
		geometry:   geometry,
		box:        &box,
		properties: PropertyList(properties),
	}
}

// Feature consists of a specific geometry type and a list of properties.
type Feature[G Geometry] struct {
	geometry   G
	box        *BoundingBox
	properties PropertyList
}

// Geometry contains the points represented by a particular geometry type.
type Geometry interface {
	json.Marshaler
	json.Unmarshaler
	Type() GeometryType
	Validate() error
}

// Validate the feature geometry.
func (f Feature[G]) Validate() error {
	return f.geometry.Validate()
}

// Geometry returns the stored geometry.
func (f Feature[G]) Geometry() Geometry {
	return f.geometry
}

// BoundingBox returns the stored bounding box.
func (f Feature[G]) BoundingBox() *BoundingBox {
	return f.box
}

// Properties returns the stored properties.
func (f Feature[G]) Properties() PropertyList {
	return f.properties
}

// MarshalJSON returns the JSON encoding of the Feature.
func (f Feature[G]) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string       `json:"type"`
		Box        *BoundingBox `json:"bbox,omitempty"`
		Geometry   Geometry     `json:"geometry"`
		Properties PropertyList `json:"properties,omitempty"`
	}{
		Type:       TypePropFeature,
		Box:        f.box,
		Geometry:   f.geometry,
		Properties: f.properties,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (f *Feature[G]) UnmarshalJSON(data []byte) error {
	var feature struct {
		Type       string          `json:"type"`
		Box        *BoundingBox    `json:"bbox,omitempty"`
		Geometry   json.RawMessage `json:"geometry"`
		Properties PropertyList    `json:"properties,omitempty"`
	}

	if err := json.Unmarshal(data, &feature); err != nil {
		return err
	} else if feature.Type != TypePropFeature {
		return fmt.Errorf("type is '%s', expecting '%s'", feature.Type, TypePropFeature)
	}

	f.box = feature.Box
	f.properties = feature.Properties

	geo, err := unmarshalGeometry(feature.Geometry)
	if err != nil {
		return err
	}
	f.geometry = geo.(G)
	return nil
}

// NewFeatureCollection creates a new feature collection.
func NewFeatureCollection(features ...Feature[Geometry]) FeatureCollection {
	return FeatureCollection{
		features: features,
	}
}

// NewFeatureCollectionWithBoundingBox creates a new feature collection with the supplied bounding box.
func NewFeatureCollectionWithBoundingBox(box BoundingBox, features ...Feature[Geometry]) FeatureCollection {
	return FeatureCollection{
		features: features,
		box:      &box,
	}
}

// FeatureCollection is a list of Features.
type FeatureCollection struct {
	features []Feature[Geometry]
	box      *BoundingBox
}

// MarshalJSON returns the JSON encoding of the FeatureCollection.
func (c FeatureCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&featureCollection{
		Type:     TypePropFeatureCollection,
		Box:      c.box,
		Features: c.features,
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

	c.box = col.Box
	c.features = col.Features
	return nil
}

type featureCollection struct {
	Type     string              `json:"type"`
	Box      *BoundingBox        `json:"bbox,omitempty"`
	Features []Feature[Geometry] `json:"features"`
}
