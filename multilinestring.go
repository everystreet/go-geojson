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
func (*MultiLineString) Type() GeometryType {
	return MultiLineStringGeometryType
}

// MarshalJSON returns the JSON encoding of the MultiLineString.
func (m *MultiLineString) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	return json.Marshal([][]Position(*m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiLineString) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[][]Position)(m)); err != nil {
		return err
	}
	return m.validate()
}

func (m *MultiLineString) validate() error {
	for _, ls := range *m {
		if len(ls) < 2 {
			return errLineStringTooShort
		}
	}
	return nil
}
