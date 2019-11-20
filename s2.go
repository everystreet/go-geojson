package geojson

import (
	"fmt"

	"github.com/golang/geo/s2"
)

// LineStringToS2 returns an S2 Geometry polyline.
func LineStringToS2(linestring LineString) (*s2.Polyline, error) {
	points := make([]s2.LatLng, len(linestring))
	for i, pos := range linestring {
		points[i] = s2.LatLngFromDegrees(pos.Latitude, pos.Longitude)
	}
	polyline := s2.PolylineFromLatLngs(points)
	return polyline, polyline.Validate()
}

// PolygonToS2 returns an S2 Geometry polygon.
func PolygonToS2(polygon Polygon) (*s2.Polygon, error) {
	if err := polygon.Validate(); err != nil {
		return nil, err
	}

	loops := make([]*s2.Loop, len(polygon))
	for i, loop := range polygon {
		loops[i] = LoopToS2(loop)
		if err := loops[i].Validate(); err != nil {
			return nil, fmt.Errorf("invalid loop '%d': %w", i, err)
		}
	}
	p := s2.PolygonFromLoops(loops)
	return p, p.Validate()
}

// LoopToS2 returns an S2 Loop.
func LoopToS2(loop []Position) *s2.Loop {
	if n := len(loop); n == 0 {
		return s2.EmptyLoop()
	} else if n == 1 {
		return s2.LoopFromPoints([]s2.Point{
			s2.PointFromLatLng(s2.LatLngFromDegrees(loop[0].Latitude, loop[0].Longitude)),
		})
	}

	points := make([]s2.Point, len(loop)-1)
	for i := 0; i < len(loop)-1; i++ {
		points[i] = s2.PointFromLatLng(s2.LatLngFromDegrees(loop[i].Latitude, loop[i].Longitude))
	}
	return s2.LoopFromPoints(points)
}
