package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestLineString(t *testing.T) {
	linestring := geojson.NewLineString(
		geojson.NewPosition(12, 34),
		geojson.NewPosition(56, 78),
		geojson.NewPosition(90, 12),
	)

	err := linestring.Geometry.Validate()
	require.NoError(t, err)

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
	err := geojson.LineString{
		geojson.NewPosition(12, 34),
	}.Validate()
	require.Error(t, err)
}

func TestMultiLineString(t *testing.T) {
	linestrings := geojson.NewMultiLineString(
		[]geojson.Position{
			geojson.NewPosition(12, 34),
			geojson.NewPosition(56, 78),
			geojson.NewPosition(90, 12),
		},
		[]geojson.Position{
			geojson.NewPosition(23, 45),
			geojson.NewPosition(67, 89),
		},
	)

	err := linestrings.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&linestrings)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
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
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, linestrings, &unmarshalled)
}

func TestMultiLineStringTooShort(t *testing.T) {
	err := geojson.NewMultiLineString(
		[]geojson.Position{
			geojson.NewPosition(12, 34),
		}).Geometry.Validate()
	require.Error(t, err)
}
