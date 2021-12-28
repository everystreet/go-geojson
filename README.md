# go-geojson

`go-geojson` is a Go package for working with the GeoJSON format, as standardized by [RFC 7946](https://tools.ietf.org/html/rfc7946).

This package supports marshalling and unmarshalling of all geometry types: `Point`, `MultiPoint`, `LineString`, `MultiLineString`, `Polygon`, `MultiPolygon` and `GeometryCollection`.

## Usage

1. The command below adds the latest version of `go-geojson` as a dependency to your module:

    ```bash
    go get -u github.com/everystreet/go-geojson/v3
    ```

2. Then the following line can be used to import the package and use the examples below:

    ```go
    import "github.com/everystreet/go-geojson/v3"
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

GeoJSON features can be created to contain any of the supported geometry types. Additionally, an optional bounding box and property list can be added. Once the feature is constructed, it can be marshaled to JSON.

```go
feature := geojson.Feature[*geojson.LineString]{
    Geometry: geojson.NewLineString(
        geojson.MakePosition(34, 12),
        geojson.MakePosition(78, 56),
        geojson.MakePosition(12, 90),
    ),
    BBox: &geojson.BoundingBox{
        BottomLeft: geojson.MakePosition(1, 1),
        TopRight:   geojson.MakePosition(100, 100),
    },
    Properties: geojson.NewPropertyList(
        geojson.Property{
            Name:  "foo",
            Value: "bar",
        },
    ),
}

data, _ := json.Marshal(linestring)
```
