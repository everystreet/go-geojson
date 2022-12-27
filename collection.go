package geojson

import (
	"encoding/json"
	"fmt"
)

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
	GeometryCollectionType      GeometryType = "GeometryCollection"
)

// GeometryCollection is a heterogeneous collection of Geometry objects.
type GeometryCollection []Geometry

// NewGeometryCollection returns a GeometryCollection Feature.
func NewGeometryCollection(geometries ...Geometry) *GeometryCollection {
	return (*GeometryCollection)(&geometries)
}

// Type returns the geometry type.
func (c GeometryCollection) Type() GeometryType {
	return GeometryCollectionType
}

// Validate the GeometryCollection.
func (c GeometryCollection) Validate() error {
	return nil
}

// MarshalJSON returns the JSON encoding of the GeometryCollection.
func (c GeometryCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(geometryCollection{
		Type:       GeometryCollectionType,
		Geometries: c,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (c *GeometryCollection) UnmarshalJSON(data []byte) error {
	var collection struct {
		Geometries []json.RawMessage `json:"geometries"`
	}

	if err := json.Unmarshal(data, &collection); err != nil {
		return err
	}

	*c = make(GeometryCollection, len(collection.Geometries))
	for i, data := range collection.Geometries {
		geo, err := unmarshalGeometry(data)
		if err != nil {
			return err
		}
		(*c)[i] = geo
	}
	return nil
}

func unmarshalGeometry(data json.RawMessage) (Geometry, error) {
	var typ struct {
		Type GeometryType `json:"type"`
	}

	if err := json.Unmarshal(data, &typ); err != nil {
		return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
	}

	var geo Geometry
	switch typ.Type {
	case PointGeometryType:
		geo = &Point{}
	case MultiPointGeometryType:
		geo = &MultiPoint{}
	case LineStringGeometryType:
		geo = &LineString{}
	case MultiLineStringGeometryType:
		geo = &MultiLineString{}
	case PolygonGeometryType:
		geo = &Polygon{}
	case MultiPolygonGeometryType:
		geo = &MultiPolygon{}
	case GeometryCollectionType:
		geo = &GeometryCollection{}
	default:
		return nil, fmt.Errorf("unknown geometry type '%v'", typ.Type)
	}

	if err := json.Unmarshal(data, geo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal geometry '%v': %w", typ.Type, err)
	}
	return geo, nil
}

type geometry struct {
	Type        GeometryType `json:"type"`
	Coordinates interface{}  `json:"coordinates"`
}

type geometryCollection struct {
	Type       GeometryType `json:"type"`
	Geometries []Geometry   `json:"geometries"`
}
