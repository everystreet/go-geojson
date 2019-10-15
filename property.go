package geojson

import (
	"fmt"
	"reflect"
)

// Property represents a single property of arbitrary type.
type Property struct {
	Name  string
	Value interface{}
}

// GetValue assigns the value to dest if the types are equal.
func (p *Property) GetValue(dest interface{}) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be pointer")
	}

	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				err = fmt.Errorf("type error: %v", r)
			}
		}()
		reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(p.Value))
	}()
	return err
}
