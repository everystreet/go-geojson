package geojson

// Point is a single set of Position.
type Point Position

// NewPoint returns a Point Feature with the specified longitude and latitude.
func NewPoint(long, lat float64) *Feature {
	return &Feature{
		Geometry: &Point{
			Longitude: long,
			Latitude:  lat,
		},
	}
}

// NewPointWithElevation returns a Point Feature with the specified longitude, latitude and elevation.
func NewPointWithElevation(long, lat, elevation float64) *Feature {
	return &Feature{
		Geometry: &Point{
			Longitude: long,
			Latitude:  lat,
			Elevation: NewOptionalFloat64(elevation),
		},
	}
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
	return (*Position)(&p).MarshalJSON()
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Point) UnmarshalJSON(data []byte) error {
	return (*Position)(p).UnmarshalJSON(data)
}
