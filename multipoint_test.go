package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestMarshal(t *testing.T) {
	data, err := json.Marshal(geojson.NewMultiPoint([]geojson.Coordinates{
		{Longitude: 12, Latitude: 34},
		{Longitude: 12, Latitude: 34, Elevation: geojson.NewOptionalFloat64(56)},
	}))
	require.NoError(t, err)
	require.JSONEq(t,
		`{
			"type": "Feature",
			"geometry": {
				"type": "MultiPoint",
				"coordinates": [
					[12, 34],
					[12, 34, 56]
				]
			}
		}`,
		string(data))
}

func TestUnmarshal(t *testing.T) {
	f := geojson.Feature{}
	err := json.Unmarshal([]byte(
		`{
			"type": "Feature",
			"geometry": {
				"type": "MultiPoint",
				"coordinates": [
					[12, 34],
					[12, 34, 56]
				]
			}
		}`), &f)
	require.NoError(t, err)
	require.Equal(t, geojson.NewMultiPoint([]geojson.Coordinates{
		{Longitude: 12, Latitude: 34},
		{Longitude: 12, Latitude: 34, Elevation: geojson.NewOptionalFloat64(56)},
	}), &f)
}