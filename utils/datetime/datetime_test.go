package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMiliseconds_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := GetCurrentMiliseconds()

	// THEN
	assert.True(t, millis > 1631359343911)
}

func TestConvertCurrentLocalTimeToString_SimpleInput_Success(t *testing.T) {
	// GIVEN
	tables := []struct {
		format string
	}{
		{format: YYYY_MM_DD},
		{format: YYYY_MM_DD_HH_MM_SS},
		{format: YYYY_MM_DD_HH_MM_SS_SSS},
		{format: DD_MM_YYYY},
		{format: DD_MM_YYYY_HH_MM_SS},
		{format: DD_MM_YYYY_HH_MM_SS_SSS},
	}
	for _, table := range tables {
		// WHEN
		datetime := ConvertCurrentLocalTimeToString(table.format)

		// THEN
		assert.NotEmpty(t, datetime)
	}
}

func TestConvertMillisecondsToString_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		millis           int64
		format           string
		expectedDatetime string
	}{
		{
			millis:           1631357757780,
			format:           YYYY_MM_DD,
			expectedDatetime: "2021-09-11",
		},
		{
			millis:           1631357757780,
			format:           YYYY_MM_DD_HH_MM_SS,
			expectedDatetime: "2021-09-11 17:55:57",
		},
		{
			millis:           1631357757780,
			format:           YYYY_MM_DD_HH_MM_SS_SSS,
			expectedDatetime: "2021-09-11 17:55:57.780",
		},
		{
			millis:           1631357757780,
			format:           DD_MM_YYYY,
			expectedDatetime: "11-09-2021",
		},
		{
			millis:           1631357757780,
			format:           DD_MM_YYYY_HH_MM_SS,
			expectedDatetime: "11-09-2021 17:55:57",
		},
		{
			millis:           1631357757780,
			format:           DD_MM_YYYY_HH_MM_SS_SSS,
			expectedDatetime: "11-09-2021 17:55:57.780",
		},
	}
	for _, table := range tables {
		// WHEN
		datetime := ConvertMillisecondsToString(table.millis, table.format)

		// THEN
		assert.NotEmpty(datetime)
		assert.Equal(table.expectedDatetime, datetime)
	}
}

func TestConvertStringToMilliseconds_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		actualDatetime string
		format         string
		expectedMillis int64
	}{
		{
			actualDatetime: "2021-09-11",
			format:         YYYY_MM_DD,
			expectedMillis: 1631293200000,
		},
		{
			actualDatetime: "2021-09-11 17:55:57",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedMillis: 1631357757000,
		},
		{
			actualDatetime: "2021-09-11 17:55:57.780",
			format:         YYYY_MM_DD_HH_MM_SS_SSS,
			expectedMillis: 1631357757780,
		},
		{
			actualDatetime: "11-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1631293200000,
		},
		{
			actualDatetime: "11-09-2021 17:55:57",
			format:         DD_MM_YYYY_HH_MM_SS,
			expectedMillis: 1631357757000,
		},
		{
			actualDatetime: "11-09-2021 17:55:57.780",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1631357757780,
		},
	}
	for _, table := range tables {
		// WHEN
		millis, err := ConvertStringToMilliseconds(table.actualDatetime, table.format)

		// THEN
		assert.Nil(err)
		assert.Equal(table.expectedMillis, millis)
	}
}

func TestConvertStringToMilliseconds_WrongFormat_FailedToConvert(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		datetimeStr    string
		format         string
		expectedErrStr string
	}{
		{
			datetimeStr:    "11-09-2021",
			format:         YYYY_MM_DD,
			expectedErrStr: "parsing time \"11-09-2021\" as \"2006-01-02\": cannot parse \"9-2021\" as \"2006\"",
		},
		{
			datetimeStr:    "11-09-2021",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedErrStr: "parsing time \"11-09-2021\" as \"2006-01-02 15:04:05\": cannot parse \"9-2021\" as \"2006\"",
		},
		{
			datetimeStr:    "123",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedErrStr: "parsing time \"123\" as \"2006-01-02 15:04:05\": cannot parse \"123\" as \"2006\"",
		},
		{
			datetimeStr:    "-123",
			format:         "abcdef",
			expectedErrStr: "parsing time \"-123\" as \"abcdef\": cannot parse \"-123\" as \"abcdef\"",
		},
		{
			datetimeStr:    "",
			format:         YYYY_MM_DD,
			expectedErrStr: "parsing time \"\" as \"2006-01-02\": cannot parse \"\" as \"2006\"",
		},
	}
	for _, table := range tables {
		// WHEN
		millis, err := ConvertStringToMilliseconds(table.datetimeStr, table.format)

		// THEN
		assert.Equal(int64(-1), millis)
		assert.NotNil(err)
		assert.Equal(table.expectedErrStr, err.Error())
	}
}

