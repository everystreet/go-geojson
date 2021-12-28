package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestPolygon(t *testing.T) {
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

	err := feature.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Polygon",
				"coordinates": [
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
				]
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.Polygon]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestMultiPolygon(t *testing.T) {
	feature := geojson.Feature[*geojson.MultiPolygon]{
		Geometry: geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.MakePosition(7, 7),
					geojson.MakePosition(4, 8),
					geojson.MakePosition(3, 4),
					geojson.MakePosition(5, 2),
					geojson.MakePosition(7, 3),
					geojson.MakePosition(7, 7),
				},
				{
					geojson.MakePosition(4, 4),
					geojson.MakePosition(4, 6),
					geojson.MakePosition(5, 7),
					geojson.MakePosition(6, 4),
					geojson.MakePosition(4, 4),
				},
			},
			[][]geojson.Position{
				{
					geojson.MakePosition(7, 7),
					geojson.MakePosition(3, 4),
					geojson.MakePosition(5, 2),
					geojson.MakePosition(7, 7),
				},
			},
		),
	}

	err := feature.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&feature)
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
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.MultiPolygon]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestMultiPolygonErrors(t *testing.T) {
	t.Run("ring too short", func(t *testing.T) {
		err := geojson.Feature[*geojson.MultiPolygon]{
			Geometry: geojson.NewMultiPolygon(
				[][]geojson.Position{
					{
						geojson.MakePosition(7, 7),
						geojson.MakePosition(4, 8),
						geojson.MakePosition(3, 4),
					},
				}),
		}.Geometry.Validate()

		require.Error(t, err)
		require.Contains(t, err.Error(), "is too short")
	})

	t.Run("ring not closed", func(t *testing.T) {
		err := geojson.Feature[*geojson.MultiPolygon]{
			Geometry: geojson.NewMultiPolygon(
				[][]geojson.Position{
					{
						geojson.MakePosition(7, 7),
						geojson.MakePosition(4, 8),
						geojson.MakePosition(3, 4),
						geojson.MakePosition(5, 2),
					},
				}),
		}.Geometry.Validate()

		require.Error(t, err)
		require.Contains(t, err.Error(), "must be closed")
	})

	t.Run("counter-clockwise exterior ring", func(t *testing.T) {
		err := geojson.Feature[*geojson.MultiPolygon]{
			Geometry: geojson.NewMultiPolygon(
				[][]geojson.Position{
					{
						geojson.MakePosition(4, 4),
						geojson.MakePosition(4, 6),
						geojson.MakePosition(5, 7),
						geojson.MakePosition(4, 4),
					},
				}),
		}.Geometry.Validate()

		require.Error(t, err)
		require.Contains(t, err.Error(), "exterior ring must be clockwise")
	})

	t.Run("clockwise interior ring", func(t *testing.T) {
		err := geojson.Feature[*geojson.MultiPolygon]{
			Geometry: geojson.NewMultiPolygon(
				[][]geojson.Position{
					{
						geojson.MakePosition(7, 7),
						geojson.MakePosition(4, 8),
						geojson.MakePosition(3, 4),
						geojson.MakePosition(7, 7),
					},
					{
						geojson.MakePosition(7, 7),
						geojson.MakePosition(4, 8),
						geojson.MakePosition(3, 4),
						geojson.MakePosition(7, 7),
					},
				}),
		}.Geometry.Validate()

		require.Error(t, err)
		require.Contains(t, err.Error(), "interior ring must be counter-clockwise")
	})
}
