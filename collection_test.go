package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/everystreet/go-geojson/v2"
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
					"coordinates": [45, 9]
				},
				{
					"type": "MultiLineString",
					"coordinates": [
						[
							[34, 12],
							[78, 56],
							[12, 90]
						],
						[
							[45, 23],
							[89, 67]
						]
					]
				},
				{
					"type": "GeometryCollection",
					"geometries": [
						{
							"type": "Point",
							"coordinates": [12, 37]
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
