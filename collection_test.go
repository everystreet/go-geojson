package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestGeometryCollection(t *testing.T) {
	collection := geojson.NewGeometryCollection(
		geojson.NewPoint(9, 45).Geometry,
		geojson.NewMultiLineString(
			[]geojson.Position{
				geojson.MakePosition(12, 34),
				geojson.MakePosition(56, 78),
				geojson.MakePosition(90, 12),
			},
			[]geojson.Position{
				geojson.MakePosition(23, 45),
				geojson.MakePosition(67, 89),
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
