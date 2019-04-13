package geojson

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	// FeatureType string.
	FeatureType = "Feature"

	// PointType string.
	PointType = "Point"
	// MultiPointType string.
	MultiPointType = "MultiPoint"
	// LineStringType string.
	LineStringType = "LineString"
	// MultiLineStringType string.
	MultiLineStringType = "MultiLineString"
	// PolygonType string.
	PolygonType = "Polygon"
	// MultiPolygonType string.
	MultiPolygonType = "MultiPolygon"
)

// Feature consists of a specific geometry type and a list of properties.
type Feature struct {
	Geometry interface {
		json.Marshaler
		json.Unmarshaler
	}
	Properties PropertyList
}

// MarshalJSON returns the JSON encoding of the Feature.
func (f *Feature) MarshalJSON() ([]byte, error) {
	var geom geo

	switch g := f.Geometry.(type) {
	case *Point:
		geom.Type = PointType
		geom.Coords = g
	case *MultiPoint:
		fmt.Println("MultiPoint")
		geom.Type = MultiPointType
		geom.Coords = g
	default:
		return nil, fmt.Errorf("unknown geometry type: %v", g)
	}

	var feat interface{}
	if len(f.Properties) > 0 {
		feat = struct {
			Type  string        `json:"type"`
			Geo   geo           `json:"geometry"`
			Props *PropertyList `json:"properties"`
		}{
			Type:  FeatureType,
			Geo:   geom,
			Props: &f.Properties,
		}
	} else {
		feat = struct {
			Type string `json:"type"`
			Geo  geo    `json:"geometry"`
		}{
			Type: FeatureType,
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
		return errors.New("missing 'type'")
	} else if err := json.Unmarshal(*data, &typ); err != nil {
		return errors.Wrap(err, "failed to unmarshal 'type'")
	} else if typ != FeatureType {
		return fmt.Errorf("type is '%s', expecting '%s'", typ, FeatureType)
	}

	if data, ok := objs["properties"]; ok {
		if err := json.Unmarshal(*data, &f.Properties); err != nil {
			return errors.Wrap(err, "failed to unmarshal 'properties'")
		}
	}

	geo := struct {
		Type   string           `json:"type"`
		Coords *json.RawMessage `json:"coordinates"`
	}{}

	if data, ok := objs["geometry"]; !ok {
		return errors.New("missing 'geometry'")
	} else if err := json.Unmarshal(*data, &geo); err != nil {
		return errors.Wrap(err, "failed to unmarshal 'geometry'")
	}

	switch geo.Type {
	case PointType:
		p := Point{}
		if err := json.Unmarshal(*geo.Coords, &p); err != nil {
			return errors.Wrap(err, "failed to unmarshal "+PointType)
		}
		f.Geometry = &p
	case MultiPointType:
		m := MultiPoint{}
		if err := json.Unmarshal(*geo.Coords, &m); err != nil {
			return errors.Wrap(err, "failed to unmarshal "+MultiPointType)
		}
		f.Geometry = &m
	default:
		return errors.New("unknown geometry type " + geo.Type)
	}

	return nil
}

type feature struct {
	Type  string       `json:"type"`
	Geo   geo          `json:"geometry"`
	Props PropertyList `json:"properties"`
}

type geo struct {
	Type   string `json:"type"`
	Coords interface {
		json.Marshaler
		json.Unmarshaler
	} `json:"coordinates"`
}
