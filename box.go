package geojson

import (
	"encoding/json"
	"fmt"
)

// BoundingBox represents a bounding box in either 2D or 3D space.
type BoundingBox struct {
	BottomLeft Position
	TopRight   Position
}

// MarshalJSON returns the JSON encoding of the box.
// It is an error if only 1 elevation value is set - either both positions or neither must have it.
func (b BoundingBox) MarshalJSON() ([]byte, error) {
	if b.BottomLeft.Elevation.IsSet() || b.TopRight.Elevation.IsSet() {
		if !b.BottomLeft.Elevation.IsSet() || !b.TopRight.Elevation.IsSet() {
			return nil, fmt.Errorf("bounding box positions must be in the same dimension")
		}

		return json.Marshal(&position{
			b.BottomLeft.LatLng.Lng.Degrees(),
			b.BottomLeft.LatLng.Lat.Degrees(),
			b.BottomLeft.Elevation.Value(),
			b.TopRight.LatLng.Lng.Degrees(),
			b.TopRight.LatLng.Lat.Degrees(),
			b.TopRight.Elevation.Value(),
		})
	}

	return json.Marshal(&position{
		b.BottomLeft.LatLng.Lng.Degrees(),
		b.BottomLeft.LatLng.Lat.Degrees(),
		b.TopRight.LatLng.Lng.Degrees(),
		b.TopRight.LatLng.Lat.Degrees(),
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
		b.BottomLeft = MakePosition(pos[1], pos[0])
		b.TopRight = MakePosition(pos[3], pos[2])
	case 6:
		b.BottomLeft = MakePositionWithElevation(pos[1], pos[0], pos[2])
		b.TopRight = MakePositionWithElevation(pos[4], pos[3], pos[5])
	default:
		return fmt.Errorf("invalid position")
	}
	return nil
}

// Validate the bounding box.
func (b BoundingBox) Validate() error {
	if (b.BottomLeft.Elevation.IsSet() && b.TopRight.Elevation.IsSet()) ||
		(!b.BottomLeft.Elevation.IsSet() && !b.TopRight.Elevation.IsSet()) {
		return fmt.Errorf("bounding box positions must be in the same dimension")
	} else if !b.BottomLeft.LatLng.IsValid() {
		return fmt.Errorf("bottom left is invalid")
	} else if !b.TopRight.LatLng.IsValid() {
		return fmt.Errorf("top right is invalid")
	}
	return nil
}
