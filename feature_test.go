package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestFeature(t *testing.T) {
	feature := geojson.NewPoint(9.189982, 45.4642035)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982,45.4642035]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, &unmarshalled)
}

func TestFeatureWithBoundingBox(t *testing.T) {
	feature := geojson.NewPoint(9.189982, 45.4642035).
		WithBoundingBox(
			geojson.NewPosition(7.1827768, 43.7032932),
			geojson.NewPosition(11.2387051, 47.2856026),
		)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"bbox": [
				7.1827768,  43.7032932,
				11.2387051, 47.2856026
			],
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982,45.4642035]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, &unmarshalled)
}

func TestFeatureWithProperties(t *testing.T) {
	feature := geojson.NewPoint(9.189982, 45.4642035).
		AddProperty("city", "Milan")

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982,45.4642035]
			},
			"properties": {
				"city": "Milan"
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, &unmarshalled)
}
