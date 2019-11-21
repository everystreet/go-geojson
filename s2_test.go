package geojson_test

import (
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestLineStringToS2(t *testing.T) {
	s := geojson.NewLineString(
		geojson.NewPosition(7, 4),
		geojson.NewPosition(6, 9),
		geojson.NewPosition(2, 6),
	).Geometry.(*geojson.LineString)

	linestring, err := geojson.LineStringToS2(*s)
	require.NoError(t, err)
	require.Equal(t, 2, linestring.NumEdges())
}

func TestPolygonToS2(t *testing.T) {
	p := geojson.NewPolygon(
		[]geojson.Position{
			geojson.NewPosition(7, 7),
			geojson.NewPosition(4, 8),
			geojson.NewPosition(3, 4),
			geojson.NewPosition(5, 2),
			geojson.NewPosition(7, 3),
			geojson.NewPosition(7, 7),
		},
		[]geojson.Position{
			geojson.NewPosition(4, 4),
			geojson.NewPosition(4, 6),
			geojson.NewPosition(5, 7),
			geojson.NewPosition(6, 4),
			geojson.NewPosition(4, 4),
		},
	).Geometry.(*geojson.Polygon)

	polygon, err := geojson.PolygonToS2(*p)
	require.NoError(t, err)
	require.Equal(t, 2, polygon.NumLoops())
	require.Equal(t, 9, polygon.NumEdges())
}
