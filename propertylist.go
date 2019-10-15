package geojson

import (
	"encoding/json"
	"fmt"
)

// PropertyList is a list of Properties.
type PropertyList []Property

// NewPropertyList creates a new PropertyList from the supplied Properties.
func NewPropertyList(props ...Property) PropertyList {
	return PropertyList(props)
}

// MarshalJSON returns the JSON encoding of the PropertyList.
func (l *PropertyList) MarshalJSON() ([]byte, error) {
	props := properties{}
	for _, p := range *l {
		props[p.Name] = p.Value
	}
	return json.Marshal(&props)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (l *PropertyList) UnmarshalJSON(data []byte) error {
	props := properties{}
	if err := json.Unmarshal(data, &props); err != nil {
		return err
	}

	list := make([]Property, len(props))
	i := 0
	for k, v := range props {
		list[i] = Property{
			Name:  k,
			Value: v,
		}
		i++
	}

	*l = list
	return nil
}

// Get a Property from the list.
func (l *PropertyList) Get(name string) (*Property, bool) {
	for _, p := range *l {
		if p.Name == name {
			return &p, true
		}
	}
	return nil, false
}

// GetValue assigns a named property to dest if the types are equal.
func (l *PropertyList) GetValue(name string, dest interface{}) error {
	for _, p := range *l {
		if p.Name == name {
			return p.GetValue(dest)
		}
	}
	return fmt.Errorf("property '%s' doesn't exist", name)
}

type properties map[string]interface{}
