package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToIntAndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum int
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "-100",
			expectedNum: -100,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "2147483647",
			expectedNum: 2147483647,
		},
		{
			numStr:      "-2147483647",
			expectedNum: -2147483647,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseInt: parsing \"1a2345\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToInt(table.numStr)
		numStr := ConvertIntToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)

		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}
	}
}

func TestConvertStringToInt8AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum int8
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "-100",
			expectedNum: -100,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "127",
			expectedNum: 127,
		},
		{
			numStr:      "-128",
			expectedNum: -128,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseInt: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 127,
			expectedStr: "127",
			expectedErr: "strconv.ParseInt: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-9223372036854775807",
			expectedNum: -128,
			expectedStr: "-128",
			expectedErr: "strconv.ParseInt: parsing \"-9223372036854775807\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToInt8(table.numStr)
		numStr := ConvertInt8ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}

	}
}

func TestConvertStringToInt16AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum int16
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "-100",
			expectedNum: -100,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "32767",
			expectedNum: 32767,
		},
		{
			numStr:      "-32768",
			expectedNum: -32768,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseInt: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 32767,
			expectedStr: "32767",
			expectedErr: "strconv.ParseInt: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-9223372036854775807",
			expectedNum: -32768,
			expectedStr: "-32768",
			expectedErr: "strconv.ParseInt: parsing \"-9223372036854775807\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToInt16(table.numStr)
		numStr := ConvertInt16ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}
	}
}

func TestConvertStringToInt32AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum int32
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "-100",
			expectedNum: -100,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "2147483647",
			expectedNum: 2147483647,
		},
		{
			numStr:      "-2147483648",
			expectedNum: -2147483648,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseInt: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 2147483647,
			expectedStr: "2147483647",
			expectedErr: "strconv.ParseInt: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-9223372036854775807",
			expectedNum: -2147483648,
			expectedStr: "-2147483648",
			expectedErr: "strconv.ParseInt: parsing \"-9223372036854775807\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToInt32(table.numStr)
		numStr := ConvertInt32ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}
	}
}

func TestConvertStringToInt64AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum int64
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "-100",
			expectedNum: -100,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "2147483647",
			expectedNum: 2147483647,
		},
		{
			numStr:      "-2147483648",
			expectedNum: -2147483648,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseInt: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "92233720368547758071",
			expectedNum: 9223372036854775807,
			expectedStr: "9223372036854775807",
			expectedErr: "strconv.ParseInt: parsing \"92233720368547758071\": value out of range",
		},
		{
			numStr:      "-92233720368547758071",
			expectedNum: -9223372036854775808,
			expectedStr: "-9223372036854775808",
			expectedErr: "strconv.ParseInt: parsing \"-92233720368547758071\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToInt64(table.numStr)
		numStr := ConvertInt64ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}
	}
}

func TestConvertStringToUIntAndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum uint
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "4294967295",
			expectedNum: 4294967295,
		},
		{
			numStr:      "-2147483647",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"-2147483647\": invalid syntax",
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"1a2345\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToUInt(table.numStr)
		numStr := ConvertUIntToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)

		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}
	}
}

func TestConvertStringToUInt8AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum uint8
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "255",
			expectedNum: 255,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 255,
			expectedStr: "255",
			expectedErr: "strconv.ParseUint: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-100",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"-100\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToUInt8(table.numStr)
		numStr := ConvertUInt8ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}

	}
}

func TestConvertStringToUInt16AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum uint16
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "65535",
			expectedNum: 65535,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 65535,
			expectedStr: "65535",
			expectedErr: "strconv.ParseUint: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-100",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"-100\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToUInt16(table.numStr)
		numStr := ConvertUInt16ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}

	}
}

func TestConvertStringToUInt32AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum uint32
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "4294967295",
			expectedNum: 4294967295,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "9223372036854775807",
			expectedNum: 4294967295,
			expectedStr: "4294967295",
			expectedErr: "strconv.ParseUint: parsing \"9223372036854775807\": value out of range",
		},
		{
			numStr:      "-100",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"-100\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToUInt32(table.numStr)
		numStr := ConvertUInt32ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}

	}
}

func TestConvertStringToUInt64AndBack_MultipleCase_Success(t *testing.T) {
	// GIVNE
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum uint64
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
		},
		{
			numStr:      "100",
			expectedNum: 100,
		},
		{
			numStr:      "18446744073709551615",
			expectedNum: 18446744073709551615,
		},
		{
			numStr:      "1a2345",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"1a2345\": invalid syntax",
		},
		{
			numStr:      "92233720368547758071",
			expectedNum: 18446744073709551615,
			expectedStr: "18446744073709551615",
			expectedErr: "strconv.ParseUint: parsing \"92233720368547758071\": value out of range",
		},
		{
			numStr:      "-100",
			expectedNum: 0,
			expectedStr: "0",
			expectedErr: "strconv.ParseUint: parsing \"-100\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToUInt64(table.numStr)
		numStr := ConvertUInt64ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.numStr, numStr)
		}

	}
}

