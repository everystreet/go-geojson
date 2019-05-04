package geojson

import (
	"encoding/json"
	"errors"
)

// Polygon is a set of linear rings (closed LineStrings).
type Polygon [][]Position

// NewPolygon returns a new Polygon from the supplied linear rings.
func NewPolygon(rings ...[]Position) *Feature {
	return &Feature{
		Geometry: (*Polygon)(&rings),
	}
}

// MarshalJSON returns the JSON encoding of the Polygon.
func (p *Polygon) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	return json.Marshal([][]Position(*p))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Polygon) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[][]Position)(p)); err != nil {
		return err
	}
	return p.validate()
}

func (p *Polygon) validate() error {
	for _, r := range *p {
		if len(r) < 4 {
			return errLinearRingTooShort
		}
	}
	return nil
}

var errLinearRingTooShort = errors.New("Polygon ring must contain at least 4 positions")
