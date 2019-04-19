package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestMultiPoint(t *testing.T) {
	multipoint := geojson.NewMultiPoint(
		geojson.NewPosition(12, 34),
		geojson.NewPositionWithElevation(12, 34, 56),
	)

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
