package randomiser

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
	return fmt.Sprintf("randomise: field '%s' requires %v but provded value '%v' is %[3]T", t.fieldName, t.typeRequired, t.value)
}

type MalformedProviderArgument struct {
	message string
}

func (t MalformedProviderArgument) Error() string {
	return t.message
}

type MalformedProviderUnsupportedType struct {
	typeProvided reflect.Type
	fieldName    string
}

func (t MalformedProviderUnsupportedType) Error() string {
	return fmt.Sprintf("randomise: field '%s' has unsupported type '%s' for provider", t.fieldName, t.typeProvided)
}
