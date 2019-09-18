package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
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

func TestPropertyListGetType(t *testing.T) {
	t.Run("string property", func(t *testing.T) {
		props := geojson.NewPropertyList(
			geojson.Property{Name: "string_prop", Value: "hello"},
		)

		var prop string
		err := props.GetType("string_prop", &prop)
		require.NoError(t, err)
		require.Equal(t, "hello", prop)
	})

	t.Run("integer property", func(t *testing.T) {
		props := geojson.NewPropertyList(
			geojson.Property{Name: "int_prop", Value: 4},
		)

		var prop int
		err := props.GetType("int_prop", &prop)
		require.NoError(t, err)
		require.Equal(t, 4, prop)
	})

	t.Run("incorrect type", func(t *testing.T) {
		props := geojson.NewPropertyList(
			geojson.Property{Name: "float_prop", Value: 4.5},
		)

		var prop int
		err := props.GetType("float_prop", &prop)
		require.Error(t, err)
	})
}
