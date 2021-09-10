package ero

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorWrapper_WrapError_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	err := errors.New("Failed to open file")

	// WHEN
	errOpenFile := Wrap(err)

	// THEN
	assert.NotNil(errOpenFile)
	assert.NotNil(errOpenFile.Detail())
	assert.Equal("Failed to open file", errOpenFile.Error())
}

func TestNewErrorWrapper_MessageString_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	errOpenFile := New("Failed to open file")

	// THEN
	assert.NotNil(errOpenFile)
	assert.NotNil(errOpenFile.Detail())
	assert.Equal("Failed to open file", errOpenFile.Error())
}

func TestNewErrorWrapperf_MessageString_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	errOpenFile := Newf("Failed to open file=%s", "test.csv")

	// THEN
	assert.NotNil(errOpenFile)
	assert.NotNil(errOpenFile.Detail())
	assert.Equal("Failed to open file=test.csv", errOpenFile.Error())
}

func TestGetError_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		actualErr   string
		expectedErr string
	}{
		{
			actualErr:   "Failed to open file",
			expectedErr: "Failed to open file",
		},
		{
			actualErr:   "",
			expectedErr: "",
		},
		{
			actualErr:   "abc~!@#$%^&*()_+-=:;'[]<>,.?/1234",
			expectedErr: "abc~!@#$%^&*()_+-=:;'[]<>,.?/1234",
		},
	}

	for _, table := range tables {
		// WHEN
		err := New(table.actualErr)

		// THEN
		assert.NotNil(err)
		assert.NotNil(err.Detail())
		assert.Equal(table.expectedErr, err.Error())
	}
}

func TestGetErrorf_MultipleCase_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	tables := []struct {
		actualErr   string
		formatStr   string
		formatInt   int
		expectedErr string
	}{
		{
			actualErr:   "Failed to open file=%s code=%d",
			formatStr:   "system.go",
			formatInt:   8080,
			expectedErr: "Failed to open file=system.go code=8080",
		},
		{
			actualErr:   "",
			expectedErr: "%!(EXTRA string=, int=0)",
		},
		{
			actualErr:   "Failed to open file=%s code=%d",
			formatStr:   "abc~!@#$%^&*()_+-=:;'[]<>,.?/1234",
			formatInt:   -8080,
			expectedErr: "Failed to open file=abc~!@#$%^&*()_+-=:;'[]<>,.?/1234 code=-8080",
		},
	}

	for _, table := range tables {
		// WHEN
		err := Newf(table.actualErr, table.formatStr, table.formatInt)

		// THEN
		assert.NotNil(err)
		assert.NotNil(err.Detail())
		assert.Equal(table.expectedErr, err.Error())
	}
}

func TestAddStackTrace_SimpleInput_Success(t *testing.T) {
	//GIVEN
	assert := assert.New(t)

	// WHEN
	err := New("Failed to open file")
	errA := err.AddStackTrace("Component A called")
	errB := errA.AddStackTrace("Component B called")

	// THEN
	assert.NotNil(err)
	assert.NotNil(errA)
	assert.NotNil(errB)
	assert.Equal("Failed to open file", err.Error())
	assert.Equal("Component A called: Failed to open file", errA.Error())
	assert.Equal("Component B called: Component A called: Failed to open file", errB.Error())
}

func TestAddStackTracef_SimpleInput_Success(t *testing.T) {
	//GIVEN
	assert := assert.New(t)

	// WHEN
	err := New("Failed to open file")
	errA := err.AddStackTracef("Component=%s called code=%d", "A", 123)
	errB := errA.AddStackTracef("Component=%s called code=%d", "B", 456)

	// THEN
	assert.NotNil(err)
	assert.NotNil(errA)
	assert.NotNil(errB)
	assert.Equal("Failed to open file", err.Error())
	assert.Equal("Component=A called code=123: Failed to open file", errA.Error())
	assert.Equal("Component=B called code=456: Component=A called code=123: Failed to open file", errB.Error())
}

