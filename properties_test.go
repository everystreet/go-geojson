package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestPropertyList(t *testing.T) {
	props := geojson.NewPropertyList(
		geojson.StringProp("prop1", "val1"),
		geojson.StringProp("prop2", "val2"),
	)

	data, err := json.Marshal(&props)
	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"prop1": "val1",
			"prop2": "val2"
		}`,
		string(data))

	unmarshalled := geojson.PropertyList{}
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	require.Equal(t, &props, &unmarshalled)
}
