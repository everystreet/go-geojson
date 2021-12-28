package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestFeature(t *testing.T) {
	feature := geojson.NewPoint(45.4642035, 9.189982)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, &unmarshalled)
}

func TestFeatureWithBoundingBox(t *testing.T) {
	feature := geojson.NewPoint(45.4642035, 9.189982).WithBoundingBox(
		geojson.MakePosition(43.7032932, 7.1827761),
		geojson.MakePosition(47.2856026, 11.2387051),
	)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"bbox": [
				7.1827761,  43.7032932,
				11.2387051, 47.2856026
			],
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, &unmarshalled)
}

func TestFeatureWithProperties(t *testing.T) {
	feature := geojson.NewPoint(45.4642035, 9.189982).
		AddProperty("city", "Milan")

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035]
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
		geojson.NewPoint(45.4642035, 9.189982),
		geojson.NewPoint(13.0473748, 79.9288064),
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
						"coordinates": [9.189982, 45.4642035]
					}
				},
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [79.9288064, 13.0473748]
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
		geojson.NewPoint(45.4642035, 9.189982),
		geojson.NewPoint(13.0473748, 79.9288064),
	).WithBoundingBox(
		geojson.MakePosition(43.7032932, 7.1827761),
		geojson.MakePosition(47.2856026, 11.2387051),
	)

	data, err := json.Marshal(collection)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "FeatureCollection",
			"bbox": [
				7.1827761,  43.7032932,
				11.2387051, 47.2856026
			],
			"features": [
				{
					"type": "Feature",
					"geometry": {
						"type": "Point",
						"coordinates": [9.189982, 45.4642035]
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
