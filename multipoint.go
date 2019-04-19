package geojson

import (
	"encoding/json"
)

// MultiPoint is a set of Position.
type MultiPoint []Position

// NewMultiPoint returns a MultiPoint from the specified set of position.
func NewMultiPoint(pos ...Position) *Feature {
	return &Feature{
		Geometry: (*MultiPoint)(&pos),
	}
}

// MarshalJSON returns the JSON encoding of the MultiPoint.
func (m *MultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Position(*m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPoint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]Position)(m))
}
