package geojson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	geojson "github.com/mercatormaps/go-geojson"
)

func TestMarshalFeature(t *testing.T) {
	tests := []struct {
		name     string
		feature  *geojson.Feature
		expected string
	}{
		{
			"with properties",
			geojson.NewPoint(9.189982, 45.4642035, geojson.StringProp("city", "Milan")),
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				},
				"properties": {
					"city": "Milan"
				}
			}`,
		},
		{
			"without properties",
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
			data, err := json.Marshal(tt.feature)
			require.NoError(t, err)
			require.JSONEq(t, tt.expected, string(data))
		})
	}
}

func TestUnmarshalFeature(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *geojson.Feature
	}{
		{
			"with properties",
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				},
				"properties": {
					"city": "Milan"
				}
			}`,
			geojson.NewPoint(9.189982, 45.4642035, geojson.StringProp("city", "Milan")),
		},
		{
			"without properties",
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
