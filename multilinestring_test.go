package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMultiLineString(t *testing.T) {
	mls := geojson.NewMultiLineString(
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

	data, err := json.Marshal(&mls)
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
	require.Equal(t, mls, &unmarshalled)
}
