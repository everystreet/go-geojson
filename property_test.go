package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/everystreet/go-geojson/v3"
	"github.com/stretchr/testify/require"
)

func TestPropertyGetValue(t *testing.T) {
	t.Run("string property", func(t *testing.T) {
		prop := geojson.Property{Name: "my_prop", Value: "hello"}

		var value string
		err := prop.GetValue(&value)
		require.NoError(t, err)
		require.Equal(t, "hello", value)
	})

	t.Run("integer property", func(t *testing.T) {
		prop := geojson.Property{Name: "my_prop", Value: 4}

		var value int
		err := prop.GetValue(&value)
		require.NoError(t, err)
		require.Equal(t, 4, value)
	})

	t.Run("incorrect type", func(t *testing.T) {
		prop := geojson.Property{Name: "my_prop", Value: 4.5}

		var value int
		err := prop.GetValue(&value)
		require.Error(t, err)
	})
}

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
