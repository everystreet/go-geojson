package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/mercatormaps/go-geojson"
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
	t.Run("marshal", func(t *testing.T) {
		_, err := json.Marshal(geojson.NewMultiPolygon(
			[][]geojson.Position{
				{
					geojson.NewPosition(12, 34),
					geojson.NewPosition(56, 78),
					geojson.NewPosition(12, 34),
				},
			}))
		require.Error(t, err)
	})

	t.Run("unmarshal", func(t *testing.T) {
		err := json.Unmarshal([]byte(`
		{
			"type": "Feature",
			"geometry": {
				"type": "MultiPolygon",
				"coordinates": [
					[
						[
							[12, 34],
							[56, 78],
							[12, 34]
						]
					]
				]
			}
		}`), &geojson.Feature{})
		require.Error(t, err)
	})
}
