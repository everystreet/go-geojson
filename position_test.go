package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestPosition(t *testing.T) {
	pos := geojson.MakePosition(45.4642035, 9.189982)

	data, err := json.Marshal(&pos)
	require.NoError(t, err)
	require.JSONEq(t, "[9.189982, 45.4642035]", string(data))

	unmarshalled := geojson.Position{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, &pos, &unmarshalled)
}

func TestPositionWithElevation(t *testing.T) {
	pos := geojson.MakePositionWithElevation(45.4642035, 9.189982, 125)

	data, err := json.Marshal(&pos)
	require.NoError(t, err)
	require.JSONEq(t, "[9.189982, 45.4642035, 125]", string(data))

	unmarshalled := geojson.Position{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, &pos, &unmarshalled)
}
