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

func (err *Error[_]) Kind() kind.Code {
	switch data := any(err.Data).(type) {
	case kind.Code:
		return data
	case AllOptions:
		return data.Code
	default:
		return kind.Unknown
	}
}

func (err *Error[_]) Severity() severity.Level {
	switch data := any(err.Data).(type) {
	case severity.Level:
		return data
	case AllOptions:
		return data.Level
	default:
		return severity.Unknown
	}
}
