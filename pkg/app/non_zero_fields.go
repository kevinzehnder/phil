package app

import (
	"fmt"
	"reflect"
)

// returns all the fieldnames of a struct, that have nonZero values
func GetNonZeroFields(item interface{}) ([]string, error) {
	var fields []string

	// Get the type and value of the item
	typ := reflect.TypeOf(item)
	val := reflect.ValueOf(item)

	// Make sure the item is a pointer to a struct
	if typ.Kind() != reflect.Ptr || val.IsNil() {
		return fields, fmt.Errorf("item must be a non-nil pointer to a struct")
	}
	elemType := typ.Elem()
	if elemType.Kind() != reflect.Struct {
		return fields, fmt.Errorf("item must be a pointer to a struct")
	}

	// Iterate over the fields of the struct and check if they are non-zero
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		value := val.Elem().Field(i)
		if !value.IsZero() {
			fields = append(fields, field.Name)
		}
	}

	return fields, nil
}
