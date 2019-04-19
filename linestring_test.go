package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestLineString(t *testing.T) {
	linestring := geojson.NewLineString(
		geojson.NewPosition(12, 34),
		geojson.NewPosition(56, 78),
		geojson.NewPosition(90, 12),
	)

	data, err := json.Marshal(linestring)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "LineString",
				"coordinates": [
					[12, 34],
					[56, 78],
					[90, 12]
				]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, linestring, &unmarshalled)
}

func TestLineStringTooShort(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		_, err := json.Marshal(&geojson.LineString{
			geojson.NewPosition(12, 34),
		})
		require.Error(t, err)
	})

	t.Run("unmarshal", func(t *testing.T) {
		err := json.Unmarshal([]byte(`
			{
				"type": "Feature",
				"geometry": {
					"type": "LineString",
					"coordinates": [
						[12, 34]
					]
				}
			}`), &geojson.Feature{})
		require.Error(t, err)
	})
}
