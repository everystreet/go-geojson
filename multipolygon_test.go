package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMultiPolygon(t *testing.T) {
	polygons := geojson.NewMultiPolygon(
		[][]geojson.Position{
			{
				geojson.NewPosition(7, 7),
				geojson.NewPosition(8, 4),
				geojson.NewPosition(4, 3),
				geojson.NewPosition(2, 5),
				geojson.NewPosition(3, 7),
				geojson.NewPosition(7, 7),
			},
			[]geojson.Position{
				geojson.NewPosition(4, 4),
				geojson.NewPosition(6, 4),
				geojson.NewPosition(7, 5),
				geojson.NewPosition(4, 6),
				geojson.NewPosition(4, 4),
			},
		},
		[][]geojson.Position{
			{
				geojson.NewPosition(7, 7),
				geojson.NewPosition(4, 3),
				geojson.NewPosition(2, 5),
				geojson.NewPosition(7, 7),
			},
		},
	)

	err := polygons.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&polygons)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "MultiPolygon",
				"coordinates": [
					[
						[
							[7, 7],
							[8, 4],
							[4, 3],
							[2, 5],
							[3, 7],
							[7, 7]
						],
						[
							[4, 4],
							[6, 4],
							[7, 5],
							[4, 6],
							[4, 4]
						]
					],
					[
						[
							[7, 7],
							[4, 3],
							[2, 5],
							[7, 7]
						]
					]
				]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, polygons, &unmarshalled)
}

func TestMultiPolygonErrors(t *testing.T) {
	t.Run("ring too short", func(t *testing.T) {
		err := geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.NewPosition(7, 7),
					geojson.NewPosition(8, 4),
					geojson.NewPosition(4, 3),
				},
			}).Geometry.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "is too short")
	})

	t.Run("ring not closed", func(t *testing.T) {
		err := geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.NewPosition(7, 7),
					geojson.NewPosition(8, 4),
					geojson.NewPosition(4, 3),
					geojson.NewPosition(2, 5),
				},
			}).Geometry.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "must be closed")
	})

	t.Run("counter-clockwise exterior ring", func(t *testing.T) {
		err := geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.NewPosition(4, 4),
					geojson.NewPosition(6, 4),
					geojson.NewPosition(7, 5),
					geojson.NewPosition(4, 4),
				},
			}).Geometry.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "exterior ring must be clockwise")
	})

	t.Run("clockwise interior ring", func(t *testing.T) {
		err := geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.NewPosition(7, 7),
					geojson.NewPosition(8, 4),
					geojson.NewPosition(4, 3),
					geojson.NewPosition(7, 7),
				},
				{
					geojson.NewPosition(7, 7),
					geojson.NewPosition(8, 4),
					geojson.NewPosition(4, 3),
					geojson.NewPosition(7, 7),
				},
			}).Geometry.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "interior ring must be counter-clockwise")
	})
}
