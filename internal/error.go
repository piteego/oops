package internal

import (
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type Error[T any] struct {
	Msg  string
	Data T
}

func (err *Error[_]) Error() string { return err.Msg }

func (err *Error[_]) Kind() (kind.Code, bool) {
	switch d := any(err.Data).(type) {
	case kind.Code:
		return d, true
	case AllOptions:
		return d.Code, true
	default:
		return kind.Unknown, false
	}
}

func (err *Error[_]) Severity() (severity.Level, bool) {
	switch d := any(err.Data).(type) {
	case severity.Level:
		return d, true
	case AllOptions:
		return d.Level, true
	default:
		return severity.Unknown, false
	}
}

func (err *Error[_]) CausedBy() (error, bool) {
	switch d := any(err.Data).(type) {
	case CauseOption:
		return d.Error, true
	case AllOptions:
		return d.Cause, true
	default:
		return nil, false
	}
}
