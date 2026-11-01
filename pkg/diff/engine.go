package diff

import (
	"errors"
	"fmt"
	"github.com/bradyhelms/netsink/pkg/config"
	"reflect"
)

type DriftItem struct {
	Field    string `json:"field"`
	Desired  any    `json:"desired"`
	Observed any    `json:"observed"`
}

func iterateFields(v any, prefix string) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return
		}
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := prefix + fieldType.Name

		if fieldVal.Kind() == reflect.Struct {
			fmt.Printf("Nested Struct: %s\n", fieldName)
			iterateFields(fieldVal.Interface(), fieldName+".")
		} else {
			fmt.Printf("%s = %v (Type: %s)\n", fieldName, fieldVal.Interface(), fieldType.Type)
		}
	}
}

// Compare two structs and return a slice of DriftItems
func compareStructs(
	desiredState config.NetworkState,
	observedState config.NetworkState,
) ([]DriftItem, error) {

	var item []DriftItem
	var err error

	if err != nil {
		return nil, err
	}

	return item, nil
}
