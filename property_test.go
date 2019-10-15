package geojson_test

import (
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
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
