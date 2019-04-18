package geojson

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BoundingBox represents a bounding box in either 2D or 3D space.
type BoundingBox struct {
	BottomLeft Coordinates
	TopRight   Coordinates
}

// MarshalJSON returns the JSON encoding of the box.
// It is an error if only 1 elevation value is set - either both coordinates or neither must have it.
func (b *BoundingBox) MarshalJSON() ([]byte, error) {
	if b.BottomLeft.Elevation.IsSet() || b.TopRight.Elevation.IsSet() {
		if !b.BottomLeft.Elevation.IsSet() || !b.TopRight.Elevation.IsSet() {
			return nil, fmt.Errorf("bounding box coordinates must be the same setting")
		}

		return json.Marshal(&coordinates{
			b.BottomLeft.Longitude,
			b.BottomLeft.Latitude,
			b.BottomLeft.Elevation.Value(),
			b.TopRight.Longitude,
			b.TopRight.Latitude,
			b.TopRight.Elevation.Value(),
		})
	}

	return json.Marshal(&coordinates{
		b.BottomLeft.Longitude,
		b.BottomLeft.Latitude,
		b.TopRight.Longitude,
		b.TopRight.Latitude,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the results.
func (b *BoundingBox) UnmarshalJSON(data []byte) error {
	coords := coordinates{}
	if err := json.Unmarshal(data, &coords); err != nil {
		return err
	}

	switch len(coords) {
	case 4:
		b.BottomLeft = NewCoordinates(coords[0], coords[1])
		b.TopRight = NewCoordinates(coords[2], coords[3])
	case 6:
		b.BottomLeft = NewCoordinatesWithElevation(coords[0], coords[1], coords[2])
		b.TopRight = NewCoordinatesWithElevation(coords[3], coords[4], coords[5])
	default:
		return errors.New("invalid coordinates")
	}
	return nil
}
