package geojson_test

import (
	"encoding/json"
	"testing"

	geojson "github.com/mercatormaps/go-geojson"
	"github.com/stretchr/testify/require"
)

func TestMarshalPropertyList(t *testing.T) {
	data, err := json.Marshal(&geojson.PropertyList{
		geojson.StringProp("prop1", "val1"),
		geojson.StringProp("prop2", "val2"),
	})

	require.NoError(t, err)
	require.JSONEq(t, `
		{
			"prop1": "val1",
			"prop2": "val2"
		}`,
		string(data))
}

func TestUnmarshalPropertyList(t *testing.T) {
	props := geojson.PropertyList{}
	err := json.Unmarshal([]byte(`
		{
			"prop1": "val1",
			"prop2": "val2"
		}`),
		&props)

	require.NoError(t, err)
	require.ElementsMatch(t, geojson.PropertyList{
		geojson.StringProp("prop1", "val1"),
		geojson.StringProp("prop2", "val2"),
	}, props)
}
