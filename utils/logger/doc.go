// The package provides a way to create logger. We only need to create it once and use it anywhere in the project.
//
// For example Usage
//
// 	import (
// 		"errors"
// 		"fmt"
//
// 		"github.com/phamtai97/go-utils/utils/logger"
//
// 		"go.uber.org/zap"
// 		"go.uber.org/zap/zapcore"
// 	)
//
// 	type user struct {
// 		Username string
// 		Age      int
// 	}
//
// 	func (u *user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
// 		enc.AddString("name", u.Username)
// 		enc.AddInt("age", u.Age)
// 		return nil
// 	}
//
// 	func main() {
// 		// Can customize the logger config
// 		// cfg := logger.Config{
// 		// 	Level: logger.INFO,
// 		// 	FileLogConfig: logger.FileLogConfig{
// 		// 		IsUseFile: true,
// 		// 		FilePath:  "./logs.log",
// 		// 	},
// 		// }
//
// 		// New default config
// 		// cfg := logger.NewDefaultConfig()
//
// 		// New production config, write logs to file
// 		// cfg := logger.NewProductionConfig(true, "./logs.log")
//
// 		// New production config, write logs to console
// 		// cfg := logger.NewProductionConfig("")
//
// 		// Can init logger with simple line
// 		// logger.InitProduction("./logs.log")
//
// 		cfg := logger.Config{
// 			Level: logger.DEBUG,
// 		}
//
// 		if err := logger.Init(cfg); err != nil {
// 			fmt.Printf("Failed to init logger: %v\n", err)
// 		}
// 		defer logger.Sync()
//
// 		logger.Debug("Test debug logger",
// 			zap.String("Hey, ", "I am a software engineer"),
// 			zap.Object("My information: ", &user{
// 				Username: "AJPham",
// 				Age:      1997,
// 			}))
//
// 		logger.Info("Test info logger",
// 			zap.String("Hey, ", "I am a software engineer"),
// 			zap.Object("My information: ", &user{
// 				Username: "AJPham",
// 				Age:      1997,
// 			}))
//
// 		logger.Error("Test error logger",
// 			zap.String("Hey, ", "I am a software engineer"),
// 			zap.Object("My information: ", &user{
// 				Username: "AJPham",
// 				Age:      1997,
// 			}), zap.Error(errors.New("Failed to write log")))
//
// 		logger.Warn("Test warn logger",
// 			zap.String("Hey, ", "I am a software engineer"),
// 			zap.Object("My information: ", &user{
// 				Username: "AJPham",
// 				Age:      1997,
// 			}))
//
// 		logger.Fatal("Test fatal logger",
// 			zap.String("Hey, ", "I am a software engineer"),
// 			zap.Object("My information: ", &user{
// 				Username: "AJPham",
// 				Age:      1997,
// 			}))
// 	}
package logger
