package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestPropertyList(t *testing.T) {
	props := geojson.NewPropertyList(
		geojson.Property{Name: "prop1", Value: "val1"},
		geojson.Property{Name: "prop2", Value: "val2"},
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

	require.Len(t, unmarshalled, 2)
	require.Contains(t, unmarshalled, props[0])
	require.Contains(t, unmarshalled, props[1])
}
