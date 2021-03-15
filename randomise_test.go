package randomise_test

import (
	"testing"
	"time"

	randomise "github.com/amwill04/randomiser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CustomIntType int
type CustomInt8Type int8
type CustomInt16Type int16
type CustomInt32Type int32
type CustomInt64Type int64
type CustomFloat32Type float32
type CustomFloat64Type float64
type CustomUintType uint
type CustomUint8Type uint8
type CustomUint16Type uint16
type CustomUint32Type uint32
type CustomUint64Type uint64
type CustomBoolType bool
type CustomStringType string
type CustomTimeType time.Time
type CustomSliceStringType []string

func (t CustomTimeType) String() string {
	ts := time.Time(t)
	return ts.Format(time.RFC3339)
}

type StructA struct {
	ColStructB StructB
	ColNested  string
}

type StructB struct {
	StructC
}

type StructC struct {
	ColDeepNest int
}

type BaseTypes struct {
	ColInt                         int
	ColInt8                        int8
	ColInt16                       int16
	ColInt32                       int32
	ColInt64                       int64
	ColFloat32                     float32
	ColFloat64                     float64
	ColUint                        uint
	ColUint8                       uint8
	ColUint16                      uint16
	ColUint32                      uint32
	ColUint64                      uint64
	ColBool                        bool
	ColString                      string
	ColTime                        time.Time
	ColByte                        []byte
	ColSliceInt                    []int
	ColSliceInt8                   []int8
	ColSliceInt16                  []int16
	ColSliceInt32                  []int32
	ColSliceInt64                  []int64
	ColSliceFloat32                []float32
	ColSliceFloat64                []float64
	ColSliceUint                   []uint
	ColSliceUint8                  []uint8
	ColSliceUint16                 []uint16
	ColSliceUint32                 []uint32
	ColSliceUint64                 []uint64
	ColSliceBool                   []bool
	ColSliceString                 []string
	ColSliceTime                   []time.Time
	ColPtrByte                     []*byte
	ColSlicePtrInt                 []*int
	ColSlicePtrInt8                []*int8
	ColSlicePtrInt16               []*int16
	ColSlicePtrInt32               []*int32
	ColSlicePtrInt64               []*int64
	ColSlicePtrFloat32             []*float32
	ColSlicePtrFloat64             []*float64
	ColSlicePtrUint                []*uint
	ColSlicePtrUint8               []*uint8
	ColSlicePtrUint16              []*uint16
	ColSlicePtrUint32              []*uint32
	ColSlicePtrUint64              []*uint64
	ColSlicePtrBool                []*bool
	ColSlicePtrString              []*string
	ColSlicePtrTime                []*time.Time
	ColStruct                      StructA
	ColMapStringString             map[string]string
	ColMapStringPtrString          map[string]*string
	ColMapStringInt                map[string]int
	ColMapStringPtrInt             map[string]*int
	ColMapIntSliceString           map[int][]string
	ColMapIntPtrSliceString        map[int]*[]string
	ColMapIntSlicePtrString        map[int][]*string
	ColMapIntPtrSlicePtrString     map[int]*[]*string
	ColMapStringMapStringString    map[string]map[string]string
	ColMapStringPtrMapStringString map[string]*map[string]string
	ColSliceSliceString            [][]string
	ColSliceSlicePtrString         [][]*string
	ColInterface                   interface{}
	ColSliceInterface              []interface{}
	ColMapStringInterface          map[string]interface{}
	ColArrayString                 [2]string
	ColArrayArrayInt8Ptr           [2]*[2]*int8
	ColChanString                  chan string
	ColSliceChanString             []chan string
}

type BaseTypesPtr struct {
	ColInt                         *int
	ColInt8                        *int8
	ColInt16                       *int16
	ColInt32                       *int32
	ColInt64                       *int64
	ColFloat32                     *float32
	ColFloat64                     *float64
	ColUint                        *uint
	ColUint8                       *uint8
	ColUint16                      *uint16
	ColUint32                      *uint32
	ColUint64                      *uint64
	ColBool                        *bool
	ColString                      *string
	ColTime                        *time.Time
	ColByte                        *[]byte
	ColSliceInt                    *[]int
	ColSliceInt8                   *[]int8
	ColSliceInt16                  *[]int16
	ColSliceInt32                  *[]int32
	ColSliceInt64                  *[]int64
	ColSliceFloat32                *[]float32
	ColSliceFloat64                *[]float64
	ColSliceUint                   *[]uint
	ColSliceUint8                  *[]uint8
	ColSliceUint16                 *[]uint16
	ColSliceUint32                 *[]uint32
	ColSliceUint64                 *[]uint64
	ColSliceBool                   *[]bool
	ColSliceString                 *[]string
	ColSliceTime                   *[]time.Time
	ColPtrByte                     *[]*byte
	ColSlicePtrInt                 *[]*int
	ColSlicePtrInt8                *[]*int8
	ColSlicePtrInt16               *[]*int16
	ColSlicePtrInt32               *[]*int32
	ColSlicePtrInt64               *[]*int64
	ColSlicePtrFloat32             *[]*float32
	ColSlicePtrFloat64             *[]*float64
	ColSlicePtrUint                *[]*uint
	ColSlicePtrUint8               *[]*uint8
	ColSlicePtrUint16              *[]*uint16
	ColSlicePtrUint32              *[]*uint32
	ColSlicePtrUint64              *[]*uint64
	ColSlicePtrBool                *[]*bool
	ColSlicePtrString              *[]*string
	ColSlicePtrTime                *[]*time.Time
	ColStruct                      *StructA
	ColMapStringString             *map[string]string
	ColMapStringPtrString          *map[string]*string
	ColMapStringInt                *map[string]int
	ColMapStringPtrInt             *map[string]*int
	ColMapIntSliceString           *map[int][]string
	ColMapIntPtrSliceString        *map[int]*[]string
	ColMapIntSlicePtrString        *map[int][]*string
	ColMapIntPtrSlicePtrString     *map[int]*[]*string
	ColMapStringMapStringString    *map[string]map[string]string
	ColMapStringPtrMapStringString *map[string]*map[string]string
	ColSliceSliceString            *[][]string
	ColSliceSlicePtrString         *[][]*string
	ColInterface                   *interface{}
	ColSliceInterface              *[]interface{}
	ColMapStringInterface          *map[string]interface{}
	ColArrayString                 *[2]string
	ColArrayArrayInt8Ptr           *[2]*[2]*int8
	ColChanString                  *chan string
	ColSliceChanString             *[]chan string
}

type CustomTypes struct {
	ColCustomIntType         CustomIntType
	ColCustomInt8Type        CustomInt8Type
	ColCustomInt16Type       CustomInt16Type
	ColCustomInt32Type       CustomInt32Type
	ColCustomInt64Type       CustomInt64Type
	ColCustomFloat32Type     CustomFloat32Type
	ColCustomFloat64Type     CustomFloat64Type
	ColCustomUintType        CustomUintType
	ColCustomUint8Type       CustomUint8Type
	ColCustomUint16Type      CustomUint16Type
	ColCustomUint32Type      CustomUint32Type
	ColCustomUint64Type      CustomUint64Type
	ColCustomBoolType        CustomBoolType
	ColCustomStringType      CustomStringType
	ColCustomTimeType        CustomTimeType
	ColCustomSliceStringType CustomSliceStringType
	ColSliceCustomBool       []CustomBoolType
}

type CustomTypesPtr struct {
	ColCustomIntType         *CustomIntType
	ColCustomInt8Type        *CustomInt8Type
	ColCustomInt16Type       *CustomInt16Type
	ColCustomInt32Type       *CustomInt32Type
	ColCustomInt64Type       *CustomInt64Type
	ColCustomFloat32Type     *CustomFloat32Type
	ColCustomFloat64Type     *CustomFloat64Type
	ColCustomUintType        *CustomUintType
	ColCustomUint8Type       *CustomUint8Type
	ColCustomUint16Type      *CustomUint16Type
	ColCustomUint32Type      *CustomUint32Type
	ColCustomUint64Type      *CustomUint64Type
	ColCustomBoolType        *CustomBoolType
	ColCustomStringType      *CustomStringType
	ColCustomTimeType        *CustomTimeType
	ColCustomSliceStringType *CustomSliceStringType
	ColSliceCustomBool       *[]CustomBoolType
}

var _ = Describe("Randomise", func() {
	var (
		r randomise.Random

		mockDate = time.Date(1989, 4, 10, 0, 0, 0, 0, time.UTC)

		// returned values for seed
		mockColInt                     = 608169601
		mockColInt8            int8    = 3
		mockColInt16           int16   = 14083
		mockColInt32           int32   = 608169604
		mockColInt64           int64   = 608169605
		mockColFloat32         float32 = 7.599999904632568
		mockColFloat64                 = 9.8
		mockColUint            uint    = 608169610
		mockColUint8           uint8   = 221
		mockColUint16          uint16  = 4812
		mockColUint32          uint32  = 608169613
		mockColUint64          uint64  = 608169614
		mockColBool                    = true
		mockColString                  = "l"
		mockColTime                    = time.Date(1988, 6, 19, 0, 0, 0, 0, time.UTC)
		mockColByte                    = []byte{229}
		mockColSliceInt                = []int{608169620}
		mockColSliceInt8               = []int8{22}
		mockColSliceInt16              = []int16{14102}
		mockColSliceInt32              = []int32{608169623}
		mockColSliceInt64              = []int64{608169624}
		mockColSliceFloat32            = []float32{6.5}
		mockColSliceFloat64            = []float64{8.7}
		mockColSliceUint               = []uint{608169629}
		mockColSliceUint8              = []uint8{240}
		mockColSliceUint16             = []uint16{4831}
		mockColSliceUint32             = []uint32{608169632}
		mockColSliceUint64             = []uint64{608169633}
		mockColSliceBool               = []bool{true}
		mockColSliceString             = []string{"E"}
		mockColSliceTime               = []time.Time{time.Date(2007, 1, 13, 0, 0, 0, 0, time.UTC)}
		mockColSlicePtrByte            = []*byte{uint8Ptr(248)}
		mockColSlicePtrInt             = []*int{intPtr(608169639)}
		mockColSlicePtrInt8            = []*int8{int8Ptr(41)}
		mockColSlicePtrInt16           = []*int16{int16Ptr(14121)}
		mockColSlicePtrInt32           = []*int32{int32Ptr(608169642)}
		mockColSlicePtrInt64           = []*int64{int64Ptr(608169643)}
		mockColSlicePtrFloat32         = []*float32{float32Ptr(5.400000095367432)}
		mockColSlicePtrFloat64         = []*float64{float64Ptr(7.6)}
		mockColSlicePtrUint            = []*uint{uintPtr(608169648)}
		mockColSlicePtrUint8           = []*uint8{uint8Ptr(4)}
		mockColSlicePtrUint16          = []*uint16{uint16Ptr(4850)}
		mockColSlicePtrUint32          = []*uint32{uint32Ptr(608169651)}
		mockColSlicePtrUint64          = []*uint64{uint64Ptr(608169652)}
		mockColSlicePtrBool            = []*bool{boolPtr(true)}
		mockColSlicePtrString          = []*string{stringPtr("X")}
		mockColSlicePtrTime            = []*time.Time{timePtr(time.Date(2026, 8, 7, 0, 0, 0, 0, time.UTC))}
		mockColStruct                  = StructA{
			ColStructB: StructB{
				StructC: StructC{
					ColDeepNest: 608169666,
				},
			},
			ColNested: "b",
		}
		mockColMapStringString                               = map[string]string{"cde": "f"}
		mockColMapStringPtrString                            = map[string]*string{"ghi": stringPtr("j")}
		mockColMapStringInt                                  = map[string]int{"klm": 608169679}
		mockColMapStringPtrInt                               = map[string]*int{"opq": intPtr(608169683)}
		mockColMapIntSliceString                             = map[int][]string{608169684: {"t"}}
		mockColMapIntPtrSliceString                          = map[int]*[]string{608169686: {"v"}}
		mockColMapIntSlicePtrString                          = map[int][]*string{608169688: {stringPtr("x")}}
		mockColMapIntPtrSlicePtrString                       = map[int]*[]*string{608169690: {stringPtr("z")}}
		mockColMapStringMapStringString                      = map[string]map[string]string{"ABC": {"DEF": "G"}}
		mockColMapStringPtrMapStringString                   = map[string]*map[string]string{"HIJ": {"KLM": "N"}}
		mockColSliceSliceString                              = [][]string{{"O"}}
		mockColSliceSlicePtrString                           = [][]*string{{stringPtr("P")}}
		mockColArrayString                                   = [2]string{"Q", "R"}
		mockColArrayInt8ArrayPtr                             = [2]*[2]*int8{{int8Ptr(111), int8Ptr(112)}, {int8Ptr(113), int8Ptr(114)}}
		mockColCustomIntType               CustomIntType     = 608169601
		mockColCustomInt8Type              CustomInt8Type    = 3
		mockColCustomInt16Type             CustomInt16Type   = 14083
		mockColCustomInt32Type             CustomInt32Type   = 608169604
		mockColCustomInt64Type             CustomInt64Type   = 608169605
		mockColCustomFloat32Type           CustomFloat32Type = 7.599999904632568
		mockColCustomFloat64Type           CustomFloat64Type = 9.8
		mockColCustomUintType              CustomUintType    = 608169610
		mockColCustomUint8Type             CustomUint8Type   = 221
		mockColCustomUint16Type            CustomUint16Type  = 4812
		mockColCustomUint32Type            CustomUint32Type  = 608169613
		mockColCustomUint64Type            CustomUint64Type  = 608169614
		mockColCustomBoolType              CustomBoolType    = true
		mockColCustomStringType            CustomStringType  = "l"
		mockColCustomTimeType                                = CustomTimeType(time.Date(1988, 6, 19, 0, 0, 0, 0, time.UTC))
		mockColCustomSliceStringType                         = CustomSliceStringType{"p"}
		mockColSliceCustomBool                               = []CustomBoolType{true}
		//mockColOneOfString                                   = "one_of_a"
		//mockColOneOfSliceTime                                = []time.Time{time.Date(2017, 4, 10, 0, 0, 0, 0, time.UTC)}

		//typeProvider = func(value reflect.Value, typ reflect.Type, _ string) error {
		//	values := []TestTypeEnum{
		//		TestTypeEnumValueA,
		//		TestTypeEnumValueB,
		//	}
		//	v := values[rand.Intn(len(values))]
		//	var newValue reflect.Value
		//	if typ.Kind() == reflect.Ptr {
		//		newType := reflect.New(typ.Elem())
		//		newValue = reflect.ValueOf(&v).Convert(newType.Type())
		//	} else {
		//		newType := reflect.New(typ).Elem()
		//		newValue = reflect.ValueOf(v).Convert(newType.Type())
		//	}
		//	value.Set(newValue)
		//	return nil
		//}
	)

	BeforeEach(func() {
		r = randomise.NewRandomise()
		r.SetSeed(mockDate.Unix())
	})

	Context("when a struct is passed with base types", func() {
		It("should return randomised struct", func() {
			t := BaseTypes{}
			Expect(r.Struct(&t)).To(Succeed())
			Expect(t).To(Equal(BaseTypes{
				ColInt:                         mockColInt,
				ColInt8:                        mockColInt8,
				ColInt16:                       mockColInt16,
				ColInt32:                       mockColInt32,
				ColInt64:                       mockColInt64,
				ColFloat32:                     mockColFloat32,
				ColFloat64:                     mockColFloat64,
				ColUint:                        mockColUint,
				ColUint8:                       mockColUint8,
				ColUint16:                      mockColUint16,
				ColUint32:                      mockColUint32,
				ColUint64:                      mockColUint64,
				ColBool:                        mockColBool,
				ColString:                      mockColString,
				ColTime:                        mockColTime,
				ColByte:                        mockColByte,
				ColSliceInt:                    mockColSliceInt,
				ColSliceInt8:                   mockColSliceInt8,
				ColSliceInt16:                  mockColSliceInt16,
				ColSliceInt32:                  mockColSliceInt32,
				ColSliceInt64:                  mockColSliceInt64,
				ColSliceFloat32:                mockColSliceFloat32,
				ColSliceFloat64:                mockColSliceFloat64,
				ColSliceUint:                   mockColSliceUint,
				ColSliceUint8:                  mockColSliceUint8,
				ColSliceUint16:                 mockColSliceUint16,
				ColSliceUint32:                 mockColSliceUint32,
				ColSliceUint64:                 mockColSliceUint64,
				ColSliceBool:                   mockColSliceBool,
				ColSliceString:                 mockColSliceString,
				ColSliceTime:                   mockColSliceTime,
				ColPtrByte:                     mockColSlicePtrByte,
				ColSlicePtrInt:                 mockColSlicePtrInt,
				ColSlicePtrInt8:                mockColSlicePtrInt8,
				ColSlicePtrInt16:               mockColSlicePtrInt16,
				ColSlicePtrInt32:               mockColSlicePtrInt32,
				ColSlicePtrInt64:               mockColSlicePtrInt64,
				ColSlicePtrFloat32:             mockColSlicePtrFloat32,
				ColSlicePtrFloat64:             mockColSlicePtrFloat64,
				ColSlicePtrUint:                mockColSlicePtrUint,
				ColSlicePtrUint8:               mockColSlicePtrUint8,
				ColSlicePtrUint16:              mockColSlicePtrUint16,
				ColSlicePtrUint32:              mockColSlicePtrUint32,
				ColSlicePtrUint64:              mockColSlicePtrUint64,
				ColSlicePtrBool:                mockColSlicePtrBool,
				ColSlicePtrString:              mockColSlicePtrString,
				ColSlicePtrTime:                mockColSlicePtrTime,
				ColStruct:                      mockColStruct,
				ColMapStringString:             mockColMapStringString,
				ColMapStringPtrString:          mockColMapStringPtrString,
				ColMapStringInt:                mockColMapStringInt,
				ColMapStringPtrInt:             mockColMapStringPtrInt,
				ColMapIntSliceString:           mockColMapIntSliceString,
				ColMapIntPtrSliceString:        mockColMapIntPtrSliceString,
				ColMapIntSlicePtrString:        mockColMapIntSlicePtrString,
				ColMapIntPtrSlicePtrString:     mockColMapIntPtrSlicePtrString,
				ColMapStringMapStringString:    mockColMapStringMapStringString,
				ColMapStringPtrMapStringString: mockColMapStringPtrMapStringString,
				ColSliceSliceString:            mockColSliceSliceString,
				ColSliceSlicePtrString:         mockColSliceSlicePtrString,
				ColInterface:                   nil,
				ColSliceInterface:              nil,
				ColMapStringInterface:          nil,
				ColArrayString:                 mockColArrayString,
				ColArrayArrayInt8Ptr:           mockColArrayInt8ArrayPtr,
				ColChanString:                  nil,
				ColSliceChanString:             nil,
			}))
		})
	})

	Context("when a struct is passed with base types as pointers", func() {
		It("should return randomised struct", func() {
			t := BaseTypesPtr{}
			Expect(r.Struct(&t)).To(Succeed())
			Expect(t).To(Equal(BaseTypesPtr{
				ColInt:                         &mockColInt,
				ColInt8:                        &mockColInt8,
				ColInt16:                       &mockColInt16,
				ColInt32:                       &mockColInt32,
				ColInt64:                       &mockColInt64,
				ColFloat32:                     &mockColFloat32,
				ColFloat64:                     &mockColFloat64,
				ColUint:                        &mockColUint,
				ColUint8:                       &mockColUint8,
				ColUint16:                      &mockColUint16,
				ColUint32:                      &mockColUint32,
				ColUint64:                      &mockColUint64,
				ColBool:                        &mockColBool,
				ColString:                      &mockColString,
				ColTime:                        &mockColTime,
				ColByte:                        &mockColByte,
				ColSliceInt:                    &mockColSliceInt,
				ColSliceInt8:                   &mockColSliceInt8,
				ColSliceInt16:                  &mockColSliceInt16,
				ColSliceInt32:                  &mockColSliceInt32,
				ColSliceInt64:                  &mockColSliceInt64,
				ColSliceFloat32:                &mockColSliceFloat32,
				ColSliceFloat64:                &mockColSliceFloat64,
				ColSliceUint:                   &mockColSliceUint,
				ColSliceUint8:                  &mockColSliceUint8,
				ColSliceUint16:                 &mockColSliceUint16,
				ColSliceUint32:                 &mockColSliceUint32,
				ColSliceUint64:                 &mockColSliceUint64,
				ColSliceBool:                   &mockColSliceBool,
				ColSliceString:                 &mockColSliceString,
				ColSliceTime:                   &mockColSliceTime,
				ColPtrByte:                     &mockColSlicePtrByte,
				ColSlicePtrInt:                 &mockColSlicePtrInt,
				ColSlicePtrInt8:                &mockColSlicePtrInt8,
				ColSlicePtrInt16:               &mockColSlicePtrInt16,
				ColSlicePtrInt32:               &mockColSlicePtrInt32,
				ColSlicePtrInt64:               &mockColSlicePtrInt64,
				ColSlicePtrFloat32:             &mockColSlicePtrFloat32,
				ColSlicePtrFloat64:             &mockColSlicePtrFloat64,
				ColSlicePtrUint:                &mockColSlicePtrUint,
				ColSlicePtrUint8:               &mockColSlicePtrUint8,
				ColSlicePtrUint16:              &mockColSlicePtrUint16,
				ColSlicePtrUint32:              &mockColSlicePtrUint32,
				ColSlicePtrUint64:              &mockColSlicePtrUint64,
				ColSlicePtrBool:                &mockColSlicePtrBool,
				ColSlicePtrString:              &mockColSlicePtrString,
				ColSlicePtrTime:                &mockColSlicePtrTime,
				ColStruct:                      &mockColStruct,
				ColMapStringString:             &mockColMapStringString,
				ColMapStringPtrString:          &mockColMapStringPtrString,
				ColMapStringInt:                &mockColMapStringInt,
				ColMapStringPtrInt:             &mockColMapStringPtrInt,
				ColMapIntSliceString:           &mockColMapIntSliceString,
				ColMapIntPtrSliceString:        &mockColMapIntPtrSliceString,
				ColMapIntSlicePtrString:        &mockColMapIntSlicePtrString,
				ColMapIntPtrSlicePtrString:     &mockColMapIntPtrSlicePtrString,
				ColMapStringMapStringString:    &mockColMapStringMapStringString,
				ColMapStringPtrMapStringString: &mockColMapStringPtrMapStringString,
				ColSliceSliceString:            &mockColSliceSliceString,
				ColSliceSlicePtrString:         &mockColSliceSlicePtrString,
				ColInterface:                   nil,
				ColSliceInterface:              nil,
				ColMapStringInterface:          nil,
				ColArrayString:                 &mockColArrayString,
				ColArrayArrayInt8Ptr:           &mockColArrayInt8ArrayPtr,
				ColChanString:                  nil,
				ColSliceChanString:             nil,
			}))
		})
	})

	Describe("custom types", func() {
		Context("non-pointers", func() {
			It("should return randomised struct", func() {
				t := CustomTypes{}
				Expect(r.Struct(&t)).To(Succeed())
				Expect(t).To(Equal(CustomTypes{
					ColCustomIntType:         mockColCustomIntType,
					ColCustomInt8Type:        mockColCustomInt8Type,
					ColCustomInt16Type:       mockColCustomInt16Type,
					ColCustomInt32Type:       mockColCustomInt32Type,
					ColCustomInt64Type:       mockColCustomInt64Type,
					ColCustomFloat32Type:     mockColCustomFloat32Type,
					ColCustomFloat64Type:     mockColCustomFloat64Type,
					ColCustomUintType:        mockColCustomUintType,
					ColCustomUint8Type:       mockColCustomUint8Type,
					ColCustomUint16Type:      mockColCustomUint16Type,
					ColCustomUint32Type:      mockColCustomUint32Type,
					ColCustomUint64Type:      mockColCustomUint64Type,
					ColCustomBoolType:        mockColCustomBoolType,
					ColCustomStringType:      mockColCustomStringType,
					ColCustomTimeType:        mockColCustomTimeType,
					ColCustomSliceStringType: mockColCustomSliceStringType,
					ColSliceCustomBool:       mockColSliceCustomBool,
				}))
			})
		})
	})

	Context("pointers", func() {
		It("should return randomised struct", func() {
			t := CustomTypesPtr{}
			Expect(r.Struct(&t)).To(Succeed())
			Expect(t).To(Equal(CustomTypesPtr{
				ColCustomIntType:         &mockColCustomIntType,
				ColCustomInt8Type:        &mockColCustomInt8Type,
				ColCustomInt16Type:       &mockColCustomInt16Type,
				ColCustomInt32Type:       &mockColCustomInt32Type,
				ColCustomInt64Type:       &mockColCustomInt64Type,
				ColCustomFloat32Type:     &mockColCustomFloat32Type,
				ColCustomFloat64Type:     &mockColCustomFloat64Type,
				ColCustomUintType:        &mockColCustomUintType,
				ColCustomUint8Type:       &mockColCustomUint8Type,
				ColCustomUint16Type:      &mockColCustomUint16Type,
				ColCustomUint32Type:      &mockColCustomUint32Type,
				ColCustomUint64Type:      &mockColCustomUint64Type,
				ColCustomBoolType:        &mockColCustomBoolType,
				ColCustomStringType:      &mockColCustomStringType,
				ColCustomTimeType:        &mockColCustomTimeType,
				ColCustomSliceStringType: &mockColCustomSliceStringType,
				ColSliceCustomBool:       &mockColSliceCustomBool,
			}))
		})
	})

	Describe("AddTypeConfig() is called", func() {
		Describe("when OneOf() Provider", func() {

			type EnumType string
			type Test struct {
				Field EnumType
			}

			var (
				optionA = EnumType("option_a")
				optionB = EnumType("option_b")
			)

			Context("when is declared with incorrect type", func() {
				It("it should return MalformedProviderType", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.Config{
						Provider: randomise.OneOf("option_a", "option_b"),
					})
					err := r.Struct(&t)
					Expect(err).To(BeAssignableToTypeOf(randomise.MalformedProviderType{}))
				})
			})

			Context("when is declared with correct type", func() {
				It("it should set field", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.Config{
						Provider: randomise.OneOf(optionA, optionB),
					})
					Expect(r.Struct(&t)).To(Succeed())
					Expect(t.Field).To(Equal(optionB))
				})
			})
		})

		Describe("when As() Provider", func() {

			type Test struct {
				Field string
			}

			Context("when is declared with incorrect type", func() {
				It("it should return MalformedProviderType", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.Config{
						Provider: randomise.As(100),
					})
					err := r.Struct(&t)
					Expect(err).To(BeAssignableToTypeOf(randomise.MalformedProviderType{}))
				})
			})

			Context("when is declared with correct type", func() {
				It("it should set field", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.Config{
						Provider: randomise.As("option_a"),
					})
					Expect(r.Struct(&t)).To(Succeed())
					Expect(t.Field).To(Equal("option_a"))
				})
			})
		})
	})
})

func timePtr(v time.Time) *time.Time {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}

func uint32Ptr(v uint32) *uint32 {
	return &v
}

func uint16Ptr(v uint16) *uint16 {
	return &v
}

func uintPtr(v uint) *uint {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}

func float32Ptr(v float32) *float32 {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func int32Ptr(v int32) *int32 {
	return &v
}

func int16Ptr(v int16) *int16 {
	return &v
}

func int8Ptr(v int8) *int8 {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func uint8Ptr(v uint8) *uint8 {
	return &v
}

type NotTaggedStruct struct {
	Latitude         float32
	Long             float32
	CreditCardType   string
	CreditCardNumber string
	Email            string
	IPV4             string
	IPV6             string
}

func BenchmarkRandom_AddTypeProvider(b *testing.B) {
	mockDate := time.Date(1989, 4, 10, 0, 0, 0, 0, time.UTC)
	r := randomise.NewRandomise()
	r.SetSeed(mockDate.Unix())
	for i := 0; i < b.N; i++ {
		t := NotTaggedStruct{}
		if err := r.Struct(&t); err != nil {
			b.Fatal(err)
		}
	}
}
