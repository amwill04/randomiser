package randomise

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"sync/atomic"
	"time"
)

var (
	ignoreRegex = regexp.MustCompile("(\\binterface\\b)|\\bchan\\b")
)

func NewRandomise() Random {
	configs := make(map[string]Config)
	defaultConfig := Config{
		Provider:     nil,
		MapKeyLength: 3,
		StringLength: 1,
		SliceLength:  1,
		MapLength:    1,
	}
	return Random{
		seed:           nil,
		configs:        configs,
		currentConfig:  nil,
		defaultConfig:  &defaultConfig,
		createdWithNew: true,
	}
}

type Config struct {
	Provider     Provider
	MapKeyLength int
	StringLength int
	SliceLength  int
	MapLength    int
}

type Random struct {
	seed           *seed
	configs        map[string]Config
	currentConfig  *Config
	defaultConfig  *Config
	createdWithNew bool
}

func (r *Random) SetStringLength(length int) {
	r.defaultConfig.StringLength = length
}

func (r *Random) SetSliceLength(length int) {
	r.defaultConfig.SliceLength = length
}

func (r *Random) SetMapLength(length int) {
	r.defaultConfig.SliceLength = length
}

func (r *Random) SetMapKeyLength(length int) {
	r.defaultConfig.MapKeyLength = length
}

func (r *Random) SetSeed(s int64) {
	rand.Seed(s)
	newSeed := seed(s)
	r.seed = &newSeed
}

func (r *Random) AddTypeConfig(name string, config Config) {
	r.configs[name] = config
}

type seed int64

func (s *seed) nextInt() int64 {
	return atomic.AddInt64((*int64)(s), 1)
}

func (r *Random) Struct(dst interface{}) error {
	if !r.createdWithNew {
		return MalformedRandom{reason: reasonMalformedNoCreatedWithNew}
	}
	if r.seed == nil {
		panic(MalformedRandom{reason: reasonMalformedSeedNotSet})
	}
	value := reflect.ValueOf(dst)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("struct should be a pointer, been given a non-pointer %T", dst)
	}

	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("element should be a struct, given a non-struct: %T", dst)
	}

	typ := value.Type()
	nFields := value.NumField()
	for i := 0; i < nFields; i++ {
		fieldVal := value.Field(i)
		fieldTyp := typ.Field(i)
		r.currentConfig = r.defaultConfig
		if config, ok := r.configs[fieldTyp.Name]; ok {
			if config.Provider != nil {
				if err := config.Provider(fieldVal, fieldTyp.Type, fieldTyp.Name); err != nil {
					return err
				}
				continue
			}
			r.currentConfig = &config
		}
		if err := r.randomiseField(fieldVal, fieldTyp.Type, nil); err != nil {
			return err
		}
	}
	return nil
}

func (r Random) randomiseField(value reflect.Value, typ reflect.Type, length *int) error {
	switch value.Interface().(type) {
	case int8, *int8:
		r.randomiseInt8(value, typ)
	case int16, *int16:
		r.randomiseInt16(value, typ)
	case int32, *int32:
		r.randomiseInt32(value, typ)
	case int64, *int64:
		r.randomiseInt64(value, typ)
	case int, *int:
		r.randomiseInt(value, typ)
	case float32, *float32:
		r.randomiseFloat32(value, typ)
	case float64, *float64:
		r.randomiseFloat64(value, typ)
	case uint8, *uint8:
		r.randomiseUInt8(value, typ)
	case uint16, *uint16:
		r.randomiseUInt16(value, typ)
	case uint32, *uint32:
		r.randomiseUInt32(value, typ)
	case uint64, *uint64:
		r.randomiseUInt64(value, typ)
	case uint, *uint:
		r.randomiseUInt(value, typ)
	case bool, *bool:
		randomiseBool(value, typ)
	case string, *string:
		l := r.currentConfig.StringLength
		if length != nil {
			l = *length
		}
		r.randomiseString(value, typ, l)
	case time.Time, *time.Time:
		_ = r.randomiseTime(value, typ)
	default:
		kind := value.Kind()
		if value.Kind() == reflect.Ptr {
			kind = typ.Elem().Kind()
		}
		if ignoreRegex.MatchString(typ.String()) {
			return nil
		}
		if err := r.randomiseCustomField(value, kind, typ, length); err != nil {
			return err
		}
	}
	return nil
}

