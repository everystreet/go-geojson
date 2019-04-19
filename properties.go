package geojson

import (
	"encoding/json"
)

// Property represents a single property of arbitrary type.
type Property struct {
	Name  string
	Value interface{}
}

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

// StringProp returns a new Property with the specified name and string value.
func StringProp(name, value string) Property {
	return Property{
		Name:  name,
		Value: value,
	}
}

type properties map[string]interface{}
