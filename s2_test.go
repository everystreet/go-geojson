package geojson_test

import (
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestLineStringToS2(t *testing.T) {
	linestring := geojson.NewLineString(
		geojson.MakePosition(7, 4),
		geojson.MakePosition(6, 9),
		geojson.MakePosition(2, 6),
	)

	polyline, err := geojson.LineStringToS2(*linestring)
	require.NoError(t, err)
	require.Equal(t, 2, polyline.NumEdges())
}

func TestPolygonToS2(t *testing.T) {
	polygon := geojson.NewPolygon(
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
	)

	s2Polygon, err := geojson.PolygonToS2(*polygon)
	require.NoError(t, err)
	require.Equal(t, 2, s2Polygon.NumLoops())
	require.Equal(t, 9, s2Polygon.NumEdges())
}
