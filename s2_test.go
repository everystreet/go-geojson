package geojson_test

import (
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestLineStringToS2(t *testing.T) {
	s := geojson.NewLineString(
		geojson.MakePosition(7, 4),
		geojson.MakePosition(6, 9),
		geojson.MakePosition(2, 6),
	).Geometry.(*geojson.LineString)

	linestring, err := geojson.LineStringToS2(*s)
	require.NoError(t, err)
	require.Equal(t, 2, linestring.NumEdges())
}

func TestPolygonToS2(t *testing.T) {
	p := geojson.NewPolygon(
		[]geojson.Position{
			geojson.MakePosition(7, 7),
			geojson.MakePosition(4, 8),
			geojson.MakePosition(3, 4),
			geojson.MakePosition(5, 2),
			geojson.MakePosition(7, 3),
			geojson.MakePosition(7, 7),
		},
		[]geojson.Position{
			geojson.MakePosition(4, 4),
			geojson.MakePosition(4, 6),
			geojson.MakePosition(5, 7),
			geojson.MakePosition(6, 4),
			geojson.MakePosition(4, 4),
		},
	).Geometry.(*geojson.Polygon)

	polygon, err := geojson.PolygonToS2(*p)
	require.NoError(t, err)
	require.Equal(t, 2, polygon.NumLoops())
	require.Equal(t, 9, polygon.NumEdges())
}
