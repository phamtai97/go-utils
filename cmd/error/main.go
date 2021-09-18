// Package main contains examples of how to use the error package
package main

import (
	ero "github.com/phamtai97/go-utils/utils/error"
	"github.com/phamtai97/go-utils/utils/logger"

	"go.uber.org/zap"
)

func doError() error {
	return ero.New("Failed to open file")
}

func main() {
	logger.InitProduction("")

	// err := ero.Newf("Not found file")
	err := ero.New("Failed to open file")

	// Add context
	errA := err.AddContext("Call by component A")
	errB := errA.AddContext("Call by component B")
	errC := errB.AddContextf("Call by component C %d", 123)

	// Create error by pre-existing error
	errWrap := ero.Wrap(err)

	// Add stackstrace
	errWrapA := errWrap.AddStackTrace("Call by component wrapper A")
	errWrapB := errWrapA.AddStackTrace("Call by component wrapperB")
	errWrapC := errWrapB.AddStackTracef("Call by component wrapper C %d", 123)

	// Logger error
	logger.Error("Test", zap.Error(errC.Detail()))
	logger.Error("Test", zap.Error(errWrapC.Detail()))
	logger.Error("Test", zap.Error(errC))
	logger.Error("Test", zap.Error(errWrapC))
	logger.Error("Test", zap.String("Failed", errA.Error()))
	logger.Error("Test", zap.String("Failed", errWrapA.Error()))

	// Check error with Is
	if errC.Is(errA) {
		logger.Info("Error C is error A", zap.String("ErrorC", errC.Error()), zap.String("ErrorA", errA.Error()))
	}

	if !errWrap.Is(err) {
		logger.Info("errWrap is not err", zap.String("errWrap", errWrap.Error()), zap.String("err", err.Error()))
	}

	// Get root cause of error
	logger.Error("Root cause", zap.String("Root cause", errC.RootCause().Error()))
	logger.Error("Root cause", zap.String("Root cause", errC.RootCauseStr()))
	logger.Error("Root cause", zap.Error(errC.RootCause().Detail()))

	// Case error to ErrorWrapper
	errCast := doError().(*ero.ErrorWrapper)
	logger.Error("Cast error to ErrorWrapper", zap.Error(errCast.Detail()))
}
