package geojson

import (
	"encoding/json"
	"fmt"
)

// TypePropFeature is the value of the "type" property for Features.
const TypePropFeature = "Feature"

// GeometryType is a supported geometry type.
type GeometryType string

// Types of geometry.
const (
	PointGeometryType           GeometryType = "Point"
	MultiPointGeometryType      GeometryType = "MultiPoint"
	LineStringGeometryType      GeometryType = "LineString"
	MultiLineStringGeometryType GeometryType = "MultiLineString"
	PolygonGeometryType         GeometryType = "Polygon"
	MultiPolygonGeometryType    GeometryType = "MultiPolygon"
)

// Feature consists of a specific geometry type and a list of properties.
type Feature struct {
	Geometry   Geometry
	BBox       *BoundingBox
	Properties PropertyList
}

// Geometry contains the points represented by a particular geometry type.
type Geometry interface {
	Type() GeometryType
	json.Marshaler
	json.Unmarshaler
}

// MarshalJSON returns the JSON encoding of the Feature.
func (f *Feature) MarshalJSON() ([]byte, error) {
	geom := geo{
		Type: f.Geometry.Type(),
		Pos:  f.Geometry,
	}

	var feat interface{}
	if len(f.Properties) > 0 {
		feat = struct {
			Type  string        `json:"type"`
			BBox  *BoundingBox  `json:"bbox,omitempty"`
			Geo   geo           `json:"geometry"`
			Props *PropertyList `json:"properties"`
		}{
			Type:  TypePropFeature,
			BBox:  f.BBox,
			Geo:   geom,
			Props: &f.Properties,
		}
	} else {
		feat = struct {
			Type string       `json:"type"`
			BBox *BoundingBox `json:"bbox,omitempty"`
			Geo  geo          `json:"geometry"`
		}{
			Type: TypePropFeature,
			BBox: f.BBox,
			Geo:  geom,
		}
	}
	return json.Marshal(&feat)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (f *Feature) UnmarshalJSON(data []byte) error {
	var objs map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objs); err != nil {
		return err
	}

	var typ string
	if data, ok := objs["type"]; !ok {
		return fmt.Errorf("missing 'type'")
	} else if err := json.Unmarshal(*data, &typ); err != nil {
		return fmt.Errorf("failed to unmarshal 'type': %w", err)
	} else if typ != TypePropFeature {
		return fmt.Errorf("type is '%s', expecting '%s'", typ, TypePropFeature)
	}

	if data, ok := objs["bbox"]; ok {
		f.BBox = &BoundingBox{}
		if err := json.Unmarshal(*data, f.BBox); err != nil {
			return fmt.Errorf("failed to unmarshal 'bbox' (bounding box): %w", err)
		}
	}

	if data, ok := objs["properties"]; ok {
		if err := json.Unmarshal(*data, &f.Properties); err != nil {
			return fmt.Errorf("failed to unmarshal 'properties': %w", err)
		}
	}

	geo := struct {
		Type GeometryType     `json:"type"`
		Pos  *json.RawMessage `json:"coordinates"`
	}{}

	if data, ok := objs["geometry"]; !ok {
		return fmt.Errorf("missing 'geometry'")
	} else if err := json.Unmarshal(*data, &geo); err != nil {
		return fmt.Errorf("failed to unmarshal 'geometry': %w", err)
	}

	switch geo.Type {
	case PointGeometryType:
		p := Point{}
		if err := json.Unmarshal(*geo.Pos, &p); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", PointGeometryType, err)
		}
		f.Geometry = &p
	case MultiPointGeometryType:
		m := MultiPoint{}
		if err := json.Unmarshal(*geo.Pos, &m); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", MultiPointGeometryType, err)
		}
		f.Geometry = &m
	case LineStringGeometryType:
		l := LineString{}
		if err := json.Unmarshal(*geo.Pos, &l); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", LineStringGeometryType, err)
		}
		f.Geometry = &l
	case MultiLineStringGeometryType:
		m := MultiLineString{}
		if err := json.Unmarshal(*geo.Pos, &m); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", MultiLineStringGeometryType, err)
		}
		f.Geometry = &m
	case PolygonGeometryType:
		p := Polygon{}
		if err := json.Unmarshal(*geo.Pos, &p); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", PolygonGeometryType, err)
		}
		f.Geometry = &p
	case MultiPolygonGeometryType:
		m := MultiPolygon{}
		if err := json.Unmarshal(*geo.Pos, &m); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", MultiPolygonGeometryType, err)
		}
		f.Geometry = &m
	default:
		return fmt.Errorf("unknown geometry type %s", geo.Type)
	}

	return nil
}

// WithBoundingBox sets the optional bounding box.
func (f *Feature) WithBoundingBox(bottomLeft, topRight Position) *Feature {
	f.BBox = &BoundingBox{
		BottomLeft: bottomLeft,
		TopRight:   topRight,
	}
	return f
}

// WithProperties sets the optional properties, removing all existing properties.
func (f *Feature) WithProperties(props ...Property) *Feature {
	f.Properties = PropertyList(props)
	return f
}

// AddProperty appends a new property.
func (f *Feature) AddProperty(name string, value interface{}) *Feature {
	f.Properties = append(f.Properties, Property{
		Name:  name,
		Value: value,
	})
	return f
}

type feature struct {
	Type  string       `json:"type"`
	Geo   geo          `json:"geometry"`
	Props PropertyList `json:"properties"`
}

type geo struct {
	Type GeometryType `json:"type"`
	Pos  interface {
		json.Marshaler
		json.Unmarshaler
	} `json:"coordinates"`
}
