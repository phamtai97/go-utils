// Package main contains examples of how to use the conv package
package main

import (
	conv "github.com/phamtai97/go-utils/utils/convertor"
	"github.com/phamtai97/go-utils/utils/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitProduction("")

	numInt, err := conv.ConvertStringToInt("123456")
	if err != nil {
		logger.Fatal("Failed to convert string to int")
	}
	logger.Info("Convert string to int", zap.Int("Value", numInt))

	numInt32, err := conv.ConvertStringToInt32("123456")
	if err != nil {
		logger.Fatal("Failed to convert string to int32")
	}
	logger.Info("Convert string to int32", zap.Int32("Value", numInt32))

	numInt64, err := conv.ConvertStringToInt64("123456")
	if err != nil {
		logger.Fatal("Failed to convert string to int64")
	}
	logger.Info("Convert string to int64", zap.Int64("Value", numInt64))

	bTrue, err := conv.ConvertStringToBool("TRUE")
	if err != nil {
		logger.Fatal("Failed to convert string to bool")
	}
	logger.Info("Convert string to bool", zap.Bool("Value", bTrue))

	bFalse, err := conv.ConvertStringToBool("false")
	if err != nil {
		logger.Fatal("Failed to convert string to bool")
	}
	logger.Info("Convert string to bool", zap.Bool("Value", bFalse))

	numFloat32, err := conv.ConvertStringToFloat32("1.23243")
	if err != nil {
		logger.Fatal("Failed to convert string to float32")
	}
	logger.Info("Convert string to float32", zap.Float32("Value", numFloat32))

	numFloat64, err := conv.ConvertStringToFloat64("1.23243")
	if err != nil {
		logger.Fatal("Failed to convert string to float64")
	}
	logger.Info("Convert string to float64", zap.Float64("Value", numFloat64))

	numUInt, err := conv.ConvertStringToUInt("123456")
	if err != nil {
		logger.Fatal("Failed to convert string to uint")
	}
	logger.Info("Convert string to uint", zap.Uint("Value", numUInt))

	strInt := conv.ConvertIntToString(123456)
	logger.Info("Convert int to string", zap.String("Value", strInt))

	strInt32 := conv.ConvertInt32ToString(123456)
	logger.Info("Convert int32 to string", zap.String("Value", strInt32))

	strInt64 := conv.ConvertInt64ToString(123456)
	logger.Info("Convert int64 to string", zap.String("Value", strInt64))

	strBool := conv.ConvertBoolToString(true)
	logger.Info("Convert bool to string", zap.String("Value", strBool))

	strFloat32 := conv.ConvertFloat32ToString(1.2345)
	logger.Info("Convert float32 to string", zap.String("Value", strFloat32))

	strFloat64 := conv.ConvertFloat64ToString(1.2345)
	logger.Info("Convert float64 to string", zap.String("Value", strFloat64))

	strUInt := conv.ConvertUIntToString(123456)
	logger.Info("Convert uint to string", zap.String("Value", strUInt))
}
