// Package main contains examples of how to use the datetime package
package main

import (
	"time"

	"github.com/phamtai97/go-utils/utils/datetime"
	"github.com/phamtai97/go-utils/utils/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitProduction("")

	// Convert current milliseconds to different formats
	logger.Info("Convert current milliseconds to format YYYY-MM-DD", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.YYYY_MM_DD)))
	logger.Info("Convert current milliseconds to format YYYY-MM-DD HH:mm:ss", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.YYYY_MM_DD_HH_MM_SS)))
	logger.Info("Convert current milliseconds to format YYYY-MM-DD HH:mm:ss.SSS", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.YYYY_MM_DD_HH_MM_SS_SSS)))
	logger.Info("Convert current milliseconds to format DD-MM-YYYY", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.DD_MM_YYYY)))
	logger.Info("Convert current milliseconds to format DD-MM-YYYY HH:mm:ss", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.DD_MM_YYYY_HH_MM_SS)))
	logger.Info("Convert current milliseconds to format DD-MM-YYYY HH:mm:ss.SSS", zap.String("value", datetime.ConvertCurrentLocalTimeToString(datetime.DD_MM_YYYY_HH_MM_SS_SSS)))

	// Get current millisenconds
	millis := datetime.GetCurrentMiliseconds()
	logger.Info("Current milliseconds", zap.Int64("value", millis))

	// Convert milliseconds to different formats
	yyyymmdd := datetime.ConvertMillisecondsToString(millis, datetime.YYYY_MM_DD)
	logger.Info("Convert milliseconds to format YYYY-MM-DD", zap.String("value", yyyymmdd))

	yyyymmddHhmmss := datetime.ConvertMillisecondsToString(millis, datetime.YYYY_MM_DD_HH_MM_SS)
	logger.Info("Convert milliseconds to format YYYY-MM-DD HH:mm:ss", zap.String("value", yyyymmddHhmmss))

	yyyymmddHhmmssSss := datetime.ConvertMillisecondsToString(millis, datetime.YYYY_MM_DD_HH_MM_SS_SSS)
	logger.Info("Convert milliseconds to format YYYY-MM-DD HH:mm:ss.SSS", zap.String("value", yyyymmddHhmmssSss))

	ddmmyyyy := datetime.ConvertMillisecondsToString(millis, datetime.DD_MM_YYYY)
	logger.Info("Convert milliseconds to format DD-MM-YYYY", zap.String("value", ddmmyyyy))

	ddmmyyyyHhmmss := datetime.ConvertMillisecondsToString(millis, datetime.DD_MM_YYYY_HH_MM_SS)
	logger.Info("Convert milliseconds to format DD-MM-YYYY HH:mm:ss", zap.String("value", ddmmyyyyHhmmss))

	ddmmyyyyHhmmssSss := datetime.ConvertMillisecondsToString(millis, datetime.DD_MM_YYYY_HH_MM_SS_SSS)
	logger.Info("Convert milliseconds to format DD-MM-YYYY HH:mm:ss.SSS", zap.String("value", ddmmyyyyHhmmssSss))

	// Convert string format to millisecond
	millis, _ = datetime.ConvertStringToMilliseconds(yyyymmdd, datetime.YYYY_MM_DD)
	logger.Info("Convert string format YYYY-MM-DD to millisecond", zap.Int64("value", millis))

	millis, _ = datetime.ConvertStringToMilliseconds(yyyymmddHhmmss, datetime.YYYY_MM_DD_HH_MM_SS)
	logger.Info("Convert string format YYYY-MM-DD HH:mm:ss to millisecond", zap.Int64("value", millis))

	millis, _ = datetime.ConvertStringToMilliseconds(yyyymmddHhmmssSss, datetime.YYYY_MM_DD_HH_MM_SS_SSS)
	logger.Info("Convert string format YYYY-MM-DD HH:mm:ss.SSS to millisecond", zap.Int64("value", millis))

	millis, _ = datetime.ConvertStringToMilliseconds(ddmmyyyy, datetime.DD_MM_YYYY)
	logger.Info("Convert string format DD-MM-YYYY to millisecond", zap.Int64("value", millis))

	millis, _ = datetime.ConvertStringToMilliseconds(ddmmyyyyHhmmss, datetime.DD_MM_YYYY_HH_MM_SS)
	logger.Info("Convert string format DD-MM-YYYY HH:mm:ss to millisecond", zap.Int64("value", millis))

	millis, _ = datetime.ConvertStringToMilliseconds(ddmmyyyyHhmmssSss, datetime.DD_MM_YYYY_HH_MM_SS_SSS)
	logger.Info("Convert string format DD-MM-YYYY HH:mm:ss.SSS to millisecond", zap.Int64("value", millis))

	// other functions
	logger.Info("Year", zap.Int("value", datetime.GetYear()))
	logger.Info("Day of year", zap.Int("value", datetime.GetDayOfYear()))
	logger.Info("Day of month", zap.Int("value", datetime.GetDayOfMonth()))
	logger.Info("Month of year", zap.Int("value", datetime.GetMonthOfYear()))
	logger.Info("Start local time of year", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetStartLocalTimeOfYear())))
	logger.Info("End local time of year", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetEndLocalTimeOfYear())))
	logger.Info("Start local time of month", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetStartLocalTimeOfMonth())))
	logger.Info("End local time of month", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetEndLocalTimeOfMonth())))
	logger.Info("Start local time of day", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetStartLocalTimeOfDay())))
	logger.Info("End local time of day", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetEndLocalTimeOfDay())))
	logger.Info("Start local time of time", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetStartLocalTimeOfTime(time.Now()))))
	logger.Info("Start local time of time", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetEndLocalTimeOfTime(time.Now()))))
	logger.Info("Get before local time of time", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetBeforeLocalTimeOfTime(time.Now(), 9, true))))
	logger.Info("Get after local time of time", zap.Int64("value", datetime.ConvertLocalTimeToMilliseconds(datetime.GetAfterLocalTimeOfTime(time.Now(), 9, false))))
}
