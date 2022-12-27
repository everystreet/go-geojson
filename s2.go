package geojson

import (
	"fmt"

	"github.com/golang/geo/s2"
)

// LineStringToS2 returns an S2 Geometry polyline.
func LineStringToS2(linestring LineString) (*s2.Polyline, error) {
	latlngs := make([]s2.LatLng, len(linestring))
	for i, pos := range linestring {
		latlngs[i] = pos.pos
	}
	polyline := s2.PolylineFromLatLngs(latlngs)
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
		return s2.LoopFromPoints([]s2.Point{s2.PointFromLatLng(loop[0].pos)})
	}

	points := make([]s2.Point, len(loop)-1)
	for i := 0; i < len(loop)-1; i++ {
		points[i] = s2.PointFromLatLng(loop[i].pos)
	}
	return s2.LoopFromPoints(points)
}
