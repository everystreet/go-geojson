package geojson

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func loopToS2(loop []Position) *s2.Loop {
	points := make([]s2.Point, len(loop))
	for i, pos := range loop {
		points[i] = s2.PointFromLatLng(
			s2.LatLngFromDegrees(pos.Latitude, pos.Longitude))
	}
	return s2.LoopFromPoints(points)
}
