package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestPolygon(t *testing.T) {
	p := geojson.NewPolygon(
		[]geojson.Position{
			geojson.NewPosition(7, 7),
			geojson.NewPosition(8, 4),
			geojson.NewPosition(4, 3),
			geojson.NewPosition(2, 5),
			geojson.NewPosition(3, 7),
			geojson.NewPosition(7, 7),
		},
		[]geojson.Position{
			geojson.NewPosition(4, 4),
			geojson.NewPosition(6, 4),
			geojson.NewPosition(7, 5),
			geojson.NewPosition(4, 6),
			geojson.NewPosition(4, 4),
		},
	)

	err := p.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(&p)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Polygon",
				"coordinates": [
					[
						[7, 7],
						[8, 4],
						[4, 3],
						[2, 5],
						[3, 7],
						[7, 7]
					],
					[
						[4, 4],
						[6, 4],
						[7, 5],
						[4, 6],
						[4, 4]
					]
				]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, p, &unmarshalled)
}