func TestConvertStringToLocalTime_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		actualDatetime string
		format         string
		expectedMillis int64
	}{
		{
			actualDatetime: "2021-09-11",
			format:         YYYY_MM_DD,
			expectedMillis: 1631293200000,
		},
		{
			actualDatetime: "2021-09-11 17:55:57",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedMillis: 1631357757000,
		},
		{
			actualDatetime: "2021-09-11 17:55:57.780",
			format:         YYYY_MM_DD_HH_MM_SS_SSS,
			expectedMillis: 1631357757780,
		},
		{
			actualDatetime: "11-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1631293200000,
		},
		{
			actualDatetime: "11-09-2021 17:55:57",
			format:         DD_MM_YYYY_HH_MM_SS,
			expectedMillis: 1631357757000,
		},
		{
			actualDatetime: "11-09-2021 17:55:57.780",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1631357757780,
		},
	}
	for _, table := range tables {
		// WHEN
		localTime, err := ConvertStringToLocalTime(table.actualDatetime, table.format)

		// THEN
		assert.Nil(err)
		assert.NotNil(localTime)
		assert.Equal(table.expectedMillis, ConvertLocalTimeToMilliseconds(localTime))
	}
}

func TestConvertStringToLocal_WrongFormat_FailedToConvert(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		datetimeStr    string
		format         string
		expectedErrStr string
	}{
		{
			datetimeStr:    "11-09-2021",
			format:         YYYY_MM_DD,
			expectedErrStr: "parsing time \"11-09-2021\" as \"2006-01-02\": cannot parse \"9-2021\" as \"2006\"",
		},
		{
			datetimeStr:    "11-09-2021",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedErrStr: "parsing time \"11-09-2021\" as \"2006-01-02 15:04:05\": cannot parse \"9-2021\" as \"2006\"",
		},
		{
			datetimeStr:    "123",
			format:         YYYY_MM_DD_HH_MM_SS,
			expectedErrStr: "parsing time \"123\" as \"2006-01-02 15:04:05\": cannot parse \"123\" as \"2006\"",
		},
		{
			datetimeStr:    "-123",
			format:         "abcdef",
			expectedErrStr: "parsing time \"-123\" as \"abcdef\": cannot parse \"-123\" as \"abcdef\"",
		},
		{
			datetimeStr:    "",
			format:         YYYY_MM_DD,
			expectedErrStr: "parsing time \"\" as \"2006-01-02\": cannot parse \"\" as \"2006\"",
		},
	}
	for _, table := range tables {
		// WHEN
		localTime, err := ConvertStringToLocalTime(table.datetimeStr, table.format)

		// THEN
		assert.NotNil(localTime)
		assert.NotNil(err)
		assert.Equal("0001-01-01 00:00:00 +0000 UTC", localTime.String())
		assert.Equal(table.expectedErrStr, err.Error())
	}
}

func TestConvertLocalTimeToMilliseconds_SimpleInput_Success(t *testing.T) {
	// GIVEN
	currentTime := time.Now()

	// WHEN
	millis := ConvertLocalTimeToMilliseconds(currentTime)

	// THEN
	assert.True(t, millis > 1631359343911)
}

func TestGetCurrentLocalTime_SimpleInput_Succss(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	// WHEN
	currentTime := GetCurrentLocalTime()
	millis := ConvertLocalTimeToMilliseconds(currentTime)

	// THEN
	assert.NotNil(currentTime)
	assert.True(millis > 1631359343911)
}

func TestConvertLocalTimeToString_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	localTime, _ := ConvertStringToLocalTime("09-09-2021", DD_MM_YYYY)

	// WHEN
	datetimeStr := ConvertLocalTimeToString(localTime, DD_MM_YYYY_HH_MM_SS)

	assert.NotEmpty(datetimeStr)
	assert.Equal("09-09-2021 00:00:00", datetimeStr)
}