func (r Random) randomiseCustomField(value reflect.Value, kind reflect.Kind, typ reflect.Type, length *int) error {
	switch kind {
	case reflect.Int8:
		r.randomiseInt8(value, typ)
	case reflect.Int16:
		r.randomiseInt16(value, typ)
	case reflect.Int32:
		r.randomiseInt32(value, typ)
	case reflect.Int64:
		r.randomiseInt64(value, typ)
	case reflect.Int:
		r.randomiseInt(value, typ)
	case reflect.Float32:
		r.randomiseFloat32(value, typ)
	case reflect.Float64:
		r.randomiseFloat64(value, typ)
	case reflect.Uint8:
		r.randomiseUInt8(value, typ)
	case reflect.Uint16:
		r.randomiseUInt16(value, typ)
	case reflect.Uint32:
		r.randomiseUInt32(value, typ)
	case reflect.Uint64:
		r.randomiseUInt64(value, typ)
	case reflect.Uint:
		r.randomiseUInt(value, typ)
	case reflect.Bool:
		randomiseBool(value, typ)
	case reflect.String:
		l := r.currentConfig.StringLength
		if length != nil {
			l = *length
		}
		r.randomiseString(value, typ, l)
	case reflect.Struct:
		err := r.randomiseTime(value, typ)
		if err != nil {
			v := reflect.New(typ)
			if typ.Kind() == reflect.Ptr {
				v = reflect.New(typ.Elem())
			}
			err = r.Struct(v.Interface())
			if err != nil {
				return err
			}
			if typ.Kind() != reflect.Ptr {
				v = v.Elem()
			}
			value.Set(v)
		}
	case reflect.Slice:
		return r.randomiseSlice(value, typ)
	case reflect.Map:
		return r.randomiseMap(value, typ)
	case reflect.Array:
		return r.randomiseArray(value, typ)
	default:
		return UnknownType{Type: value.Type()}
	}
	return nil
}

func (r Random) randomInt8() int8 {
	return int8(r.seed.nextInt() % math.MaxInt8)
}

