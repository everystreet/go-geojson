package geojson

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BoundingBox represents a bounding box in either 2D or 3D space.
type BoundingBox struct {
	BottomLeft Position
	TopRight   Position
}

// MarshalJSON returns the JSON encoding of the box.
// It is an error if only 1 elevation value is set - either both positions or neither must have it.
func (b *BoundingBox) MarshalJSON() ([]byte, error) {
	if b.BottomLeft.Elevation.IsSet() || b.TopRight.Elevation.IsSet() {
		if !b.BottomLeft.Elevation.IsSet() || !b.TopRight.Elevation.IsSet() {
			return nil, fmt.Errorf("bounding box positions must be in the same setting")
		}

		return json.Marshal(&position{
			b.BottomLeft.Longitude,
			b.BottomLeft.Latitude,
			b.BottomLeft.Elevation.Value(),
			b.TopRight.Longitude,
			b.TopRight.Latitude,
			b.TopRight.Elevation.Value(),
		})
	}

	return json.Marshal(&position{
		b.BottomLeft.Longitude,
		b.BottomLeft.Latitude,
		b.TopRight.Longitude,
		b.TopRight.Latitude,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the results.
func (b *BoundingBox) UnmarshalJSON(data []byte) error {
	pos := position{}
	if err := json.Unmarshal(data, &pos); err != nil {
		return err
	}

	switch len(pos) {
	case 4:
		b.BottomLeft = NewPosition(pos[0], pos[1])
		b.TopRight = NewPosition(pos[2], pos[3])
	case 6:
		b.BottomLeft = NewPositionWithElevation(pos[0], pos[1], pos[2])
		b.TopRight = NewPositionWithElevation(pos[3], pos[4], pos[5])
	default:
		return errors.New("invalid position")
	}
	return nil
}
