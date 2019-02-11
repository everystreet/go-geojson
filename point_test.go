package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestMarshalPoint(t *testing.T) {
	tests := []struct {
		name     string
		point    *geojson.Feature
		expected string
	}{
		{
			"with elevation",
			geojson.NewPointWithElevation(9.189982, 45.4642035, 125),
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035,125]
				}
			}`,
		},
		{
			"without elevation",
			geojson.NewPoint(9.189982, 45.4642035),
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				}
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.point)
			require.NoError(t, err)
			require.JSONEq(t, tt.expected, string(data))
		})
	}
}

func TestUnmarshalPoint(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *geojson.Feature
	}{
		{
			"with elevation",
			`{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982,45.4642035,125]
			}
		}`,
			geojson.NewPointWithElevation(9.189982, 45.4642035, 125),
		},
		{
			"without elevation",
			`{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982,45.4642035]
			}
		}`,
			geojson.NewPoint(9.189982, 45.4642035),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feature := geojson.Feature{}
			err := json.Unmarshal([]byte(tt.data), &feature)
			require.NoError(t, err)
			require.Equal(t, tt.expected, &feature)
		})
	}
}