func (r Random) randomiseInt8(value reflect.Value, typ reflect.Type) {
	v := r.randomInt8()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomInt16() int16 {
	return int16(r.seed.nextInt() % math.MaxInt16)
}

func (r Random) randomiseInt16(value reflect.Value, typ reflect.Type) {
	v := r.randomInt16()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomInt32() int32 {
	return int32(r.seed.nextInt() % math.MaxInt32)
}

func (r Random) randomiseInt32(value reflect.Value, typ reflect.Type) {
	v := r.randomInt32()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomInt64() int64 {
	return r.seed.nextInt() % math.MaxInt64
}

func (r Random) randomiseInt64(value reflect.Value, typ reflect.Type) {
	v := r.randomInt64()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomInt() int {
	return int(r.seed.nextInt())
}

func (r Random) randomiseInt(value reflect.Value, typ reflect.Type) {
	v := r.randomInt()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomFloat32() float32 {
	return float32(r.seed.nextInt()%10)/10.0 + float32(r.seed.nextInt()%10)
}

func (r Random) randomiseFloat32(value reflect.Value, typ reflect.Type) {
	v := r.randomFloat32()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomFloat64() float64 {
	return float64(r.seed.nextInt()%10)/10.0 + float64(r.seed.nextInt()%10)
}

func (r Random) randomiseFloat64(value reflect.Value, typ reflect.Type) {
	v := r.randomFloat64()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomUInt8() uint8 {
	return uint8(r.seed.nextInt() % math.MaxUint8)
}

func (r Random) randomiseUInt8(value reflect.Value, typ reflect.Type) {
	v := r.randomUInt8()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomUInt16() uint16 {
	return uint16(r.seed.nextInt() % math.MaxUint16)
}

func (r Random) randomiseUInt16(value reflect.Value, typ reflect.Type) {
	v := r.randomUInt16()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomUInt32() uint32 {
	return uint32(r.seed.nextInt() % math.MaxUint32)
}

func (r Random) randomiseUInt32(value reflect.Value, typ reflect.Type) {
	v := r.randomUInt32()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomUInt64() uint64 {
	return uint64(r.seed.nextInt())
}

func (r Random) randomiseUInt64(value reflect.Value, typ reflect.Type) {
	v := r.randomUInt64()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomUInt() uint {
	return uint(r.seed.nextInt())
}

func (r Random) randomiseUInt(value reflect.Value, typ reflect.Type) {
	v := r.randomUInt()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomBool() bool {
	return rand.Intn(2) == 1
}

func randomiseBool(value reflect.Value, typ reflect.Type) {
	v := rand.Intn(2) == 1
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

const alphabetAll = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (r Random) randomString(length int) string {
	str := bytes.Buffer{}
	str.Grow(length)
	for i := 0; i < length; i++ {
		str.WriteByte(alphabetAll[r.seed.nextInt()%int64(len(alphabetAll))])
	}
	return str.String()
}

func (r Random) randomiseString(value reflect.Value, typ reflect.Type, length int) {
	v := r.randomString(length)
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
}

func (r Random) randomTime() time.Time {
	return time.Date(
		int(1972+r.seed.nextInt()%60),
		time.Month(1+(r.seed.nextInt()%12)),
		int(1+(r.seed.nextInt()%25)),
		0,
		0,
		0,
		0,
		time.UTC,
	)
}

func (r Random) randomiseTime(value reflect.Value, typ reflect.Type) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = internalNotTime{}
		}
	}()
	v := r.randomTime()
	var newValue reflect.Value
	if typ.Kind() == reflect.Ptr {
		newType := reflect.New(typ.Elem())
		newValue = reflect.ValueOf(&v).Convert(newType.Type())
	} else {
		newType := reflect.New(typ).Elem()
		newValue = reflect.ValueOf(v).Convert(newType.Type())
	}
	value.Set(newValue)
	return
}

func (r Random) randomiseSlice(value reflect.Value, typ reflect.Type) error {
	baseType := typ
	if typ.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}
	length := r.currentConfig.SliceLength
	newSlice := reflect.MakeSlice(baseType, length, length)
	for i := 0; i < length; i++ {
		v := reflect.New(baseType.Elem()).Elem()
		if err := r.randomiseField(v, v.Type(), nil); err != nil {
			return err
		}
		newSlice.Index(i).Set(v)
	}
	if typ.Kind() == reflect.Ptr {
		newPtrSlice := reflect.New(typ.Elem())
		newPtrSlice.Elem().Set(newSlice)
		newSlice = newPtrSlice
	}
	value.Set(newSlice)
	return nil
}

func (r Random) randomiseMap(value reflect.Value, typ reflect.Type) error {
	baseType := typ
	if typ.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}
	mapType := reflect.MapOf(baseType.Key(), baseType.Elem())
	newMap := reflect.MakeMap(mapType)
	for i := 0; i < r.currentConfig.MapLength; i++ {
		k := reflect.New(baseType.Key()).Elem()
		v := reflect.New(baseType.Elem()).Elem()
		if err := r.randomiseField(k, k.Type(), &r.currentConfig.MapKeyLength); err != nil {
			return err
		}
		if err := r.randomiseField(v, v.Type(), nil); err != nil {
			return err
		}
		newMap.SetMapIndex(k, v)
	}
	if typ.Kind() == reflect.Ptr {
		newPtrMap := reflect.New(typ.Elem())
		newPtrMap.Elem().Set(newMap)
		newMap = newPtrMap
	}
	value.Set(newMap)
	return nil
}

func (r *Random) randomiseArray(value reflect.Value, typ reflect.Type) error {
	baseType := typ
	if typ.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}
	newArray := reflect.New(baseType)
	for i := 0; i < newArray.Elem().Len(); i++ {
		v := newArray.Elem().Index(i)
		if err := r.randomiseField(v, v.Type(), nil); err != nil {
			return err
		}
		newArray.Elem().Index(i).Set(v)
	}
	if typ.Kind() != reflect.Ptr {
		newArray = newArray.Elem()
	}
	value.Set(newArray)
	return nil
}
