package almiconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	badVal   = "badval"
	zeroStr  = "0"
	oneStr   = "1"
	twoStr   = "2"
	threeStr = "3"
	fourStr  = "4"
	fiveStr  = "5"
	sixStr   = "6"
	sevenStr = "7"
	eightStr = "8"
	nineStr  = "9"
	tenStr   = "10"

	zero  = uintptr(0)
	one   = int(1)
	two   = int8(2)
	three = int16(3)
	four  = int32(4)
	five  = int64(5)
	six   = uint(6)
	seven = uint8(7)
	eight = uint16(8)
	nine  = uint32(9)
	ten   = uint64(10)

	comma         = ","
	badStrSlice   = "0,x"
	zeroStrSlice  = "0,0"
	oneStrSlice   = "1,1"
	twoStrSlice   = "2,2"
	threeStrSlice = "3,3"
	fourStrSlice  = "4,4"
	fiveStrSlice  = "5,5"
	sixStrSlice   = "6,6"
	sevenStrSlice = "7,7"
	eightStrSlice = "8,8"
	nineStrSlice  = "9,9"
	tenStrSlice   = "10,10"

	empty       = ""
	strKey      = "str"
	strVal      = "strVal"
	strSliceVal = "strVal,strVal"

	boolKey         = "bool"
	boolVal         = "true"
	badBoolVal      = "notBoolVal"
	boolSliceVal    = "true,true"
	badBoolSliceVal = "true,notBoolVal"
	trueVal         = true
	falseVal        = false

	runeKey         = "rune"
	runeVal         = "65"
	badRuneVal      = "notRune"
	runeSliceVal    = "65,65"
	badRuneSliceVal = "65,notRune"
	runeA           = rune(65)
	runeFail        = rune(0)
)

var (
	zeroSlice  = []uintptr{zero, zero}
	oneSlice   = []int{one, one}
	twoSlice   = []int8{two, two}
	threeSlice = []int16{three, three}
	fourSlice  = []int32{four, four}
	fiveSlice  = []int64{five, five}
	sixSlice   = []uint{six, six}
	sevenSlice = []uint8{seven, seven}
	eightSlice = []uint16{eight, eight}
	nineSlice  = []uint32{nine, nine}
	tenSlice   = []uint64{ten, ten}

	strSlice = []string{strVal, strVal}

	boolSlice = []bool{trueVal, trueVal}

	runeSlice = []rune{65, 65}

	intTypes     = []string{_uintptr, _int, _int8, _int16, _int32, _int64, _uint, _uint8, _uint16, _uint32, _uint64}
	strVals      = []string{zeroStr, oneStr, twoStr, threeStr, fourStr, fiveStr, sixStr, sevenStr, eightStr, nineStr, tenStr}
	badStrVals   = []string{badVal, badVal, badVal, badVal, badVal, badVal, badVal, badVal, badVal, badVal, badVal}
	sliceVals    = []string{zeroStrSlice, oneStrSlice, twoStrSlice, threeStrSlice, fourStrSlice, fiveStrSlice, sixStrSlice, sevenStrSlice, eightStrSlice, nineStrSlice, tenStrSlice}
	badSliceVals = []string{badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice, badStrSlice}
)

func testSetEnv(t *testing.T, key, val string) {
	if err := os.Setenv(key, val); err != nil {
		t.Fail()
	}
}

func initIntEnv(t *testing.T, vals []string) {
	for i, _type := range intTypes {
		testSetEnv(t, _type, vals[i])
	}
}

func testAlmiAton[T number](t *testing.T, key string, expect any) {
	cc := configConstraint{EnvName: key}
	envVar, err := aton[T](cc)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, expect, envVar)
}

func testAlmiAtonSlice[T number](t *testing.T, key string, expect any) {
	cc := configConstraint{EnvName: key, SliceType: true, Separator: comma}
	envVar, err := aton[T](cc)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, expect, envVar)
}

func testAlmiAtonFail[T number](t *testing.T, key string) {
	cc := configConstraint{EnvName: key}
	envVar, err := aton[T](cc)
	assert.Equal(t, T(0), envVar)
	assert.NotNil(t, err)
}

func testAlmiAtonSliceFail[T number](t *testing.T, key string) {
	cc := configConstraint{EnvName: key, SliceType: true, Separator: comma}
	envVar, err := aton[T](cc)
	assert.Equal(t, T(0), envVar)
	assert.NotNil(t, err)
}

func TestAlmiAtoi_SuccessfullyConvertInts(t *testing.T) {
	initIntEnv(t, strVals)

	testAlmiAton[uintptr](t, _uintptr, zero)
	testAlmiAton[int](t, _int, one)
	testAlmiAton[int8](t, _int8, two)
	testAlmiAton[int16](t, _int16, three)
	testAlmiAton[int32](t, _int32, four)
	testAlmiAton[int64](t, _int64, five)
	testAlmiAton[uint](t, _uint, six)
	testAlmiAton[uint8](t, _uint8, seven)
	testAlmiAton[uint16](t, _uint16, eight)
	testAlmiAton[uint32](t, _uint32, nine)
	testAlmiAton[uint64](t, _uint64, ten)
}

