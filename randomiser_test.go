package randomiser_test

import (
	"testing"
	"time"

	randomise "github.com/amwill04/randomiser"
	. "github.com/onsi/ginkgo/v2"
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

	Context("When called without using CreateNew()", func() {
		type Test struct {
			Field string
		}
		It("should return MalformedRandom", func() {
			t := Test{}
			rand := randomise.Random{}
			Expect(rand.Struct(&t)).ToNot(Succeed())
		})
	})

	var (
		r        randomise.Random
		mockDate = time.Date(1989, 4, 10, 0, 0, 0, 0, time.UTC)
		// returned values for seed
		mockColInt                  = 8787071967939076494
		mockColInt8         int8    = 107
		mockColInt16        int16   = 9439
		mockColInt32        int32   = 1390727781
		mockColInt64        int64   = 7308617895889833901
		mockColFloat32      float32 = 0.35669761896133423
		mockColFloat64              = 0.8673022323772911
		mockColUint         uint    = 16664198805663444872
		mockColUint8        uint8   = 143
		mockColUint16       uint16  = 57949
		mockColUint32       uint32  = 448464110
		mockColUint64       uint64  = 15820576590339358265
		mockColBool                 = true
		mockColString               = "szsVg"
		mockColTime                 = time.Date(2019, 7, 10, 10, 29, 10, 0, time.UTC)
		mockColByte                 = []byte{132, 214, 122}
		mockColSliceInt             = []int{3230138054492229965, 5693379029502621857, 9136199962177605632}
		mockColSliceInt8            = []int8{63, 64, 113}
		mockColSliceInt16           = []int16{34, 19738, 31068}
		mockColSliceInt32           = []int32{1272815704, 60700214, 412302932}
		mockColSliceInt64           = []int64{2754645531795864364, 2366225202029597505, 4487628403179330665}
		mockColSliceFloat32         = []float32{0.6270144581794739, 0.301972895860672, 0.5439574122428894}
		mockColSliceFloat64         = []float64{0.6523531507725464, 0.43047202351692165, 0.03410745948605001}
		mockColSliceUint            = []uint{1168718378530842714, 2176072750557506790, 14978333317498209884}
		mockColSliceUint8           = []uint8{109, 192, 211}
		mockColSliceUint16          = []uint16{37344, 48231, 18500}
		mockColSliceUint32          = []uint32{3677989657, 3361622017, 3481103586}
		mockColSliceUint64          = []uint64{8682075873787985163, 11827493833822128684, 4463540894904265705}
		mockColSliceBool            = []bool{true, false, true}
		mockColSliceString          = []string{"pirWZ", "CDYBY", "qrVma"}
		mockColSliceTime            = []time.Time{
			time.Date(2012, 3, 12, 6, 19, 32, 0, time.UTC),
			time.Date(1993, 11, 6, 19, 19, 31, 0, time.UTC),
			time.Date(2098, 11, 7, 6, 18, 2, 0, time.UTC),
		}
		mockColSlicePtrByte    = []*byte{uint8Ptr(82), uint8Ptr(213), uint8Ptr(133)}
		mockColSlicePtrInt     = []*int{intPtr(6939730804350152927), intPtr(2668961733473876525), intPtr(8453202931772817696)}
		mockColSlicePtrInt8    = []*int8{int8Ptr(41), int8Ptr(7), int8Ptr(96)}
		mockColSlicePtrInt16   = []*int16{int16Ptr(13713), int16Ptr(26634), int16Ptr(19476)}
		mockColSlicePtrInt32   = []*int32{int32Ptr(2066005117), int32Ptr(279068369), int32Ptr(111637746)}
		mockColSlicePtrInt64   = []*int64{int64Ptr(4803358365381946569), int64Ptr(1429517074973534657), int64Ptr(8973270389744226874)}
		mockColSlicePtrFloat32 = []*float32{float32Ptr(0.8568808436393738), float32Ptr(0.9400718808174133), float32Ptr(0.18035414814949036)}
		mockColSlicePtrFloat64 = []*float64{float64Ptr(0.9655046289649432), float64Ptr(0.9106112503030606), float64Ptr(0.793657749423185)}
		mockColSlicePtrUint    = []*uint{uintPtr(10627171034352290137), uintPtr(14859217705895579391), uintPtr(15812424427762282072)}
		mockColSlicePtrUint8   = []*uint8{uint8Ptr(231), uint8Ptr(187), uint8Ptr(134)}
		mockColSlicePtrUint16  = []*uint16{uint16Ptr(38239), uint16Ptr(48701), uint16Ptr(53547)}
		mockColSlicePtrUint32  = []*uint32{uint32Ptr(187490652), uint32Ptr(1814042450), uint32Ptr(3321555961)}
		mockColSlicePtrUint64  = []*uint64{uint64Ptr(5117774601994015311), uint64Ptr(14646683301499314643), uint64Ptr(14068097394368840198)}
		mockColSlicePtrBool    = []*bool{boolPtr(true), boolPtr(true), boolPtr(true)}
		mockColSlicePtrString  = []*string{stringPtr("xmQBu"), stringPtr("gsEcy"), stringPtr("AwVej")}
		mockColSlicePtrTime    = []*time.Time{
			timePtr(time.Date(2016, 6, 26, 11, 22, 8, 0, time.UTC)),
			timePtr(time.Date(2071, 3, 29, 3, 57, 12, 0, time.UTC)),
			timePtr(time.Date(2086, 10, 15, 19, 51, 38, 0, time.UTC)),
		}
		mockColStruct = StructA{
			ColStructB: StructB{
				StructC: StructC{
					ColDeepNest: 6869688774120787965,
				},
			},
			ColNested: "CgtFj",
		}
		mockColMapStringString                               = map[string]string{"bzBov": "WFmAc", "HWcll": "TEhle", "VSzpP": "IilfJ"}
		mockColMapStringPtrString                            = map[string]*string{"syBkL": stringPtr("COGNS"), "RgKdp": stringPtr("nEtbz"), "kZXwm": stringPtr("arrVv")}
		mockColMapStringInt                                  = map[string]int{"fxMpM": 2957738948002534632, "xsoju": 5243786676071560194, "qNLRA": 8718281656082008399}
		mockColMapStringPtrInt                               = map[string]*int{"bDAPC": intPtr(421457854334528867), "aNLFn": intPtr(657801808822237928), "UlbPV": intPtr(7882030331052233079)}
		mockColMapIntSliceString                             = map[int][]string{6812710852133740386: {"Ertad", "lImHf", "sGdYB"}, 3514875530947402435: {"eccsB", "XmUBX", "tZzAj"}, 1640979893917237826: {"hFnbG", "xkvXt", "mSYVx"}}
		mockColMapIntPtrSliceString                          = map[int]*[]string{689832524063967629: {"vLxga", "kksuq", "yNovZ"}, 8510854009147269140: {"rLlFf", "ZmYHe", "mYlpA"}, 1855461539089785566: {"TkYHb", "RPfsc", "JrhTA"}}
		mockColMapIntSlicePtrString                          = map[int][]*string{7796427636021573468: {stringPtr("PERcQ"), stringPtr("KXoon"), stringPtr("nRufA")}, 5170733929529154612: {stringPtr("cuvaN"), stringPtr("zfsPk"), stringPtr("bVfZv")}, 5193202225257715270: {stringPtr("JBscS"), stringPtr("PaaFD"), stringPtr("Qfsmz")}}
		mockColMapIntPtrSlicePtrString                       = map[int]*[]*string{6772078089664917898: {stringPtr("MTsVT"), stringPtr("BGlRZ"), stringPtr("UsoEr")}, 6953897037863106182: {stringPtr("weGSP"), stringPtr("CAWxF"), stringPtr("uOuvt")}, 3645567573765657738: {stringPtr("RGvRh"), stringPtr("rmfEi"), stringPtr("xziTt")}}
		mockColMapStringMapStringString                      = map[string]map[string]string{"LHmwP": {"kBAeK": "tQCvz", "lwRko": "yoVfJ", "SOWsn": "ptSvs"}, "kJkjn": {"HOceu": "Fqxig", "ZrtaI": "jvZWA", "TRiIM": "QimIp"}, "eJjSF": {"QaFvF": "OXmHH", "PUjDE": "VUlic", "lHJUS": "FuNfc"}}
		mockColMapStringPtrMapStringString                   = map[string]*map[string]string{"LWQrB": {"DqAym": "rqEfA", "yuBbg": "AmEcM", "IsvID": "McxGs"}, "jraaQ": {"SVuyJ": "dhRrn", "wKFkV": "QauLK", "OKfin": "EjZwS"}, "MFkvz": {"VCwpw": "bsazL", "qPObn": "FaabD", "QlcGX": "ufZzi"}}
		mockColSliceSliceString                              = [][]string{{"OBTld", "vIBUb", "JQlPb"}, {"eWyVF", "cljXf", "umdZs"}, {"LTJJr", "SuNsk", "ICjJX"}}
		mockColSliceSlicePtrString                           = [][]*string{{stringPtr("ZulkY"), stringPtr("oOlhb"), stringPtr("jmPOs")}, {stringPtr("zESrm"), stringPtr("GqZYz"), stringPtr("SEICU")}, {stringPtr("FQkgr"), stringPtr("LYxuy"), stringPtr("HGpsR")}}
		mockColArrayString                                   = [2]string{"fvsNc", "lvNgA"}
		mockColArrayInt8ArrayPtr                             = [2]*[2]*int8{{int8Ptr(125), int8Ptr(67)}, {int8Ptr(22), int8Ptr(127)}}
		mockColCustomIntType               CustomIntType     = 8787071967939076494
		mockColCustomInt8Type              CustomInt8Type    = 107
		mockColCustomInt16Type             CustomInt16Type   = 9439
		mockColCustomInt32Type             CustomInt32Type   = 1390727781
		mockColCustomInt64Type             CustomInt64Type   = 7308617895889833901
		mockColCustomFloat32Type           CustomFloat32Type = 0.35669761896133423
		mockColCustomFloat64Type           CustomFloat64Type = 0.8673022323772911
		mockColCustomUintType              CustomUintType    = 16664198805663444872
		mockColCustomUint8Type             CustomUint8Type   = 143
		mockColCustomUint16Type            CustomUint16Type  = 57949
		mockColCustomUint32Type            CustomUint32Type  = 448464110
		mockColCustomUint64Type            CustomUint64Type  = 15820576590339358265
		mockColCustomBoolType              CustomBoolType    = true
		mockColCustomStringType            CustomStringType  = "szsVg"
		mockColCustomTimeType                                = CustomTimeType(time.Date(2019, 7, 10, 10, 29, 10, 0, time.UTC))
		mockColCustomSliceStringType                         = CustomSliceStringType{"LWPAe", "sEDFw", "nPjMd"}
		mockColSliceCustomBool                               = []CustomBoolType{false, true, false}
	)

	BeforeEach(func() {
		r = randomise.NewRandomise(mockDate.UnixNano())
	})

	Context("When called without passing pointer", func() {
		type Test struct {
			Field string
		}
		It("should return MalformedRandom", func() {
			t := Test{}
			Expect(r.Struct(t)).ToNot(Succeed())
		})
	})

	Context("when non struct type is passed", func() {
		It("should return error", func() {
			var test string
			Expect(r.Struct(&test)).ToNot(Succeed())
		})
	})

	Context("when has private fields ", func() {
		type Test struct {
			FieldA          string
			FieldB          *string
			privateField    string
			privateFieldPtr *string
		}

		It("should not return error", func() {
			var t = Test{}
			Expect(r.Struct(&t)).To(Succeed())
		})
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

	Describe("When globals are updated", func() {
		type Test struct {
			FieldString string
			FieldSlice  []string
			FieldMap    map[string]string
		}

		Context("when when setters are called", func() {
			It("should set all values", func() {
				t := Test{}
				r.SetStringLength(10)
				r.SetMapKeyLength(5)
				r.SetMapLength(10)
				r.SetSliceLength(10)
				Expect(r.Struct(&t)).To(Succeed())
				Expect(len(t.FieldString)).To(BeNumerically("==", 10))
				Expect(len(t.FieldMap)).To(BeNumerically("==", 10))
				for k, v := range t.FieldMap {
					Expect(len(k)).To(BeNumerically("==", 5))
					Expect(len(v)).To(BeNumerically("==", 10))
				}
				Expect(len(t.FieldSlice)).To(BeNumerically("==", 10))
				for _, v := range t.FieldSlice {
					Expect(len(v)).To(BeNumerically("==", 10))
				}
			})
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
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.OneOf("option_a", "option_b")))
					err := r.Struct(&t)
					Expect(err).To(BeAssignableToTypeOf(randomise.MalformedProviderType{}))
				})
			})

			Context("when is declared with value rather than pointer", func() {
				type TestPtr struct {
					Field *EnumType
				}

				It("it should set field", func() {
					t := TestPtr{}
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.OneOf(optionA, optionB)))
					Expect(r.Struct(&t)).To(Succeed())
					Expect(*t.Field).To(Equal(optionA))
				})
			})

			Context("when is declared with correct type", func() {
				It("it should set field", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.OneOf(optionA, optionB)))
					Expect(r.Struct(&t)).To(Succeed())
					Expect(t.Field).To(Equal(optionA))
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
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.As(100)))
					err := r.Struct(&t)
					Expect(err).To(BeAssignableToTypeOf(randomise.MalformedProviderType{}))
				})
			})

			Context("when is declared with value rather than pointer", func() {
				type TestPtr struct {
					Field *string
				}

				It("it should set field", func() {
					t := TestPtr{}
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.As("option_a")))
					Expect(r.Struct(&t)).To(Succeed())
					Expect(*t.Field).To(Equal("option_a"))
				})
			})

			Context("when is declared with correct type", func() {
				It("it should set field", func() {
					t := Test{}
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.As("option_a")))
					Expect(r.Struct(&t)).To(Succeed())
					Expect(t.Field).To(Equal("option_a"))
				})
			})
		})

		Describe("when Between() Provider", func() {

			Context("when field is not supported type", func() {
				type Test struct {
					Field string
				}

				var (
					t Test
				)

				BeforeEach(func() {
					t = Test{}
				})

				It("should return MalformedProviderUnsupportedType ", func() {
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.Between("a", "c")))
					Expect(r.Struct(&t)).ToNot(Succeed())
				})
			})

			Context("when  arguments are incorrect type", func() {
				type Test struct {
					Field int
				}

				var (
					t Test
				)

				BeforeEach(func() {
					t = Test{}
				})

				It("should return MalformedProviderType with first argument", func() {
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.Between(mockColInt8, mockColInt)))
					Expect(r.Struct(&t)).ToNot(Succeed())
				})

				It("should return MalformedProviderType with second argument", func() {
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.Between(mockColInt, mockColInt8)))
					Expect(r.Struct(&t)).ToNot(Succeed())
				})
			})

			Context("when end < start number", func() {
				type Test struct {
					Field int
				}

				var (
					t Test
				)

				BeforeEach(func() {
					t = Test{}
				})

				It("should return MalformedProviderType", func() {
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.Between(10, 1)))
					Expect(r.Struct(&t)).ToNot(Succeed())
				})
			})

			Context("when end < start time", func() {
				type Test struct {
					Field time.Time
				}

				var (
					t Test
				)

				BeforeEach(func() {
					t = Test{}
				})

				It("should return MalformedProviderType", func() {
					r.AddTypeConfig("Field", randomise.WithProvider(randomise.Between(mockColTime.AddDate(0, 0, 1), mockColTime)))
					Expect(r.Struct(&t)).ToNot(Succeed())
				})
			})

			Context("when is declared with correct type", func() {
				type Test struct {
					Int           int
					IntPtr        *int
					Int8          int8
					Int8Ptr       *int8
					Int16         int16
					Int16Ptr      *int16
					Int32         int32
					Int32Ptr      *int32
					Int64         int64
					Int64Ptr      *int64
					Time          time.Time
					TimePtr       *time.Time
					CustomInt     CustomIntType
					CustomPtrInt  *CustomIntType
					CustomTime    CustomTimeType
					CustomPtrTime *CustomTimeType
				}

				It("it should set field", func() {
					t := Test{}
					r.AddTypeConfig("Int", randomise.WithProvider(randomise.Between(10, 100)))
					r.AddTypeConfig("IntPtr", randomise.WithProvider(randomise.Between(intPtr(10), intPtr(100))))
					r.AddTypeConfig("Int8", randomise.WithProvider(randomise.Between(int8(10), int8(100))))
					r.AddTypeConfig("Int8Ptr", randomise.WithProvider(randomise.Between(int8Ptr(10), int8Ptr(100))))
					r.AddTypeConfig("Int16", randomise.WithProvider(randomise.Between(int16(10), int16(100))))
					r.AddTypeConfig("Int16Ptr", randomise.WithProvider(randomise.Between(int16Ptr(10), int16Ptr(100))))
					r.AddTypeConfig("Int32", randomise.WithProvider(randomise.Between(int32(10), int32(100))))
					r.AddTypeConfig("Int32Ptr", randomise.WithProvider(randomise.Between(int32Ptr(10), int32Ptr(100))))
					r.AddTypeConfig("Int64", randomise.WithProvider(randomise.Between(int64(10), int64(100))))
					r.AddTypeConfig("Int64Ptr", randomise.WithProvider(randomise.Between(int64Ptr(10), int64Ptr(100))))
					endMockColTime := mockColTime.AddDate(0, 0, 10)
					r.AddTypeConfig("Time", randomise.WithProvider(randomise.Between(mockColTime, endMockColTime)))
					r.AddTypeConfig("TimePtr", randomise.WithProvider(randomise.Between(&mockColTime, &endMockColTime)))
					endMockColCustomIntType := mockColCustomIntType + 10
					r.AddTypeConfig("CustomInt", randomise.WithProvider(randomise.Between(mockColCustomIntType, endMockColCustomIntType)))
					r.AddTypeConfig("CustomPtrInt", randomise.WithProvider(randomise.Between(&mockColCustomIntType, &endMockColCustomIntType)))
					endMockColCustomTimeType := CustomTimeType(endMockColTime)
					r.AddTypeConfig("CustomTime", randomise.WithProvider(randomise.Between(mockColCustomTimeType, endMockColCustomTimeType)))
					r.AddTypeConfig("CustomPtrTime", randomise.WithProvider(randomise.Between(&mockColCustomTimeType, &endMockColCustomTimeType)))
					Expect(r.Struct(&t)).To(Succeed())
				})
			})
		})

		Describe("When string length is provided", func() {
			type Test struct {
				Field        string
				DefaultField string
			}

			It("should set string length different to default", func() {
				t := Test{}
				r.AddTypeConfig("Field", randomise.WithStringLength(10))
				Expect(r.Struct(&t)).To(Succeed())
				Expect(len(t.Field)).To(BeNumerically("==", 10))
				Expect(len(t.DefaultField)).To(BeNumerically("==", randomise.StringLengthDefault))
			})
		})

		Describe("When slice length is provided", func() {
			type Test struct {
				Field        []string
				DefaultField []string
			}

			It("should set slice length different to default", func() {
				t := Test{}
				r.AddTypeConfig("Field", randomise.WithSliceLength(10))
				Expect(r.Struct(&t)).To(Succeed())
				Expect(len(t.Field)).To(BeNumerically("==", 10))
				Expect(len(t.DefaultField)).To(BeNumerically("==", randomise.SliceLengthDefault))
			})
		})

		Describe("When map length is provided", func() {
			type Test struct {
				Field map[string]int
				//DefaultField map[string]string
			}

			It("should set map length different to default", func() {
				t := Test{}
				r.AddTypeConfig("Field", randomise.WithMapLength(10))
				Expect(r.Struct(&t)).To(Succeed())
				Expect(len(t.Field)).To(BeNumerically("==", 10))
				//Expect(len(t.DefaultField)).To(BeNumerically("==", randomise.MapLengthDefault))
			})
		})

		Describe("When map key length is provided", func() {
			type Test struct {
				Field        map[string]string
				DefaultField map[string]string
			}

			It("should set map key different to default", func() {
				t := Test{}
				r.AddTypeConfig("Field", randomise.WithMapKeyLength(10))
				Expect(r.Struct(&t)).To(Succeed())
				for k := range t.Field {
					Expect(len(k)).To(BeNumerically("==", 10))
				}
				for k := range t.DefaultField {
					Expect(len(k)).To(BeNumerically("==", randomise.MapKeyLengthDefault))
				}
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

func BenchmarkRandom_Generate(b *testing.B) {
	mockDate := time.Date(1989, 4, 10, 0, 0, 0, 0, time.UTC)
	r := randomise.NewRandomise(mockDate.UnixNano())
	for i := 0; i < b.N; i++ {
		t := NotTaggedStruct{}
		if err := r.Struct(&t); err != nil {
			b.Fatal(err)
		}
	}
}
