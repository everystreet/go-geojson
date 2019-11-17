package geojson

import (
	"encoding/json"
)

// MultiLineString is a set of LineStrings.
type MultiLineString [][]Position

// NewMultiLineString returns a new MultiLineString from the supplied position "strings".
func NewMultiLineString(ls ...[]Position) *Feature {
	return &Feature{
		Geometry: (*MultiLineString)(&ls),
	}
}

// Type returns the geometry type.
func (m MultiLineString) Type() GeometryType {
	return MultiLineStringGeometryType
}

// Validate the MultiLineString.
func (m MultiLineString) Validate() error {
	for _, ls := range m {
		if len(ls) < 2 {
			return errLineStringTooShort
		}
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the MultiLineString.
func (m MultiLineString) MarshalJSON() ([]byte, error) {
	return json.Marshal([][]Position(m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiLineString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[][]Position)(m))
}
