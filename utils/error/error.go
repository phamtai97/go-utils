package ero

import "github.com/pkg/errors"

// ErrorWrapper wraps any error for ease of use.
type ErrorWrapper struct {
	err error
}

// Wrap returns the *ErrorWrapper with exist error.
func Wrap(err error) *ErrorWrapper {
	return &ErrorWrapper{
		err: err,
	}
}

// New returns the *ErrorWrapper with message.
func New(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.New(message),
	}
}

// Newf returns the *ErrorWrapper with formated message.
func Newf(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Errorf(format, args...),
	}
}

// Error returns the error string.
func (e *ErrorWrapper) Error() string {
	return e.err.Error()
}

// Detail return the error in the ErrorWrapper.
func (e *ErrorWrapper) Detail() error {
	return e.err
}

// RootCause returns *ErrorWrapper that contains the root cause of error.
func (e *ErrorWrapper) RootCause() *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Cause(e.err),
	}
}

// RootCauseStr returns the root cause string.
func (e *ErrorWrapper) RootCauseStr() string {
	return e.RootCause().Error()
}

// AddStackTrace returns *ErrorWrapper containing error has been added stackstrace.
func (e *ErrorWrapper) AddStackTrace(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Wrap(e.err, message),
	}
}

// AddStackTracef returns *ErrorWrapper containing error has been added stackstrace with format message.
func (e *ErrorWrapper) AddStackTracef(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Wrapf(e.err, format, args...),
	}
}

// AddContext returns *ErrorWrapper containing error has been added context.
func (e *ErrorWrapper) AddContext(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.WithMessage(e.err, message),
	}
}

// AddContextf returns *ErrorWrapper containing error has been added context with format message.
func (e *ErrorWrapper) AddContextf(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.WithMessagef(e.err, format, args...),
	}
}

// Is checks current error is targer error.
func (e *ErrorWrapper) Is(target *ErrorWrapper) bool {
	return errors.Is(e.err, target.err)
}
