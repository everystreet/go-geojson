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
