package geojson

import (
	"encoding/json"
)

// Point is a single set of Position.
type Point Position

// NewPoint returns a point with the specified longitude and latitude.
func NewPoint(lat, lng float64) *Point {
	pos := MakePosition(lat, lng)
	return (*Point)(&pos)
}

// NewPointWithElevation returns a Point Feature with the specified longitude, latitude and elevation.
func NewPointWithElevation(lat, lng, elevation float64) *Point {
	pos := MakePositionWithElevation(lat, lng, elevation)
	return (*Point)(&pos)
}

// Type returns the geometry type.
func (p Point) Type() GeometryType {
	return PointGeometryType
}

// Validate the Point.
func (p Point) Validate() error {
	return nil
}

// MarshalJSON returns the JSON encoding of the Point.
func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(geometry{
		Type:        PointGeometryType,
		Coordinates: Position(p),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Point) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*p = Point(geo.Coordinates)
	return nil
}

// MultiPoint is a set of Position.
type MultiPoint []Position

// NewMultiPoint returns a multipoint with the specified set of positions.
func NewMultiPoint(pos ...Position) *MultiPoint {
	return (*MultiPoint)(&pos)
}

// Type returns the geometry type.
func (m MultiPoint) Type() GeometryType {
	return MultiPointGeometryType
}

// Validate the MultiPoint.
func (m MultiPoint) Validate() error {
	return nil
}

// MarshalJSON returns the JSON encoding of the MultiPoint.
func (m MultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(geometry{
		Type:        MultiPointGeometryType,
		Coordinates: []Position(m),
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPoint) UnmarshalJSON(data []byte) error {
	var geo struct {
		Coordinates []Position `json:"coordinates"`
	}

	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	*m = MultiPoint(geo.Coordinates)
	return nil
}
