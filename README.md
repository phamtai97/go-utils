# go-utils
- [go-utils](#go-utils)
  - [1. Overview](#1-overview)
  - [2. Utils package](#2-utils-package)
    - [2.1 logger](#21-logger)
## 1. Overview
In my free time, I will learn new knowledge about Golang and make notes on this project, or more simply, I will write my own components that can be reused for many different projects. This helped me review my knowledge of Golang as well as gain more experience on how to use this language.

## 2. Utils package
### 2.1 [logger](./utils/logger/logger.go) 
- I have wrapped the [zap](https://github.com/uber-go/zap) library for easy use in projects. Why zap? Because it is very [fast](https://github.com/uber-go/zap#performance).
- How to use? 
- Fisrtly, we install zap package.

```sh
# Install zap package
go get -u go.uber.org/zap
```

- Then, we only need to create it once and use it anywhere in the project. For example:

```go
...
func main(){
    // write log to console
    // logger.InitProductionLogger("")

    // write log to logs.log file
    if err := logger.InitProductionLogger("./logs.log"); err != nil {
        fmt.Printf("Failed to init logger: %v\n", err)
    }
    defer logger.Sync()

    logger.Info("I am AJPham",
        zap.String("Hey, ", "I am a software engineer"),
        zap.Int("Age: ", 1997))
}

// Result in file logs.log
// {"level":"INFO","ts":"2021-09-08 21:48:16.473","caller":"logger/main.go:53","msg":"I am AJPham","Hey, ":"I am a software engineer","Age: ":1997}
```
- Detailed examples can be see [here](cmd/logger/main.go).