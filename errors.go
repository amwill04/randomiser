package randomise

import (
	"fmt"
	"reflect"
)

type UnknownType struct {
	Type reflect.Type
}

func (t UnknownType) Error() string {
	return fmt.Sprintf("randomise: unknown type provided '%v'. Dont know how to fake", t.Type)
}

type reasonMalformed string

const (
	reasonMalformedNoCreatedWithNew reasonMalformed = "random was not created using NewRandomise()"
)

type MalformedRandom struct {
	reason reasonMalformed
}

func (t MalformedRandom) Error() string {
	return fmt.Sprintf("randomise: %s", t.reason)
}

type internalNotTime struct {
}

func (receiver internalNotTime) Error() string {
	return fmt.Sprint("randomise: struct was not time.Time")
}

type MalformedProviderType struct {
	typeRequired reflect.Type
	value        interface{}
	fieldName    string
}

func (t MalformedProviderType) Error() string {
	return fmt.Sprintf("randomise: field '%s' requires %v but provded value '%s' is %[3]T", t.fieldName, t.typeRequired, t.value)
}
