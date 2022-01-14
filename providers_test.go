package randomiser_test

import (
	"reflect"
	"testing"

	randomise "github.com/amwill04/randomiser"
)

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
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Logf("\tTest: When %s", name)

			as := randomise.As(tt.expected)

			actual := reflect.New(reflect.TypeOf(tt.expected)).Elem()

			err := as(actual, reflect.TypeOf(tt.expected), name)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tt.expected, actual.Interface()) {
				t.Errorf("actual: %v does not equal expected: %v", actual, tt.expected)
			}

		})
	}
}