func TestConvertStringToBool_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		boolStr      string
		expectedBool bool
		expectedErr  string
	}{
		{
			boolStr:      "1",
			expectedBool: true,
		},
		{
			boolStr:      "t",
			expectedBool: true,
		},
		{
			boolStr:      "T",
			expectedBool: true,
		},
		{
			boolStr:      "true",
			expectedBool: true,
		},
		{
			boolStr:      "True",
			expectedBool: true,
		},
		{
			boolStr:      "TRUE",
			expectedBool: true,
		},
		{
			boolStr:      "0",
			expectedBool: false,
		},
		{
			boolStr:      "f",
			expectedBool: false,
		},
		{
			boolStr:      "F",
			expectedBool: false,
		},
		{
			boolStr:      "false",
			expectedBool: false,
		},
		{
			boolStr:      "False",
			expectedBool: false,
		},
		{
			boolStr:      "FALSE",
			expectedBool: false,
		},
		{
			boolStr:      "atrue",
			expectedBool: false,
			expectedErr:  "strconv.ParseBool: parsing \"atrue\": invalid syntax",
		},
	}

	for _, table := range tables {
		// WHEN
		b, err := ConvertStringToBool(table.boolStr)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(false, b)
			assert.Equal(table.expectedErr, err.Error())
		} else {
			assert.Equal(table.expectedBool, b)
		}
	}
}

func TestConvertBoolToString_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	resTrue := ConvertBoolToString(true)
	resFalse := ConvertBoolToString(false)

	// THEN
	assert.Equal("true", resTrue)
	assert.Equal("false", resFalse)
}

func TestConvertStringToFloat32AndBack_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum float32
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
			expectedStr: "0E+00",
		},
		{
			numStr:      "1.33456789123",
			expectedNum: 1.33456789123,
			expectedStr: "1.3345679E+00",
		},
		{
			numStr:      "3.40282346638528859811704183484516925440e+38",
			expectedNum: 3.40282346638528859811704183484516925440e+38,
			expectedStr: "3.4028235E+38",
		},
		{
			numStr:      "-3.40282346638528859811704183484516925440e+38",
			expectedNum: -3.40282346638528859811704183484516925440e+38,
			expectedStr: "-3.4028235E+38",
		},
		{
			numStr:      "1.401298464324817070923729583289916131280e-45",
			expectedNum: 1.401298464324817070923729583289916131280e-45,
			expectedStr: "1E-45",
		},
		{
			numStr:      "1.797693134862315708145274237317043567981e+308",
			expectedErr: "strconv.ParseFloat: parsing \"1.797693134862315708145274237317043567981e+308\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToFloat32(table.numStr)
		numStr := ConvertFloat32ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		}
	}
}

func TestConvertStringToFloat64AndBack_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		numStr      string
		expectedNum float64
		expectedErr string
		expectedStr string
	}{
		{
			numStr:      "0",
			expectedNum: 0,
			expectedStr: "0E+00",
		},
		{
			numStr:      "1.33456789123",
			expectedNum: 1.33456789123,
			expectedStr: "1.33456789123E+00",
		},
		{
			numStr:      "1.797693134862315708145274237317043567981e+308",
			expectedNum: 1.797693134862315708145274237317043567981e+308,
			expectedStr: "1.7976931348623157E+308",
		},
		{
			numStr:      "-1.797693134862315708145274237317043567981e+308",
			expectedNum: -1.797693134862315708145274237317043567981e+308,
			expectedStr: "-1.7976931348623157E+308",
		},
		{
			numStr:      "4.940656458412465441765687928682213723651e-324",
			expectedNum: 4.940656458412465441765687928682213723651e-324,
			expectedStr: "5E-324",
		},
		{
			numStr:      "1.797693134862315708145274237317043567981e+309",
			expectedErr: "strconv.ParseFloat: parsing \"1.797693134862315708145274237317043567981e+309\": value out of range",
		},
	}

	for _, table := range tables {
		// WHEN
		num, err := ConvertStringToFloat64(table.numStr)
		numStr := ConvertFloat64ToString(num)

		// THEN
		if len(table.expectedErr) > 0 {
			assert.NotNil(err)
			assert.Equal(table.expectedErr, err.Error())
		} else {
			assert.Nil(err)
			assert.Equal(table.expectedNum, num)
			assert.Equal(table.expectedStr, numStr)
		}
	}
}

func BenchmarkConvertStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := ConvertStringToInt("123456")
		_ = val
	}
}

func BenchmarkConvertStringToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := ConvertStringToInt32("123456")
		_ = val
	}
}

func BenchmarkConvertStringToIn64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := ConvertStringToInt64("123456")
		_ = val
	}
}

func BenchmarkConvertIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := ConvertIntToString(123456)
		_ = val
	}
}

func BenchmarkConvertInt32ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := ConvertInt32ToString(123456)
		_ = val
	}
}

func BenchmarkConvertInt64ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := ConvertInt64ToString(123456)
		_ = val
	}
}
