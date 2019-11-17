package geojson

import (
	"encoding/json"
	"fmt"
)

// Polygon is a set of linear rings (closed LineStrings).
type Polygon [][]Position

// NewPolygon returns a new Polygon from the supplied linear rings.
func NewPolygon(rings ...[]Position) *Feature {
	return &Feature{
		Geometry: (*Polygon)(&rings),
	}
}

// Type returns the geometry type.
func (p Polygon) Type() GeometryType {
	return PolygonGeometryType
}

// Validate the Polygon.
func (p Polygon) Validate() error {
	for _, r := range p {
		if len(r) < 4 {
			return errLinearRingTooShort
		}
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the Polygon.
func (p *Polygon) MarshalJSON() ([]byte, error) {
	return json.Marshal([][]Position(*p))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Polygon) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[][]Position)(p))
}

var errLinearRingTooShort = fmt.Errorf("Polygon ring must contain at least 4 positions")
