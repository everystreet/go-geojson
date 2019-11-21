# go-geojson

`go-geojson` is a Go package for working with the GeoJSON format, as standardized by [RFC 7946](https://tools.ietf.org/html/rfc7946).

This package supports marshalling and unmarshalling of all geometry types: `Point`, `MultiPoint`, `LineString`, `MultiLineString`, `Polygon`, `MultiPolygon` and `GeometryCollection`.

## Usage

1. The command below adds the latest version of `go-geojson` as a dependency to your module:

    ```bash
    go get -u github.com/everystreet/go-geojson/v2
    ```

2. Then the following line can be used to import the package and use the examples below:

    ```go
    import "github.com/everystreet/go-geojson/v2"
    ```

## Examples

`go-geojson` implements the `json.Marshaler` and `json.Unmarshaler` interfaces. This means you can work with GeoJSON in the same way as you would with "regular" JSON.

### Unmarshal

The example below demonstrates how to unmarshal a GeoJSON Feature. Once unmarshalled into a `geojson.Feature`, you have access to the bounding box and properties, and the Geometry object. As the Geometry can be one of several types, a type switch can be used to determine the type and work with it.

```go
var feature geojson.Feature

_ = json.Unmarshal(`
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
    }`, &feature)

switch f := feature.Geometry.(type) {
case *geojson.LineString:
    for _, pos := range *f {
        fmt.Println(pos)
    }
}
```

### Marshal

The easiest way to create GeoJSON objects is using the provided helpers. The example below demonstrates creation of a simple LineString using `geojson.NewLineString`. The returned `geojson.Feature` type contains methods to add a bounding box and properties.

```go
linestring := geojson.NewLineString(
    geojson.MakePosition(34, 12),
    geojson.MakePosition(78, 56),
    geojson.MakePosition(12, 90),
).WithBoundingBox( // optionally set bounding box
    geojson.MakePosition(1, 1),
    geojson.MakePosition(100, 100),
)

data, _ := json.Marshal(linestring)
```
