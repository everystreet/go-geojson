package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMultiPolygon(t *testing.T) {
	m := geojson.NewMultiPolygon(
		[][]geojson.Position{
			{
				geojson.NewPosition(12, 34),
				geojson.NewPosition(56, 78),
				geojson.NewPosition(90, 12),
				geojson.NewPosition(12, 34),
			},
			[]geojson.Position{
				geojson.NewPosition(12, 34),
				geojson.NewPosition(56, 78),
				geojson.NewPosition(90, 12),
				geojson.NewPosition(12, 34),
			},
		},
		[][]geojson.Position{
			{
				geojson.NewPosition(12, 34),
				geojson.NewPosition(56, 78),
				geojson.NewPosition(90, 12),
				geojson.NewPosition(12, 34),
			},
		},
	)

	err := m.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&m)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "MultiPolygon",
				"coordinates": [
					[
						[
							[12, 34],
							[56, 78],
							[90, 12],
							[12, 34]
						],
						[
							[12, 34],
							[56, 78],
							[90, 12],
							[12, 34]
						]
					],
					[
						[
							[12, 34],
							[56, 78],
							[90, 12],
							[12, 34]
						]
					]
				]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, m, &unmarshalled)
}

func TestMultiPolygonTooShort(t *testing.T) {
	err := geojson.NewMultiPolygon(
		[][]geojson.Position{
			{
				geojson.NewPosition(12, 34),
				geojson.NewPosition(56, 78),
				geojson.NewPosition(12, 34),
			},
		}).Geometry.Validate()
	require.Error(t, err)
}
