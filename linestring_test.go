package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestLineString(t *testing.T) {
	feature := geojson.NewFeature(
		geojson.NewLineString(
			geojson.MakePosition(12, 34),
			geojson.MakePosition(56, 78),
			geojson.MakePosition(90, 12),
		),
	)

	err := feature.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "LineString",
				"coordinates": [
					[34, 12],
					[78, 56],
					[12, 90]
				]
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.LineString]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestLineStringTooShort(t *testing.T) {
	err := geojson.LineString{
		geojson.MakePosition(12, 34),
	}.Validate()
	require.Error(t, err)
}

func TestMultiLineString(t *testing.T) {
	feature := geojson.NewFeature(
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
		),
	)

	err := feature.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
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
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.MultiLineString]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestMultiLineStringTooShort(t *testing.T) {
	require.Error(t, geojson.NewFeature(
		geojson.NewMultiLineString(
			[]geojson.Position{
				geojson.MakePosition(12, 34),
			},
		),
	).Validate())
}