func TestRootCause_SimpleInput_Success(t *testing.T) {
	//GIVEN
	assert := assert.New(t)

	// WHEN
	err := New("Failed to open file")
	errA := err.AddStackTracef("Component=%s called", "A")
	errB := errA.AddStackTracef("Component=%s called", "B")

	// THEN
	assert.NotNil(err)
	assert.NotNil(errA)
	assert.NotNil(errB)
	assert.NotNil(errA.RootCause())
	assert.NotNil(errA.RootCause().Detail())
	assert.NotNil(errB.RootCause())
	assert.NotNil(errB.RootCause().Detail())

	assert.Equal("Failed to open file", err.RootCause().Error())
	assert.Equal("Failed to open file", errB.RootCause().Error())
	assert.Equal("Failed to open file", errA.RootCause().Error())

	assert.Equal("Failed to open file", err.RootCauseStr())
	assert.Equal("Failed to open file", err.RootCauseStr())
	assert.Equal("Failed to open file", err.RootCauseStr())
}

func TestAddContext_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	err := New("Failed to open file")
	errA := err.AddContext("Component A called")
	errB := errA.AddContext("Component B called")

	// THEN
	assert.NotNil(err)
	assert.NotNil(errA)
	assert.NotNil(errB)
	assert.Equal("Failed to open file", err.Error())
	assert.Equal("Component A called: Failed to open file", errA.Error())
	assert.Equal("Component B called: Component A called: Failed to open file", errB.Error())
}

func TestAddContextf_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	err := New("Failed to open file")
	errA := err.AddContextf("Component=%s called code=%d", "A", 123)
	errB := errA.AddContextf("Component=%s called code=%d", "B", 456)

	// THEN
	assert.NotNil(err)
	assert.NotNil(errA)
	assert.NotNil(errB)
	assert.Equal("Failed to open file", err.Error())
	assert.Equal("Component=A called code=123: Failed to open file", errA.Error())
	assert.Equal("Component=B called code=456: Component=A called code=123: Failed to open file", errB.Error())
}

func TestIs_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	errA := New("Test Is func")
	errB := errA.AddStackTrace("Component B called")
	errC := errB.AddContext("Component C called")

	isCheckBA := errB.Is(errA)
	isCheckCA := errC.Is(errA)
	isCheckCB := errC.Is(errB)

	isCheckAB := errA.Is(errB)
	isCheckAC := errA.Is(errC)
	isCheckBC := errB.Is(errC)

	// THEN
	assert.True(isCheckBA)
	assert.True(isCheckCA)
	assert.True(isCheckCB)

	assert.False(isCheckAB)
	assert.False(isCheckAC)
	assert.False(isCheckBC)
}

func TestIs_DifferentError_NotIs(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	errA := New("Test Is func")
	errB := New("Test Is func")

	isCheck := errB.Is(errA)

	// THEN
	assert.False(isCheck)
}

func BenchmarkCreateErrorWrapper(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New("Benchmark create error")
		}
	})
}

func BenchmarkCreateErrorWrapperf(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Newf("Benchmark create error=%s", "AJPham")
		}
	})
}

func BenchmarkAddStackTrace(b *testing.B) {
	b.ResetTimer()
	err := New("Benchmark create error")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err.AddStackTrace("Component benchmark")
		}
	})
}

func BenchmarkAddContext(b *testing.B) {
	b.ResetTimer()
	err := New("Benchmark create error")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err.AddContext("Component benchmark")
		}
	})
}

func BenchmarkGetRootCause(b *testing.B) {
	b.ResetTimer()
	err := New("Benchmark create error")
	errA := err.AddStackTrace("Component A called")
	errB := errA.AddStackTrace("Component B called")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			errB.RootCause()
		}
	})
}

func BenchmarkGetDetail(b *testing.B) {
	b.ResetTimer()
	err := New("Benchmark create error")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err.Detail()
		}
	})
}

func BenchmarkIs(b *testing.B) {
	b.ResetTimer()
	errA := New("Benchmark A create error")
	errB := New("Benchmark B create error")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			errA.Is(errB)
		}
	})
}
