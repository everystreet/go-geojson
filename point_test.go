package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	point := geojson.NewPoint(9.189982, 45.4642035)

	err := point.Geometry.Validate()
	require.NoError(t, err)

	data, err := json.Marshal(point)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, point, &unmarshalled)
}

func TestPointWithElevation(t *testing.T) {
	point := geojson.NewPointWithElevation(9.189982, 45.4642035, 125)

	data, err := json.Marshal(point)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [9.189982, 45.4642035, 125]
			}
		}`,
		string(data))

	unmarshalled := geojson.Feature{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, point, &unmarshalled)
}
