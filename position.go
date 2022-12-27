package geojson

import (
	"encoding/json"
	"fmt"

	"github.com/golang/geo/s2"
)

// Position represents a longitude and latitude with optional elevation/altitude.
type Position struct {
	pos       s2.LatLng
	elevation *float64
}

// MakePosition from longitude and latitude.
func MakePosition(lat, lng float64) Position {
	return Position{
		pos: s2.LatLngFromDegrees(lat, lng),
	}
}

// MakePositionWithElevation from longitude, latitude and elevation.
func MakePositionWithElevation(lat, lng, elevation float64) Position {
	return Position{
		pos:       s2.LatLngFromDegrees(lat, lng),
		elevation: &elevation,
	}
}

func (p Position) String() string {
	if p.elevation != nil {
		return fmt.Sprintf("[%G, %G, %G]", p.pos.Lng.Degrees(), p.pos.Lat.Degrees(), *p.elevation)
	}
	return fmt.Sprintf("[%G, %G]", p.pos.Lng.Degrees(), p.pos.Lat.Degrees())
}

// Validate the position.
func (p Position) Validate() error {
	if !p.pos.IsValid() {
		return fmt.Errorf("invalid latlng")
	}
	return nil
}

// MarshalJSON returns the JSON encoding of the Position.
// The JSON encoding is an array of numbers with the longitude followed by the latitude, and optional elevation.
func (p Position) MarshalJSON() ([]byte, error) {
	if p.elevation != nil {
		return json.Marshal(&position{
			p.pos.Lng.Degrees(),
			p.pos.Lat.Degrees(),
			*p.elevation,
		})
	}

	return json.Marshal(&position{
		p.pos.Lng.Degrees(),
		p.pos.Lat.Degrees(),
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
		p.elevation = &pos[2]
		fallthrough
	case 2:
		p.pos = s2.LatLngFromDegrees(pos[1], pos[0])
	default:
		return fmt.Errorf("invalid position")
	}
	return nil
}

type position []float64
