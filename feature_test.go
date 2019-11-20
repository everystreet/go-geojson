package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
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

func TestFeatureCollection(t *testing.T) {
	collection := geojson.NewFeatureCollection(
		geojson.NewPoint(9.189982, 45.4642035),
		geojson.NewPoint(79.9288064, 13.0473748),
	)

	data, err := json.Marshal(collection)
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

	unmarshalled := geojson.FeatureCollection{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, collection, &unmarshalled)
}

func TestFeatureCollectionWithBoundingBox(t *testing.T) {
	collection := geojson.NewFeatureCollection(
		geojson.NewPoint(9.189982, 45.4642035),
		geojson.NewPoint(79.9288064, 13.0473748),
	).WithBoundingBox(
		geojson.NewPosition(7.1827768, 43.7032932),
		geojson.NewPosition(11.2387051, 47.2856026),
	)

	data, err := json.Marshal(collection)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "FeatureCollection",
			"bbox": [
				7.1827768,  43.7032932,
				11.2387051, 47.2856026
			],
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

	unmarshalled := geojson.FeatureCollection{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, collection, &unmarshalled)
}
