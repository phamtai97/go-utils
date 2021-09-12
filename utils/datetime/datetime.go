package datetime

import (
	"time"
)

const (
	// YYYY_MM_DD_HH_MM_SS_SSS is the format "2006-01-02 15:04:05.000".
	YYYY_MM_DD_HH_MM_SS_SSS = "2006-01-02 15:04:05.000"
	// YYYY_MM_DD_HH_MM_SS is the format "2006-01-02 15:04:05".
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	// YYYY_MM_DD is the format "2006-01-02".
	YYYY_MM_DD = "2006-01-02"
	// DD_MM_YYYY is the format "02-01-2006".
	DD_MM_YYYY = "02-01-2006"
	// DD_MM_YYYY_HH_MM_SS is the format "02-01-2006 15:04:05".
	DD_MM_YYYY_HH_MM_SS = "02-01-2006 15:04:05"
	// DD_MM_YYYY_HH_MM_SS_SSS is the format "02-01-2006 15:04:05.000".
	DD_MM_YYYY_HH_MM_SS_SSS = "02-01-2006 15:04:05.000"
)

// GetCurrentLocalTime returns the current local time.
func GetCurrentLocalTime() time.Time {
	return time.Now()
}

// GetCurrentMiliseconds returns the current milliseconds.
func GetCurrentMiliseconds() int64 {
	return ConvertLocalTimeToMilliseconds(GetCurrentLocalTime())
}

// ConvertCurrentLocalTimeToString converts the current local time to string with the specific format.
func ConvertCurrentLocalTimeToString(format string) string {
	return GetCurrentLocalTime().Format(format)
}

// ConvertMillisecondsToString converts the milliseconds to the string with the specific format.
func ConvertMillisecondsToString(millis int64, format string) string {
	return ConvertMillisecondsToLocalTime(millis).Format(format)
}

// ConvertStringToMilliseconds converts the string with specific format to milliseconds.
func ConvertStringToMilliseconds(value, format string) (int64, error) {
	localTime, err := time.ParseInLocation(format, value, time.Local)
	if err != nil {
		return -1, err
	}

	return ConvertLocalTimeToMilliseconds(localTime), nil
}

// ConvertStringToLocalTime converts the string with specific format to the local time.
func ConvertStringToLocalTime(value, format string) (time.Time, error) {
	return time.ParseInLocation(format, value, time.Local)
}

// ConvertLocalTimeToMilliseconds converts the local time to milliseconds.
func ConvertLocalTimeToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// ConvertMillisecondsToLocalTime converts the milliseconds to local time.
func ConvertMillisecondsToLocalTime(millis int64) time.Time {
	return time.Unix(millis/1e3, (millis%1e3)*1e6)
}

// ConvertLocalTimeToString converts the local time to string with specific format.
func ConvertLocalTimeToString(localTime time.Time, format string) string {
	return localTime.Format(format)
}

// GetYear returns the current year.
func GetYear() int {
	return GetCurrentLocalTime().Year()
}

// GetDayOfYear returns the day of the year.
func GetDayOfYear() int {
	return GetCurrentLocalTime().YearDay()
}

// GetDayOfMonth returns the day of month.
func GetDayOfMonth() int {
	return GetCurrentLocalTime().Day()
}

// GetMonthOfYear returns the month of year. Value from 1 to 12.
func GetMonthOfYear() int {
	return int(GetCurrentLocalTime().Month())

}

// GetStartLocalTimeOfYear returns the start local time of year.
func GetStartLocalTimeOfYear() time.Time {
	return time.Date(GetYear(), 1, 1, 0, 0, 0, 0, time.Local)
}

// GetEndLocalTimeOfYear returns the end local time of year.
func GetEndLocalTimeOfYear() time.Time {
	return time.Date(GetYear(), 12, 31, 23, 59, 59, 999999999, time.Local)
}

// GetStartLocalTimeOfMonth returns the start local time of month.
func GetStartLocalTimeOfMonth() time.Time {
	return time.Date(GetYear(), GetCurrentLocalTime().Month(), 1, 0, 0, 0, 0, time.Local)
}

// GetEndLocalTimeOfMonth returns the end local time of month.
func GetEndLocalTimeOfMonth() time.Time {
	return time.Date(GetYear(), GetCurrentLocalTime().Month(), 1, 23, 59, 59, 999999999, time.Local).AddDate(0, 1, -1)
}

// GetStartLocalTimeOfDay return the start local time of day.
func GetStartLocalTimeOfDay() time.Time {
	return time.Date(GetYear(), GetCurrentLocalTime().Month(), GetCurrentLocalTime().Day(), 0, 0, 0, 0, time.Local)
}

// GetEndLocalTimeOfDay return the end local time of day.
func GetEndLocalTimeOfDay() time.Time {
	return time.Date(GetYear(), GetCurrentLocalTime().Month(), GetCurrentLocalTime().Day(), 23, 59, 59, 999999999, time.Local)
}

// GetStartLocalTimeOfTime return the start local time of specific local time.
func GetStartLocalTimeOfTime(inputTime time.Time) time.Time {
	return time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 0, 0, 0, 0, time.Local)
}

// GetEndLocalTimeOfTime return the end local time of specific local time.
func GetEndLocalTimeOfTime(inputTime time.Time) time.Time {
	return time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 23, 59, 59, 999999999, time.Local)
}

// GetBeforeLocalTimeOfTime returns before local time with the number of days compared to specific local time.
func GetBeforeLocalTimeOfTime(inputTime time.Time, numberDay int, isStartTime bool) time.Time {
	if isStartTime {
		return GetStartLocalTimeOfTime(inputTime.AddDate(0, 0, -numberDay))
	}

	return GetEndLocalTimeOfTime(inputTime.AddDate(0, 0, -numberDay))
}

// GetAfterLocalTimeOfTime returns after local time with the number of days compared to specific local time.
func GetAfterLocalTimeOfTime(inputTime time.Time, numberDay int, isStartTime bool) time.Time {
	if isStartTime {
		return GetStartLocalTimeOfTime(inputTime.AddDate(0, 0, numberDay))
	}

	return GetEndLocalTimeOfTime(inputTime.AddDate(0, 0, numberDay))
}

// GetMillisecondsBetween returns the number of milliseconds between 2 local time.
func GetMillisecondsBetween(startDate time.Time, endTime time.Time) int64 {
	return endTime.Sub(startDate).Milliseconds()
}
