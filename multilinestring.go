package geojson

import "encoding/json"

// MultiLineString is a set of LineStrings.
type MultiLineString [][]Position

// MarshalJSON returns the JSON encoding of the MultiLineString.
func (m *MultiLineString) MarshalJSON() ([]byte, error) {
	for _, ls := range *m {
		if len(ls) < 2 {
			return nil, errLineStringTooShort
		}
	}
	return json.Marshal([][]Position(*m))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (m *MultiLineString) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[][]Position)(m)); err != nil {
		return err
	}

	for _, ls := range *m {
		if len(ls) < 2 {
			return errLineStringTooShort
		}
	}
	return nil
}
