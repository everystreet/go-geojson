package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMarshalCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		coords   geojson.Coordinates
		expected string
	}{
		{
			"with elevation",
			geojson.Coordinates{
				Longitude: 9.189982,
				Latitude:  45.4642035,
				Elevation: geojson.NewOptionalFloat64(125),
			},
			"[9.189982,45.4642035,125]",
		},
		{
			"without elevation",
			geojson.Coordinates{
				Longitude: 9.189982,
				Latitude:  45.4642035,
			},
			"[9.189982,45.4642035]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(&tt.coords)
			require.NoError(t, err)
			require.JSONEq(t, tt.expected, string(data))
		})
	}
}

func TestUnmarshalCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected geojson.Coordinates
	}{
		{
			"with elevation",
			"[9.189982,45.4642035,125]",
			geojson.Coordinates{
				Longitude: 9.189982,
				Latitude:  45.4642035,
				Elevation: geojson.NewOptionalFloat64(125),
			},
		},
		{
			"without elevation",
			"[9.189982,45.4642035]",
			geojson.Coordinates{
				Longitude: 9.189982,
				Latitude:  45.4642035,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coords := geojson.Coordinates{}
			err := json.Unmarshal([]byte(tt.data), &coords)
			require.NoError(t, err)
			require.Equal(t, tt.expected, coords)
		})
	}
}
