package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	feature := geojson.NewFeature(
		geojson.NewPoint(45.4642035, 9.189982),
	)

	err := feature.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035]
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.Point]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestPointWithElevation(t *testing.T) {
	feature := geojson.NewFeature(
		geojson.NewPointWithElevation(45.4642035, 9.189982, 125),
	)

	data, err := json.Marshal(feature)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035, 125]
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.Point]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}

func TestMultiPoint(t *testing.T) {
	feature := geojson.NewFeature(
		geojson.NewMultiPoint(
			geojson.MakePosition(12, 34),
			geojson.MakePositionWithElevation(56, 78, 4),
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
				"type": "MultiPoint",
				"coordinates": [
					[34, 12],
					[78, 56, 4]
				]
			}
		}`, string(data))

	var unmarshalled geojson.Feature[*geojson.MultiPoint]
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, feature, unmarshalled)
}
