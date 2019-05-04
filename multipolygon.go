package geojson

import "encoding/json"

// MultiPolygon is a set of Polygons.
type MultiPolygon [][][]Position

// NewMultiPolygon returns a new MultiPolygon from the supplied polygons.
func NewMultiPolygon(p ...[][]Position) *Feature {
	return &Feature{
		Geometry: (*MultiPolygon)(&p),
	}
}

// MarshalJSON returns the JSON encoding of the MultiPolygon.
func (m *MultiPolygon) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	return json.Marshal([][][]Position(*m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiPolygon) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[][][]Position)(m)); err != nil {
		return err
	}
	return m.validate()
}

func (m *MultiPolygon) validate() error {
	for _, p := range *m {
		for _, r := range p {
			if len(r) < 4 {
				return errLinearRingTooShort
			}
		}
	}
	return nil
}
