package ero

import "github.com/pkg/errors"

type ErrorWrapper struct {
	err error
}

func Wrap(err error) *ErrorWrapper {
	return &ErrorWrapper{
		err: err,
	}
}

func New(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.New(message),
	}
}

func Newf(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Errorf(format, args...),
	}
}

func (e *ErrorWrapper) Error() string {
	return e.err.Error()
}

func (e *ErrorWrapper) Detail() error {
	return e.err
}

func (e *ErrorWrapper) RootCause() *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Cause(e.err),
	}
}

func (e *ErrorWrapper) RootCauseStr() string {
	return e.RootCause().Error()
}

func (e *ErrorWrapper) AddStackTrace(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Wrap(e.err, message),
	}
}

func (e *ErrorWrapper) AddStackTracef(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.Wrapf(e.err, format, args...),
	}
}

func (e *ErrorWrapper) AddContext(message string) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.WithMessage(e.err, message),
	}
}

func (e *ErrorWrapper) AddContextf(format string, args ...interface{}) *ErrorWrapper {
	return &ErrorWrapper{
		err: errors.WithMessagef(e.err, format, args...),
	}
}

func (e *ErrorWrapper) Is(target *ErrorWrapper) bool {
	return errors.Is(e.err, target.err)
}
