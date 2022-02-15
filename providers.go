package randomiser

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Provider func(*rand.Rand, reflect.Value, reflect.Type, string) error

func OneOf(values ...interface{}) Provider {
	return func(r *rand.Rand, value reflect.Value, typ reflect.Type, fieldName string) error {
		baseType := typ
		var isPtr bool
		if baseType.Kind() == reflect.Ptr {
			isPtr = true
			baseType = baseType.Elem()
		}
		newValue := reflect.New(baseType)
		v := values[r.Int63n(int64(len(values)))]
		setValue := reflect.ValueOf(v)
		if !setValue.CanConvert(newValue.Elem().Type()) {
			return MalformedProviderType{
				typeRequired: typ,
				value:        v,
				fieldName:    fieldName,
			}
		}
		setValue = setValue.Convert(newValue.Elem().Type())
		newValue.Elem().Set(setValue)
		if isPtr {
			value.Set(newValue)
		} else {
			value.Set(newValue.Elem())
		}
		return nil
	}
}

func As(v interface{}) Provider {
	return func(_ *rand.Rand, value reflect.Value, typ reflect.Type, fieldName string) error {
		if v == nil {
			return nil
		}
		baseType := typ
		setValue := reflect.ValueOf(v)
		var isPtr bool
		if baseType.Kind() == reflect.Ptr {
			isPtr = true
			baseType = baseType.Elem()
		}
		if setValue.Kind() == reflect.Ptr {
			setValue = setValue.Elem()
		}
		newValue := reflect.New(baseType)
		if !setValue.CanConvert(newValue.Elem().Type()) {
			return MalformedProviderType{
				typeRequired: typ,
				value:        v,
				fieldName:    fieldName,
			}
		}
		setValue = setValue.Convert(newValue.Elem().Type())
		newValue.Elem().Set(setValue)
		if isPtr {
			value.Set(newValue)
		} else {
			value.Set(newValue.Elem())
		}
		return nil
	}
}

var (
	timeType = reflect.TypeOf(time.Time{})
)

func Between(start interface{}, end interface{}) Provider {
	return func(r *rand.Rand, value reflect.Value, typ reflect.Type, fieldName string) error {
		vStart := reflect.ValueOf(start)
		if vStart.Type() != typ {
			return MalformedProviderType{
				typeRequired: typ,
				value:        start,
				fieldName:    fieldName,
			}
		}
		vEnd := reflect.ValueOf(end)
		if vEnd.Type() != typ {
			return MalformedProviderType{
				typeRequired: typ,
				value:        start,
				fieldName:    fieldName,
			}
		}
		isPtr := false
		t := typ
		if t.Kind() == reflect.Ptr {
			isPtr = true
			vStart = vStart.Elem()
			vEnd = vEnd.Elem()
			t = typ.Elem()
		}
		newValue := reflect.New(t)
		var setValue reflect.Value
		switch t.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			if vEnd.Int() < vStart.Int() {
				return MalformedProviderArgument{
					message: fmt.Sprintf("randomise: provided end '%v' is less than start '%v'", end, start),
				}
			}
			span := vEnd.Int() - vStart.Int()
			setValue = reflect.ValueOf(r.Int63n(span) + vStart.Int()).Convert(t)
		default:
			isTime := typ.ConvertibleTo(reflect.TypeOf(time.Time{})) || typ.ConvertibleTo(reflect.TypeOf(&time.Time{}))
			if !isTime {
				return MalformedProviderUnsupportedType{
					typeProvided: typ,
					fieldName:    fieldName,
				}
			}
			timeStart := vStart.Convert(timeType).Interface().(time.Time)
			timeEnd := vEnd.Convert(timeType).Interface().(time.Time)
			if timeEnd.Before(timeStart) {
				return MalformedProviderArgument{
					message: fmt.Sprintf("randomise: provided end '%v' is less than start '%v'", end, start),
				}
			}
			span := timeEnd.Unix() - timeStart.Unix()
			newTime := time.Unix(r.Int63n(span)+timeStart.Unix(), 0).UTC()
			setValue = reflect.ValueOf(newTime).Convert(t)
		}
		newValue.Elem().Set(setValue)
		if isPtr {
			value.Set(newValue)
		} else {
			value.Set(newValue.Elem())
		}
		return nil
	}
}