func TestGetYear_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	year := GetYear()

	// THEN
	assert.True(t, year > 2020)
}

func TestGetDayOfYear_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	dayOfYear := GetDayOfYear()

	// THEN
	assert.True(t, dayOfYear >= 1 && dayOfYear < 366)
}

func TestGetDayOfMonth_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	dayOfMonth := GetDayOfMonth()

	// THEN
	assert.True(t, dayOfMonth >= 1 && dayOfMonth <= 31)
}

func TestGetMonthOfYear_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	monthOfYear := GetMonthOfYear()

	// THEN
	assert.True(t, monthOfYear >= 1 && monthOfYear <= 12)
}

func TestGetStartLocalTimeOfYear_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetStartLocalTimeOfYear())

	// THEN
	assert.True(t, millis > 1 && millis <= GetCurrentMiliseconds())
}

func TestGetEndLocalTimeOfYear_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetEndLocalTimeOfYear())

	// THEN
	assert.True(t, millis >= GetCurrentMiliseconds())
}

func TestGetStartLocalTimeOfMonth_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetStartLocalTimeOfMonth())

	// THEN
	assert.True(t, millis > 1 && millis <= GetCurrentMiliseconds())
}

func TestGetEndLocalTimeOfMonth_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetEndLocalTimeOfMonth())

	// THEN
	assert.True(t, millis >= GetCurrentMiliseconds())
}

func TestGetStartLocalTimeOfDay_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetStartLocalTimeOfDay())

	// THEN
	assert.True(t, millis > 1 && millis <= GetCurrentMiliseconds())
}

func TestGetEndLocalTimeOfDay_SimpleInput_Success(t *testing.T) {
	// GIVEN
	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetEndLocalTimeOfDay())

	// THEN
	assert.True(t, millis >= GetCurrentMiliseconds())
}

func TestGetStartLocalTimeOfTime_SimpleInput_Success(t *testing.T) {
	// GIVEN
	localTime, _ := ConvertStringToLocalTime("09-09-2021", DD_MM_YYYY)

	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetStartLocalTimeOfTime(localTime))

	// THEN
	assert.Equal(t, int64(1631120400000), millis)
}

func TestGetEndLocalTimeOfTime_SimpleInput_Success(t *testing.T) {
	// GIVEN
	localTime, _ := ConvertStringToLocalTime("09-09-2021", DD_MM_YYYY)

	// WHEN
	millis := ConvertLocalTimeToMilliseconds(GetEndLocalTimeOfTime(localTime))

	// THEN
	assert.Equal(t, int64(1631206799999), millis)
}

func TestGetBeforeLocalTimeOfTime_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		datetimeStr    string
		format         string
		expectedMillis int64
		numberDay      int
		isStart        bool
	}{
		{
			datetimeStr:    "01-01-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609347600000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "01-01-2021 23:59:59.999",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609347600000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "01-01-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609347600000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "01-01-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609261200000,
			numberDay:      2,
			isStart:        true,
		},
		{
			datetimeStr:    "01-02-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612026000000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "01-02-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1611853200000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "09-09-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1630861200000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "01-01-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1609174800000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "01-02-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1611853200000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "09-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1630688400000,
			numberDay:      5,
			isStart:        true,
		},
		{
			datetimeStr:    "01-01-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609433999999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "01-01-2021 23:59:59.999",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609433999999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "01-01-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609433999999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "01-01-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1609347599999,
			numberDay:      2,
			isStart:        false,
		},
		{
			datetimeStr:    "01-02-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612112399999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "01-02-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1611939599999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "09-09-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1630947599999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "01-01-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1609261199999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "01-02-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1611939599999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "09-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1630774799999,
			numberDay:      5,
			isStart:        false,
		},
	}

	for _, table := range tables {
		// WHEN
		localTime, _ := ConvertStringToLocalTime(table.datetimeStr, table.format)
		actualMillis := ConvertLocalTimeToMilliseconds(GetBeforeLocalTimeOfTime(localTime, table.numberDay, table.isStart))

		// THEN
		assert.Equal(table.expectedMillis, actualMillis)
	}
}

