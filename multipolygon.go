package geojson

import (
	"encoding/json"
)

// MultiPolygon is a set of Polygons.
type MultiPolygon [][][]Position

// NewMultiPolygon returns a new MultiPolygon from the supplied polygons.
func NewMultiPolygon(p ...[][]Position) *Feature {
	return &Feature{
		Geometry: (*MultiPolygon)(&p),
	}
}

// Type returns the geometry type.
func (*MultiPolygon) Type() GeometryType {
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
	return json.Marshal([][][]Position(m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPolygon) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[][][]Position)(m))
}
