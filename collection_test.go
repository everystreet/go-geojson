package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestMarshalFeatureCollection(t *testing.T) {
	data, err := json.Marshal(geojson.NewFeatureCollection(
		geojson.NewPoint(9.189982, 45.4642035),
		geojson.NewPoint(79.9288064, 13.0473748),
	))

	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "FeatureCollection",
			"features": [
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [9.189982,45.4642035]
					}
				},
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [79.9288064,13.0473748]
					}
				}
			]
		}`,
		string(data))
}

func TestUnmarshalFeatureCollection(t *testing.T) {
	coll := geojson.FeatureCollection{}
	err := json.Unmarshal([]byte(`
		{
			"type": "FeatureCollection",
			"features": [
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [9.189982,45.4642035]
					}
				},
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [79.9288064,13.0473748]
					}
				}
			]
		}`), &coll)

	require.NoError(t, err)
	require.Equal(t, geojson.NewFeatureCollection(
		geojson.NewPoint(9.189982, 45.4642035),
		geojson.NewPoint(79.9288064, 13.0473748),
	), &coll)
}
