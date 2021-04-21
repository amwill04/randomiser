## randomiser

Generate and fill structs with fake data based on types for testing purposes

Will generate all values for all public fields of any type, including custom types `eg type Foo string`. The only 
type it will not generate data for will be anything that includes an `interface{}` - `e.g. map[string]interface{}` 
will not be generated as we have no knowledge to what the interface should be.

Difference with this and [faker](https://github.com/bxcodec/faker) is this module will not generate specific types 
like `IPV4` ip address. It deals with simply generating fields on their base type. It comes with `providers` to 
allow this function however.

### download

```shell
$ go get -u github.com/amwill04/randomiser
```


### simple usage

```go
package myPkg_test

import (
	"github.com/amwill04/randomiser"
	"testing"
	"time"
)



type Foo struct {
	FieldA string
	FieldB int16
	FieldC *time.Time
}

func MyTest(t *testing.T) {
	foo := Foo{}
	r := randomise.NewRandomise(time.Now().UnixNano())
	err := r.Struct(&foo)
	if err != nil {
		t.Fatal(err)
    }
    
    // can now use generated fields
}
```

### providers

Comes with the `providers` that allow you to narrow down the generated field or specify it. This allows you to use 
3rd party modules to generate phone numbers, email address etc.

- `As("use_this_value")`
- `Between(1, 10)` // only accepts `time.Time`, `[u]int[8/26/32/64]` types and any custom types of these
- `OneOf("option_a", "option_b", "option_c")`

All arguments to the provider must be of the field type, or it will error.

You can also generate your own `Provider`

```go
package myPkg_test

import (
	"github.com/amwill04/randomiser"
	"testing"
	"time"
)

type EnumType string

const (
	OptionA = "option_a"
	OptionB = "option_b"
	OptionC = "option_c"
)

type PhoneNumber string
type DateOfBirth time.Time

type Foo struct {
	FieldA EnumType
	FieldB PhoneNumber
	FieldC DateOfBirth
}

func MyTest(t *testing.T) {
	foo := Foo{}
	r := randomise.NewRandomise(time.Now().UnixNano())
	r.AddTypeConfig("FieldA", randomise.WithProvider(randomise.OneOf(OptionA, OptionB, OptionC)))
	r.AddTypeConfig("FieldB", randomise.WithProvider(randomise.As("07934536665")))
	r.AddTypeConfig("FieldC", randomise.WithProvider(randomise.Between(DateOfBirth(time.Date(2000, 1,1,0,0,0,0,time.UTC)), DateOfBirth(time.Now()))))
	err := r.Struct(&foo)
	if err != nil {
		t.Fatal(err)
    }
    
    // can now use generated fields
}
```

### generating specific types

If you wish to generate random values of a specific kind - e.g. email address

```go
package myPkg_test

import (
	"github.com/amwill04/randomiser"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)



type Foo struct {
	FieldEmail string
}

func MyTest(t *testing.T) {
	foo := Foo{}
	r := randomise.NewRandomise(time.Now().UnixNano())
	r.AddTypeConfig("FieldEmail", randomise.WithProvider(randomise.As(faker.Internet().Email())))
	err := r.Struct(&foo)
	if err != nil {
		t.Fatal(err)
    }
    
    // can now use generated fields
}
```

## benchmark

### randomiser
```bash
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRandom_Generate-12       574525              2043 ns/op             896 B/op         50 allocs/op
```

### Faker - without tags
```bash
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkFakerDataNOTTagged-12            273898              4268 ns/op            1368 B/op         63 allocs/op
```
