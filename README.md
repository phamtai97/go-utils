# go-utils
[![Go Report Card](https://goreportcard.com/badge/github.com/phamtai97/go-utils)](https://goreportcard.com/report/github.com/phamtai97/go-utils) [![Go Reference](https://pkg.go.dev/badge/github.com/phamtai97/go-utils.svg)](https://pkg.go.dev/github.com/phamtai97/go-utils)

## Table of contents
- [go-utils](#go-utils)
  - [Table of contents](#table-of-contents)
  - [1. Overview](#1-overview)
  - [2. Install](#2-install)
  - [3. Utils package](#3-utils-package)
    - [3.1 logger](#31-logger)
    - [3.2 error](#32-error)
    - [3.3 datetime](#33-datetime)
    - [3.4 config](#34-config)
    - [3.5 conv](#35-conv)
## 1. Overview
In my free time, I will learn new knowledge about Golang and make notes on this project, or more simply, I will write my own components that can be reused for many different projects. This helped me review my knowledge of Golang as well as gain more experience on how to use this language.

## 2. Install
- Run command `go get`.

```sh
go get github.com/phamtai97/go-utils
```
## 3. Utils package
### [3.1 logger](./utils/logger/logger.go) 
- I have wrapped the [zap](https://github.com/uber-go/zap) library for easy use in projects. Why zap? Because it is very [fast](https://github.com/uber-go/zap#performance).
- How to use?
- We only need to create it once and use it anywhere in the project. For example:

```go
...
func main(){
    // write log to console
    // logger.InitProduction("")

    // write log to logs.log file
    if err := logger.InitProduction("./logs.log"); err != nil {
        fmt.Printf("Failed to init logger: %v\n", err)
    }
    defer logger.Sync()

    logger.Info("I am AJPham",
        zap.String("Hey, ", "I am a software engineer"),
        zap.Int("Age: ", 1997))
    logger.Error("I am AJPham",
        zap.String("Hey, ", "I am a software engineer"),
        zap.Int("Age: ", 1997))
}

// Result in file logs.log
// {"level":"INFO","ts":"2021-09-10 21:52:04.176","caller":"error/main.go:65","msg":"I am AJPham","Hey, ":"I am a software engineer","Age: ":1997}
// {"level":"ERROR","ts":"2021-09-10 21:52:04.176","caller":"error/main.go:69","msg":"I am AJPham","Hey, ":"I am a software engineer","Age: ":1997,"stacktrace":"main.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:69\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203"}
```

- Detailed examples can be see [here](cmd/logger/main.go).

### [3.2 error](./utils/error/error.go)
- This is a simple way to create errors in the project. It uses the [github.com/pkg/errors](https://github.com/pkg/errors) package as the core.
- How to use?
- We can new error and use it. For example:

```go
...
func main() {
    logger.InitProduction("")

    err := ero.New("Not found file")
    errA := err.AddStackTrace("Component A called")
    errB := errA.AddContextf("Component %s called", "B")

    if errB.Is(err) {
        logger.Info("Error B is err")
    }

    logger.Error("This is error wrapper", zap.Error(errB))
    logger.Error("This is error detail wrapper", zap.Error(errB.Detail()))
    logger.Error("This is root cause", zap.Error(errB.RootCause().Detail()))
    logger.Error("This is root cause", zap.String("Error string", err.Error()))
}

// Result
// {"level":"INFO","ts":"2021-09-10 21:55:29.840","caller":"error/main.go:66","msg":"Error B is err"}
// {"level":"ERROR","ts":"2021-09-10 21:50:15.523","caller":"error/main.go:65","msg":"This is error wrapper","error":"Component B called: Component A called: Not found file","stacktrace":"main.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:65\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203"}
// {"level":"ERROR","ts":"2021-09-10 21:50:15.523","caller":"error/main.go:66","msg":"This is error detail wrapper","error":"Component B called: Component A called: Not found file","errorVerbose":"Not found file\ngo-utils/utils/error.New\n\t/Users/Documents/github/go-utils/utils/error/error.go:17\nmain.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:61\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203\nruntime.goexit\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/asm_amd64.s:1357\nComponent A called\ngo-utils/utils/error.(*ErrorWrapper).AddStackTrace\n\t/Users/Documents/github/go-utils/utils/error/error.go:47\nmain.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:62\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203\nruntime.goexit\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/asm_amd64.s:1357\nComponent B called","stacktrace":"main.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:66\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203"}
// {"level":"ERROR","ts":"2021-09-10 21:50:15.523","caller":"error/main.go:67","msg":"This is root cause","error":"Not found file","errorVerbose":"Not found file\ngo-utils/utils/error.New\n\t/Users/Documents/github/go-utils/utils/error/error.go:17\nmain.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:61\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203\nruntime.goexit\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/asm_amd64.s:1357","stacktrace":"main.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:67\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203"}
// {"level":"ERROR","ts":"2021-09-10 21:50:15.523","caller":"error/main.go:68","msg":"This is root cause","Error string":"Not found file","stacktrace":"main.main\n\t/Users/Documents/github/go-utils/cmd/error/main.go:68\nruntime.main\n\t/usr/local/Cellar/go@1.13/1.13.11/libexec/src/runtime/proc.go:203"}
```

- Detailed examples can be see [here](cmd/error/main.go).

### [3.3 datetime](./utils/datetime/datetime.go)
- Working with Datetime in programming is inevitable. I provide a simple enough package to play with Datetime in Golang.
- How to use?
- It is easy.

```go
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
    currMillis := datetime.GetCurrentMiliseconds()
    logger.Info("Current milliseconds", zap.Int64("value", millis))

    // Convert milliseconds to specific string
    ddmmyyyy_hhmmss_sss := datetime.ConvertMillisecondsToString(currMillis, datetime.DD_MM_YYYY_HH_MM_SS_SSS)
    logger.Info("Convert milliseconds to format DD-MM-YYYY HH:mm:ss.SSS", zap.String("value", ddmmyyyy_hhmmss_sss))

    // Convert specific string to milliseconds
    millis, err := datetime.ConvertStringToMilliseconds("2021-09-09 09:09:09.999", datetime.YYYY_MM_DD_HH_MM_SS_SSS)
    if err != nil{
        logger.Error("Failed to convert", zap.Error(err))
    }
    logger.Info("Convert string format YYYY-MM-DD HH:mm:ss.SSS to millisecond", zap.Int64("value", millis))

    // other functions
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
```

- Detailed examples can be see [here](cmd/datetime/main.go).

### [3.4 config](./utils/config/config.go)
- Most applications need configuration to run (except very simple ones). We can manage configuration by file such as yaml, json file. The package provides a way to load configuration from yaml and json files and parse it into an object.
- How to use?
- Let's go.

```go
func main() {
    logger.InitProduction("")
    serviceConfig := ServiceConfig{}

    // Load config from yaml file
    if err := config.Load(&serviceConfig, "dev.yaml"); err != nil {
        logger.Fatal("Failed to load config", zap.Error(err))
    }

    // We can provide path of config by flag to load config
    if err := config.LoadByFlag(&serviceConfig, "cfgPath"); err != nil {
        logger.Fatal("Failed to load config", zap.Error(err))
    }

    // Load config from json file
    if err := config.Load(&serviceConfig, "dev.json"); err != nil {
        logger.Fatal("Failed to load config", zap.Error(err))
    }

    // If you want omit hotkeys such as token, password,...
    if err := config.Print(serviceConfig, "Token", "Password"); err != nil {
        logger.Fatal("Failed to print config", zap.Error(err))
    }
}
```
- Detailed examples can be see [here](./cmd/config/main.go).

### [3.5 conv](./utils/convertor/convertor.go)
- How to convert numbers and strings? How to convert string to number? It's simple because there's a convertor package.
- I use [strconv](https://pkg.go.dev/strconv) for conv package.
- It includes the following data types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool.
- How to use?
- Don't go to google looking for ways to convert anymore.

```go
func main() {
    logger.InitProduction("")

    numInt, err := conv.ConvertStringToInt("123456")
    if err != nil {
        logger.Fatal("Failed to convert string to int")
    }
    logger.Info("Convert string to int", zap.Int("Value", numInt))

    strInt := conv.ConvertIntToString(123456)
    logger.Info("Convert int to string", zap.String("Value", strUInt))
}
```

- Detailed examples can be see [here](./cmd/convertor/main.go).