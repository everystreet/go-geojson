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

		if angle := LoopToS2(ring).TurningAngle(); i == 0 && angle >= 0 { // CCW
			return fmt.Errorf("exterior ring must be clockwise but angle is %f", angle)
		} else if i > 0 && angle <= 0 { // CW
			return fmt.Errorf("interior ring must be counter-clockwise but angle is %f", angle)
		}
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the Polygon.
func (p Polygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(geometry{
		Type:        PolygonGeometryType,
		Coordinates: [][]Position(p),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Polygon) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates [][]Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*p = Polygon(geo.Coordinates)
	return nil
}

// MultiPolygon is a set of Polygons.
type MultiPolygon [][][]Position

// NewMultiPolygon returns a new MultiPolygon from the supplied polygons.
func NewMultiPolygon(p ...[][]Position) *Feature {
	return &Feature{
		Geometry: (*MultiPolygon)(&p),
	}
}

// Type returns the geometry type.
func (m MultiPolygon) Type() GeometryType {
	return MultiPolygonGeometryType
}

// Validate the MultiPolygon.
func (m MultiPolygon) Validate() error {
	for _, polygon := range m {
		if err := Polygon(polygon).Validate(); err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the MultiPolygon.
func (m MultiPolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(geometry{
		Type:        MultiPolygonGeometryType,
		Coordinates: [][][]Position(m),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPolygon) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates [][][]Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*m = MultiPolygon(geo.Coordinates)
	return nil
}

var (
	errLinearRingTooShort  = fmt.Errorf("Polygon ring is too short - must contain at least 4 positions")
	errLinearRingNotClosed = fmt.Errorf("Polygon ring must be closed")
)
