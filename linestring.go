package geojson

import (
	"encoding/json"
	"errors"
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

// MarshalJSON returns the JSON encoding of the LineString.
func (l *LineString) MarshalJSON() ([]byte, error) {
	if err := l.validate(); err != nil {
		return nil, err
	}
	return json.Marshal([]Position(*l))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (l *LineString) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[]Position)(l)); err != nil {
		return err
	}
	return l.validate()
}

func (l *LineString) validate() error {
	if len(*l) < 2 {
		return errLineStringTooShort
	}
	return nil
}

var errLineStringTooShort = errors.New("LineString must contain at least 2 positions")
