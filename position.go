package geojson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Position represents a longitude and latitude with optional elevation/altitude.
type Position struct {
	Longitude float64
	Latitude  float64
	Elevation OptionalFloat64
}

// NewPosition from longitude and latitude.
func NewPosition(long, lat float64) Position {
	return Position{
		Longitude: long,
		Latitude:  lat,
	}
}

// NewPositionWithElevation from longitude, latitude and elevation.
func NewPositionWithElevation(long, lat, elevation float64) Position {
	return Position{
		Longitude: long,
		Latitude:  lat,
		Elevation: NewOptionalFloat64(elevation),
	}
}

// MarshalJSON returns the JSON encoding of the Position.
// The JSON encoding is an array of numbers with the longitude followed by the latitude, and optional elevation.
func (p Position) MarshalJSON() ([]byte, error) {
	if p.Elevation.IsSet() {
		return json.Marshal(&position{
			p.Longitude,
			p.Latitude,
			p.Elevation.Value(),
		})
	}

	return json.Marshal(&position{
		p.Longitude,
		p.Latitude,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the results.
func (p *Position) UnmarshalJSON(data []byte) error {
	pos := position{}
	if err := json.Unmarshal(data, &pos); err != nil {
		return err
	}

	switch len(pos) {
	case 3:
		p.Elevation = NewOptionalFloat64(pos[2])
		fallthrough
	case 2:
		p.Longitude = pos[0]
		p.Latitude = pos[1]
	default:
		return fmt.Errorf("invalid position")
	}
	return nil
}

func (p Position) String() string {
	if p.Elevation.IsSet() {
		return fmt.Sprintf("[%G, %G, %G]", p.Longitude, p.Latitude, p.Elevation.Value())
	}
	return fmt.Sprintf("[%G, %G]", p.Longitude, p.Latitude)
}

// OptionalFloat64 is a type that represents a float64 that can be optionally set.
type OptionalFloat64 struct {
	value *float64
}

// NewOptionalFloat64 creates a new OptionalFloat64 set to the specified value.
func NewOptionalFloat64(val float64) OptionalFloat64 {
	return OptionalFloat64{value: &val}
}

// Value returns the value. Should call this method if OptionalFloat64.IsSet() returns true.
func (o OptionalFloat64) Value() float64 {
	return *o.value
}

// IsSet returns true if the value is set, and false if not.
func (o OptionalFloat64) IsSet() bool {
	return o.value != nil
}

// Get the float64 value and whether or not it's set.
func (o OptionalFloat64) Get() (float64, bool) {
	if o.value == nil {
		return 0, false
	}
	return *o.value, true
}

func (o OptionalFloat64) String() string {
	if o.IsSet() {
		return strconv.FormatFloat(o.Value(), 'f', -1, 64)
	}
	return "{unset}"
}

type position []float64
