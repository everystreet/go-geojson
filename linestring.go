package geojson

import (
	"encoding/json"
	"fmt"
)

// LineString is a set of two or more Positions.
type LineString []Position

// NewLineString returns a LineString from the supplied positions.
func NewLineString(pos1, pos2 Position, others ...Position) *Feature {
	all := append([]Position{pos1, pos2}, others...)
	return &Feature{
		Geometry: (*LineString)(&all),
	}
}

// Type returns the geometry type.
func (l LineString) Type() GeometryType {
	return LineStringGeometryType
}

// Validate the LineString.
func (l LineString) Validate() error {
	if len(l) < 2 {
		return errLineStringTooShort
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the LineString.
func (l LineString) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Position(l))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (l *LineString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]Position)(l))
}

var errLineStringTooShort = fmt.Errorf("LineString must contain at least 2 positions")
