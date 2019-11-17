package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMultiPoint(t *testing.T) {
	multipoint := geojson.NewMultiPoint(
		geojson.NewPosition(12, 34),
		geojson.NewPositionWithElevation(12, 34, 56),
	)

	err := multipoint.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(multipoint)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "MultiPoint",
				"coordinates": [
					[12, 34],
					[12, 34, 56]
				]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, multipoint, &unmarshalled)
}
