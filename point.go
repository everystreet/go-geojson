package geojson

// Point is a single set of Coordinates.
type Point Coordinates

// MarshalJSON returns the JSON encoding of the Point.
func (p *Point) MarshalJSON() ([]byte, error) {
	return (*Coordinates)(p).MarshalJSON()
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Point) UnmarshalJSON(data []byte) error {
	return (*Coordinates)(p).UnmarshalJSON(data)
}

// NewPoint returns a Point Feature with the specified longitude and latitude.
func NewPoint(long, lat float64, props ...Property) *Feature {
	return &Feature{
		Geometry: Point{
			Longitude: long,
			Latitude:  lat,
		},
		Properties: PropertyList(props),
	}
}

// NewPointWithElevation returns a Point Feature with the specified longitude, latitude and elevation.
func NewPointWithElevation(long, lat, elevation float64, props ...Property) *Feature {
	return &Feature{
		Geometry: Point{
			Longitude: long,
			Latitude:  lat,
			Elevation: NewOptionalFloat64(elevation),
		},
		Properties: PropertyList(props),
	}
}
