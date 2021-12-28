package geojson

import (
	"encoding/json"
	"fmt"
)

// LineString is a set of two or more Positions.
type LineString []Position

// NewLineString returns a LineString from the supplied positions.
func NewLineString(pos1, pos2 Position, others ...Position) *LineString {
	all := append([]Position{pos1, pos2}, others...)
	return (*LineString)(&all)
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
	return json.Marshal(geometry{
		Type:        LineStringGeometryType,
		Coordinates: []Position(l),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (l *LineString) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates []Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*l = LineString(geo.Coordinates)
	return nil
}

// MultiLineString is a set of LineStrings.
type MultiLineString [][]Position

// NewMultiLineString returns a new MultiLineString from the supplied position "strings".
func NewMultiLineString(pos ...[]Position) *MultiLineString {
	return (*MultiLineString)(&pos)
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
	return json.Marshal(geometry{
		Type:        MultiLineStringGeometryType,
		Coordinates: [][]Position(m),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiLineString) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates [][]Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*m = MultiLineString(geo.Coordinates)
	return nil
}

var errLineStringTooShort = fmt.Errorf("LineString must contain at least 2 positions")
