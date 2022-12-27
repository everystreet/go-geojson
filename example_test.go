package geojson_test

import (
	"encoding/json"
	"fmt"

	geojson "github.com/everystreet/go-geojson/v3"
)

func ExampleMarshal() {
	feature := geojson.NewFeatureWithBoundingBox(
		geojson.NewLineString(
			geojson.MakePosition(34, 12),
			geojson.MakePosition(78, 56),
			geojson.MakePosition(12, 90),
		),
		geojson.BoundingBox{
			BottomLeft: geojson.MakePosition(1, 1),
			TopRight:   geojson.MakePosition(100, 100),
		},
		geojson.Property{
			Name:  "foo",
			Value: "bar",
		},
	)

	data, _ := json.Marshal(feature)
	fmt.Println(string(data))
	// Output: {"type":"Feature","bbox":[1,1,100,100],"geometry":{"type":"LineString","coordinates":[[12,34],[56,78],[90,12]]},"properties":{"foo":"bar"}}
}

func ExampleUnmarshal() {
	var feature geojson.Feature[geojson.Geometry]

	_ = json.Unmarshal([]byte(`
		{
			"type": "Feature",
			"geometry": {
				"type": "LineString",
				"coordinates": [
					[12, 34],
					[56, 78],
					[90, 12]
				]
			}
		}`),
		&feature,
	)

	fmt.Println(feature.Geometry().Type())
	// Output: LineString
}
