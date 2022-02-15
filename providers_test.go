package randomiser_test

import (
	"math/rand"
	"reflect"
	"testing"

	randomise "github.com/amwill04/randomiser"
)

type customInt int

func TestAs(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
	}{
		"string": {
			expected: "string",
		},
		"int": {
			expected: 100,
		},
		"pointer": {
			expected: stringPtr("string pointer"),
		},
		"custom": {
			expected: customInt(100),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Logf("\tTest: When %s", name)

			as := randomise.As(tt.expected)

			actual := reflect.New(reflect.TypeOf(tt.expected)).Elem()

			src := rand.NewSource(10)
			r := rand.New(src)

			err := as(r, actual, reflect.TypeOf(tt.expected), name)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tt.expected, actual.Interface()) {
				t.Errorf("actual: %v does not equal expected: %v", actual, tt.expected)
			}

		})
	}
}

func TestAs_CustomType(t *testing.T) {
	t.Log("\tTest: When custom type")

	expected := customInt(100)

	as := randomise.As(expected)

	actual := reflect.New(reflect.TypeOf(0)).Elem()

	src := rand.NewSource(10)
	r := rand.New(src)

	err := as(r, actual, reflect.TypeOf(0), "Int")

	if err != nil {
		t.Error(err)
	}

	if int64(expected) != actual.Int() {
		t.Errorf("actual: %v does not equal expected: %v", actual, expected)
	}

}