func TestAlmiAtoi_FailConvertInts(t *testing.T) {
	initIntEnv(t, badStrVals)

	testAlmiAtonFail[uintptr](t, _uintptr)
	testAlmiAtonFail[int](t, _int)
	testAlmiAtonFail[int8](t, _int8)
	testAlmiAtonFail[int16](t, _int16)
	testAlmiAtonFail[int32](t, _int32)
	testAlmiAtonFail[int64](t, _int64)
	testAlmiAtonFail[uint](t, _uint)
	testAlmiAtonFail[uint8](t, _uint8)
	testAlmiAtonFail[uint16](t, _uint16)
	testAlmiAtonFail[uint32](t, _uint32)
	testAlmiAtonFail[uint64](t, _uint64)
}

func TestAlmiAtoi_SuccessfullyConvertIntSlices(t *testing.T) {
	initIntEnv(t, sliceVals)

	testAlmiAtonSlice[uintptr](t, _uintptr, zeroSlice)
	testAlmiAtonSlice[int](t, _int, oneSlice)
	testAlmiAtonSlice[int8](t, _int8, twoSlice)
	testAlmiAtonSlice[int16](t, _int16, threeSlice)
	testAlmiAtonSlice[int32](t, _int32, fourSlice)
	testAlmiAtonSlice[int64](t, _int64, fiveSlice)
	testAlmiAtonSlice[uint](t, _uint, sixSlice)
	testAlmiAtonSlice[uint8](t, _uint8, sevenSlice)
	testAlmiAtonSlice[uint16](t, _uint16, eightSlice)
	testAlmiAtonSlice[uint32](t, _uint32, nineSlice)
	testAlmiAtonSlice[uint64](t, _uint64, tenSlice)
}

func TestAlmiAtoi_FailConvertIntSlices(t *testing.T) {
	initIntEnv(t, badSliceVals)

	testAlmiAtonSliceFail[uintptr](t, _uintptr)
	testAlmiAtonSliceFail[int](t, _int)
	testAlmiAtonSliceFail[int8](t, _int8)
	testAlmiAtonSliceFail[int16](t, _int16)
	testAlmiAtonSliceFail[int32](t, _int32)
	testAlmiAtonSliceFail[int64](t, _int64)
	testAlmiAtonSliceFail[uint](t, _uint)
	testAlmiAtonSliceFail[uint8](t, _uint8)
	testAlmiAtonSliceFail[uint16](t, _uint16)
	testAlmiAtonSliceFail[uint32](t, _uint32)
	testAlmiAtonSliceFail[uint64](t, _uint64)
}

func TestAlmiStr_SuccessfullyConvertString(t *testing.T) {
	if err := os.Setenv(strKey, strVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: strKey}
	envVar, err := str[string](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, strVal, envVar)
}

func TestAlmiStr_FailConvertString(t *testing.T) {
	if err := os.Setenv(strKey, strVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: strKey, SliceType: true, Separator: empty}
	envVar, err := str[string](cc)
	assert.Equal(t, empty, envVar)
	assert.NotNil(t, err)
}

// No fail equivalent for this test
func TestAlmiStr_SuccessfullyConvertStringSlice(t *testing.T) {
	if err := os.Setenv(strKey, strSliceVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: strKey, SliceType: true, Separator: comma}
	envVar, err := str[string](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, strSlice, envVar)
}

func TestAlmiAtob_SuccessfullyConvertBool(t *testing.T) {
	if err := os.Setenv(boolKey, boolVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: boolKey}
	envVar, err := atob[bool](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, trueVal, envVar)
}

func TestAlmiAtob_FailConvertBool(t *testing.T) {
	if err := os.Setenv(boolKey, badBoolVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: boolKey}
	envVar, err := atob[bool](cc)
	assert.Equal(t, falseVal, envVar)
	assert.NotNil(t, err)
}

func TestAlmiAtob_SuccessfullyConvertBoolSlice(t *testing.T) {
	if err := os.Setenv(boolKey, boolSliceVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: boolKey, SliceType: true, Separator: comma}
	envVar, err := atob[bool](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, boolSlice, envVar)
}

func TestAlmiAtob_FailConvertBoolSlice(t *testing.T) {
	if err := os.Setenv(boolKey, badBoolSliceVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: boolKey, SliceType: true, Separator: comma}
	envVar, err := atob[bool](cc)
	assert.Equal(t, falseVal, envVar)
	assert.NotNil(t, err)
}

func TestAlmiAtoRB_SuccessfullyConvertRune(t *testing.T) {
	if err := os.Setenv(runeKey, runeVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: runeKey}
	envVar, err := atoRB[rune](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, runeA, envVar)
}

func TestAlmiAtoRB_FailConvertRune(t *testing.T) {
	if err := os.Setenv(runeKey, badRuneVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: runeKey}
	envVar, err := atoRB[rune](cc)
	assert.Equal(t, runeFail, envVar)
	assert.NotNil(t, err)
}

func TestAlmiAtoRB_SuccessfullyConvertRuneSlice(t *testing.T) {
	if err := os.Setenv(runeKey, runeSliceVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: runeKey, SliceType: true, Separator: comma}
	envVar, err := atoRB[rune](cc)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, runeSlice, envVar)
}

func TestAlmiAtoRB_FailConvertRuneSlice(t *testing.T) {
	if err := os.Setenv(runeKey, badRuneSliceVal); err != nil {
		t.Fail()
	}

	cc := configConstraint{EnvName: runeKey, SliceType: true, Separator: comma}
	envVar, err := atoRB[rune](cc)
	assert.Equal(t, runeFail, envVar)
	assert.NotNil(t, err)
}
