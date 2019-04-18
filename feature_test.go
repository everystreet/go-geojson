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
			"simple",
			geojson.NewPoint(9.189982, 45.4642035),
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				}
			}`,
		},
		{
			"with bbox",
			geojson.NewPoint(9.189982, 45.4642035).WithBoundingBox(
				geojson.NewCoordinates(7.1827768, 43.7032932),
				geojson.NewCoordinates(11.2387051, 47.2856026),
			),
			`{
				"type": "Feature",
				"bbox": [
					7.1827768,  43.7032932,
					11.2387051, 47.2856026
				],
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				} 
			}`,
		},
		{
			"with properties",
			geojson.NewPoint(9.189982, 45.4642035).AddProperty("city", "Milan"),
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
			"simple",
			`{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				}
			}`,
			geojson.NewPoint(9.189982, 45.4642035),
		},
		{
			"with bbox",
			`{
				"type": "Feature",
				"bbox": [
					7.1827768,  43.7032932,
					11.2387051, 47.2856026
				],
				"geometry": {
					"type": "Point",
					"coordinates": [9.189982,45.4642035]
				} 
			}`,
			geojson.NewPoint(9.189982, 45.4642035).WithBoundingBox(
				geojson.NewCoordinates(7.1827768, 43.7032932),
				geojson.NewCoordinates(11.2387051, 47.2856026),
			),
		},
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
			geojson.NewPoint(9.189982, 45.4642035).AddProperty("city", "Milan"),
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
