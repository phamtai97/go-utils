package logger

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewDefaultConfig_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	cfg := NewDefaultConfig()

	// THEN
	assert.NotNil(cfg)
	assert.Equal(INFO, cfg.Level)

	fileLogCfg := cfg.FileLogConfig
	assert.NotNil(fileLogCfg)
	assert.Equal(false, fileLogCfg.IsUseFile)
	assert.Equal("", fileLogCfg.FilePath)
	assert.Equal(0, fileLogCfg.MaxAge)
	assert.Equal(0, fileLogCfg.MaxBackups)
	assert.Equal(0, fileLogCfg.MaxSize)
	assert.Equal(false, fileLogCfg.Compress)
}

func TestNewProductionConfig_MultipleCase_Success(t *testing.T) {
	// GIVEN
	tables := []struct {
		isUseFile bool
		filePath  string
	}{
		{false, ""},
		{false, "./abc.log"},
		{true, ""},
		{true, "./abc.log"},
	}
	assert := assert.New(t)

	for _, table := range tables {
		// WHEN
		cfg := NewProductionConfig(table.isUseFile, table.filePath)

		// THEN
		assert.NotNil(cfg)
		assert.Equal(INFO, cfg.Level)

		fileLogCfg := cfg.FileLogConfig
		assert.NotNil(fileLogCfg)
		assert.Equal(table.isUseFile, fileLogCfg.IsUseFile)
		assert.Equal(table.filePath, fileLogCfg.FilePath)
		assert.Equal(0, fileLogCfg.MaxAge)
		assert.Equal(0, fileLogCfg.MaxBackups)
		assert.Equal(DefaultLogFileSizeInMB, fileLogCfg.MaxSize)
		assert.Equal(true, fileLogCfg.Compress)
	}
}

func TestInitProductionLogger_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	err := InitProductionLogger("./logger.log")

	// THEN
	assert.Nil(err)
}
func TestInitLogger_MultipleCase_Success(t *testing.T) {
	// GIVEN
	tables := []struct {
		config        Config
		isSuccess     bool
		expectedError string
	}{
		{
			Config{
				Level: INFO,
				FileLogConfig: FileLogConfig{
					IsUseFile: false,
				},
			}, true, ""},
		{
			Config{
				Level: "INFOS",
				FileLogConfig: FileLogConfig{
					IsUseFile: false,
				},
			}, true, ""},
		{
			Config{
				Level: ERROR,
				FileLogConfig: FileLogConfig{
					IsUseFile: true,
					FilePath:  "./abc.log",
				},
			}, true, ""},
		{
			Config{
				Level: "",
				FileLogConfig: FileLogConfig{
					IsUseFile: false,
				},
			}, false, "Missing level logger"},
		{
			Config{
				Level: WARN,
				FileLogConfig: FileLogConfig{
					IsUseFile: true,
				},
			}, false, "File path must be not empty"},
		{
			Config{
				Level: FATAL,
				FileLogConfig: FileLogConfig{
					IsUseFile: false,
					MaxAge:    -1,
				},
			}, false, "MaxAge must be greater than or equal to 0"},
		{
			Config{
				Level: INFO,
				FileLogConfig: FileLogConfig{
					IsUseFile:  false,
					MaxBackups: -1,
				},
			}, false, "MaxBackups must be greater than or equal to 0"},
		{
			Config{
				Level: INFO,
				FileLogConfig: FileLogConfig{
					IsUseFile: false,
					MaxSize:   -1,
				},
			}, false, "MaxSize must be greater than or equal to 0"},
		{
			Config{
				Level: INFO,
				FileLogConfig: FileLogConfig{
					IsUseFile: true,
					FilePath:  "../logger",
				},
			}, false, "File path is invalid"},
	}
	assert := assert.New(t)

	for _, table := range tables {
		// WHEN
		err := InitLogger(table.config)

		// THEN
		if table.isSuccess {
			assert.Equal(nil, err)
		} else {
			assert.Equal(table.expectedError, err.Error())
		}
	}
}

type user struct {
	Username string
	Age      int
}

func (u *user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Username)
	enc.AddInt("email", u.Age)
	return nil
}

func TestLoggerDebug_MultipleCase_Success(t *testing.T) {
	// GIVEN
	fields := []zap.Field{
		zap.String("abc@xyz", "xyz@abc"),
		zap.Strings("abc@xyz", []string{"123@abc", "456@abc 789@", "123", "~!@#$%^&*(){}:;<>?,.-_=/``+"}),
		zap.Uint("Uint_id", 123456),
		zap.Uints("Uint_Ids", []uint{123, 456, 789}),
		zap.Int("Int_Id", 123456),
		zap.Ints("Int_Ids", []int{123, -456, 789}),
		zap.Uint32("Uint32_Id", 123456),
		zap.Uint32s("Uint32_Ids", []uint32{123, 456, 789}),
		zap.Int32("Int32_Id", 123456),
		zap.Int32s("Int32_Ids", []int32{123, -456, 789}),
		zap.Uint64("Uint64_Id", 123456),
		zap.Uint64s("Uint64_Ids", []uint64{123, 456, 789}),
		zap.Int64("Int64_Id", 123456),
		zap.Int64s("Int64_Ids", []int64{123, -456, 789}),
		zap.Object("Username", &user{
			Username: "AJPham",
			Age:      1997,
		}),
		zap.Binary("Binary", []byte("abc@123")),
		zap.Bool("IsTrue", true),
		zap.Bools("Bools", []bool{true, false, true, false}),
		zap.ByteString("Bytestring", []byte("a,b,c,1,2,3")),
		zap.ByteString("Bytestring", []byte("a,b,c,1,2,3")),
		zap.Reflect("Reflect", []int{1, 2, 3, 4, 5}),
		zap.Duration("duration", 10*time.Second),
	}

	assert := assert.New(t)
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	observedLogger := zap.New(observedZapCore)

	// WHEN
	setGlobalLog(observedLogger)
	defer Sync()
	Debug("Test debug level", fields...)
	Info("Test info level", fields...)
	Warn("Test warn level", fields...)
	Error("Test error level", fields...)

	// THEN
	assert.Equal(4, observedLogs.Len())
	assert.Equal("Test debug level", observedLogs.All()[0].Message)
	assert.Equal("Test info level", observedLogs.All()[1].Message)
	assert.Equal("Test warn level", observedLogs.All()[2].Message)
	assert.Equal("Test error level", observedLogs.All()[3].Message)
	for _, entry := range observedLogs.All() {
		assert.ElementsMatch(fields, entry.Context)
	}
}

func TestLogger_Parallel_Success(t *testing.T) {
	// GIVEN
	var wg sync.WaitGroup
	assert := assert.New(t)
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	observedLogger := zap.New(observedZapCore)

	// WHEN
	setGlobalLog(observedLogger)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Debug("Test parallel logger")
		}()

	}
	wg.Wait()

	// THEN
	assert.Equal(100, observedLogs.Len())
	for _, entry := range observedLogs.All() {
		assert.Equal("Test parallel logger", entry.Message)
	}
}

func BenchmarkLoggerInfo(b *testing.B) {
	InitLogger(NewDefaultConfig())
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("Test Benchmark", zap.String("abc@xyz", "xyz@abc"),
				zap.Int64("Int64_Id", 123456),
				zap.Object("Username", &user{
					Username: "AJPham",
					Age:      1997,
				}))
		}
	})
}
