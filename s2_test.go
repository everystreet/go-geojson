package geojson_test

import (
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestLineStringToS2(t *testing.T) {
	feature := geojson.Feature[*geojson.LineString]{
		Geometry: geojson.NewLineString(
			geojson.MakePosition(7, 4),
			geojson.MakePosition(6, 9),
			geojson.MakePosition(2, 6),
		),
	}

	linestring, err := geojson.LineStringToS2(*feature.Geometry)
	require.NoError(t, err)
	require.Equal(t, 2, linestring.NumEdges())
}

func TestPolygonToS2(t *testing.T) {
	feature := geojson.Feature[*geojson.Polygon]{
		Geometry: geojson.NewPolygon(
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
		),
	}

	polygon, err := geojson.PolygonToS2(*feature.Geometry)
	require.NoError(t, err)
	require.Equal(t, 2, polygon.NumLoops())
	require.Equal(t, 9, polygon.NumEdges())
}
