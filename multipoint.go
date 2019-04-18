package geojson

import (
	"encoding/json"
	"fmt"
)

// MultiPoint is a set of Coordinates.
type MultiPoint []Coordinates

// NewMultiPoint returns a MultiPoint from the specified set of coordinates.
func NewMultiPoint(coords ...Coordinates) *Feature {
	return &Feature{
		Geometry: (*MultiPoint)(&coords),
	}
}

// MarshalJSON returns the JSON encoding of the MultiPoint.
func (m *MultiPoint) MarshalJSON() ([]byte, error) {
	fmt.Println("marshal")
	return json.Marshal([]Coordinates(*m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPoint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]Coordinates)(m))
}
