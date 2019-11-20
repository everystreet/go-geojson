package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestGeometryCollection(t *testing.T) {
	collection := geojson.NewGeometryCollection(
		geojson.NewPoint(9, 45).Geometry,
		geojson.NewMultiLineString(
			[]geojson.Position{
				geojson.NewPosition(12, 34),
				geojson.NewPosition(56, 78),
				geojson.NewPosition(90, 12),
			},
			[]geojson.Position{
				geojson.NewPosition(23, 45),
				geojson.NewPosition(67, 89),
			},
		).Geometry,
		geojson.NewGeometryCollection(
			geojson.NewPoint(37, 12).Geometry,
		).Geometry,
	)

	err := collection.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(collection)
	require.NoError(t, err)
	require.JSONEq(t, `
	{
		"type": "Feature",
		"geometry": {
			"type": "GeometryCollection",
			"geometries": [
				{
					"type": "Point",
					"coordinates": [9, 45]
				},
				{
					"type": "MultiLineString",
					"coordinates": [
						[
							[12, 34],
							[56, 78],
							[90, 12]
						],
						[
							[23, 45],
							[67, 89]
						]
					]
				},
				{
					"type": "GeometryCollection",
					"geometries": [
						{
							"type": "Point",
							"coordinates": [37, 12]
						}
					]
				}
			]
		}
	}`, string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, collection, &unmarshalled)
}
