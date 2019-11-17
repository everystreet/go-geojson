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
			geojson.NewPosition(12, 34),
			geojson.NewPosition(56, 78),
			geojson.NewPosition(90, 12),
			geojson.NewPosition(34, 56),
			geojson.NewPosition(12, 34),
		},
		[]geojson.Position{
			geojson.NewPosition(12, 34),
			geojson.NewPosition(56, 78),
			geojson.NewPosition(90, 12),
			geojson.NewPosition(12, 34),
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
						[12, 34],
						[56, 78],
						[90, 12],
						[34, 56],
						[12, 34]
					],
					[
						[12, 34],
						[56, 78],
						[90, 12],
						[12, 34]
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

func TestPolygonTooShort(t *testing.T) {
	err := geojson.NewPolygon(
		[]geojson.Position{
			geojson.NewPosition(12, 34),
			geojson.NewPosition(56, 78),
			geojson.NewPosition(12, 34),
		}).Geometry.Validate()
	require.Error(t, err)
}
