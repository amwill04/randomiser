package randomise

import (
	"math/rand"
	"reflect"
)

type Provider func(reflect.Value, reflect.Type, string) error

func OneOf(values ...interface{}) Provider {
	return func(value reflect.Value, typ reflect.Type, fieldName string) error {
		v := values[rand.Int63n(int64(len(values)))]
		var newValue reflect.Value
		newValue = reflect.ValueOf(v)
		if newValue.Type() != typ {
			return MalformedProviderType{
				typeRequired: typ,
				value:        v,
				fieldName:    fieldName,
			}
		}
		value.Set(newValue)
		return nil
	}
}

func As(v interface{}) Provider {
	return func(value reflect.Value, typ reflect.Type, fieldName string) error {
		newValue := reflect.ValueOf(v)
		if newValue.Type() != typ {
			return MalformedProviderType{
				typeRequired: typ,
				value:        v,
				fieldName:    fieldName,
			}
		}
		value.Set(newValue)
		return nil
	}
}
