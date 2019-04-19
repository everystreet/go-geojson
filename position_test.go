package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMarshalPosition(t *testing.T) {
	tests := []struct {
		name     string
		pos      geojson.Position
		expected string
	}{
		{
			"with elevation",
			geojson.Position{
				Longitude: 9.189982,
				Latitude:  45.4642035,
				Elevation: geojson.NewOptionalFloat64(125),
			},
			"[9.189982,45.4642035,125]",
		},
		{
			"without elevation",
			geojson.Position{
				Longitude: 9.189982,
				Latitude:  45.4642035,
			},
			"[9.189982,45.4642035]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(&tt.pos)
			require.NoError(t, err)
			require.JSONEq(t, tt.expected, string(data))
		})
	}
}

func TestUnmarshalPosition(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected geojson.Position
	}{
		{
			"with elevation",
			"[9.189982,45.4642035,125]",
			geojson.Position{
				Longitude: 9.189982,
				Latitude:  45.4642035,
				Elevation: geojson.NewOptionalFloat64(125),
			},
		},
		{
			"without elevation",
			"[9.189982,45.4642035]",
			geojson.Position{
				Longitude: 9.189982,
				Latitude:  45.4642035,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := geojson.Position{}
			err := json.Unmarshal([]byte(tt.data), &pos)
			require.NoError(t, err)
			require.Equal(t, tt.expected, pos)
		})
	}
}
