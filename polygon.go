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
	for i, ring := range p {
		if len(ring) < 4 {
			return errLinearRingTooShort
		} else if ring[len(ring)-1] != ring[0] {
			return errLinearRingNotClosed
		}

		angle := loopToS2(ring).TurningAngle()
		if i == 0 && angle >= 0 { // CCW
			return fmt.Errorf("exterior ring must be clockwise but angle is %f", angle)
		} else if i > 0 && angle <= 0 { // CW
			return fmt.Errorf("interior ring must be counter-clockwise but angle is %f", angle)
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

var (
	errLinearRingTooShort  = fmt.Errorf("Polygon ring is too short - must contain at least 4 positions")
	errLinearRingNotClosed = fmt.Errorf("Polygon ring must be closed")
)