func TestGetAfterLocalTimeOfTime_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		datetimeStr    string
		format         string
		expectedMillis int64
		numberDay      int
		isStart        bool
	}{
		{
			datetimeStr:    "31-12-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1640970000000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "31-12-2021 23:59:59.999",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1640970000000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "31-12-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1640970000000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "31-12-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1641056400000,
			numberDay:      2,
			isStart:        true,
		},
		{
			datetimeStr:    "31-01-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612112400000,
			numberDay:      1,
			isStart:        true,
		},
		{
			datetimeStr:    "31-01-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612285200000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "09-09-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1631379600000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "31-12-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1641142800000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "31-01-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1612285200000,
			numberDay:      3,
			isStart:        true,
		},
		{
			datetimeStr:    "09-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1631552400000,
			numberDay:      5,
			isStart:        true,
		},
		{
			datetimeStr:    "31-12-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1641056399999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "31-12-2021 23:59:59.999",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1641056399999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "31-12-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1641056399999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "31-12-2021 10:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1641142799999,
			numberDay:      2,
			isStart:        false,
		},
		{
			datetimeStr:    "31-01-2021 00:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612198799999,
			numberDay:      1,
			isStart:        false,
		},
		{
			datetimeStr:    "31-01-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1612371599999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "09-09-2021 20:00:00.000",
			format:         DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis: 1631465999999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "31-12-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1641229199999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "31-01-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1612371599999,
			numberDay:      3,
			isStart:        false,
		},
		{
			datetimeStr:    "09-09-2021",
			format:         DD_MM_YYYY,
			expectedMillis: 1631638799999,
			numberDay:      5,
			isStart:        false,
		},
	}

	for _, table := range tables {
		// WHEN
		localTime, _ := ConvertStringToLocalTime(table.datetimeStr, table.format)
		actualMillis := ConvertLocalTimeToMilliseconds(GetAfterLocalTimeOfTime(localTime, table.numberDay, table.isStart))

		// THEN
		assert.Equal(table.expectedMillis, actualMillis)
	}
}

func TestGetMillisecondsBetween_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		startDatetimeStr string
		endDatetimeStr   string
		format           string
		expectedMillis   int64
	}{
		{
			startDatetimeStr: "01-01-2021 09:09:09",
			endDatetimeStr:   "01-01-2021 09:09:09",
			format:           DD_MM_YYYY_HH_MM_SS,
			expectedMillis:   0,
		},
		{
			startDatetimeStr: "01-01-2021 09:09:09.009",
			endDatetimeStr:   "02-01-2021 23:59:59.999",
			format:           DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis:   139850990,
		},
		{
			startDatetimeStr: "01-01-2021 09:09:09.009",
			endDatetimeStr:   "01-02-2021 10:39:39.393",
			format:           DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis:   2683830384,
		},
		{
			startDatetimeStr: "19-01-2021 09:09:09.009",
			endDatetimeStr:   "16-01-2021 10:39:39.393",
			format:           DD_MM_YYYY_HH_MM_SS_SSS,
			expectedMillis:   -253769616,
		},
	}

	for _, table := range tables {
		// WHEN
		startLocatime, _ := ConvertStringToLocalTime(table.startDatetimeStr, table.format)
		endLocalTime, _ := ConvertStringToLocalTime(table.endDatetimeStr, table.format)

		millis := GetMillisecondsBetween(startLocatime, endLocalTime)

		// THEN
		assert.Equal(table.expectedMillis, millis)
	}
}

func BenchmarkGetCurrentMiliseconds(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetCurrentLocalTime()
		}
	})
}

func BenchmarkGetCurrentTime(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetCurrentLocalTime()
		}
	})
}

func BenchmarkConvertCurrentLocalTimeToString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ConvertCurrentLocalTimeToString(DD_MM_YYYY_HH_MM_SS_SSS)
		}
	})
}

func BenchmarkConvertLocalTimeToMilliseconds(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ConvertLocalTimeToMilliseconds(time.Now())
		}
	})
}

func BenchmarkConvertLocalTimeToString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ConvertLocalTimeToString(time.Now(), DD_MM_YYYY_HH_MM_SS_SSS)
		}
	})
}

func BenchmarkConvertStringToLocalTime(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ConvertStringToLocalTime("09-09-2021 09:09:09.999", DD_MM_YYYY_HH_MM_SS_SSS)
		}
	})
}

func BenchmarkConvertMillisecondsToLocalTime(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ConvertMillisecondsToLocalTime(1631417695897)
		}
	})
}
