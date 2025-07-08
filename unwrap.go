package oops

import (
	"errors"
	"github.com/piteego/oops/internal"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

func KindOf(err error) kind.Code {
	switch e := err.(type) {
	case *Error[kind.Code]:
		return e.meta

	case *internal.Error[kind.Code]:
		return e.Data

	case *internal.Error[Standard]:
		return e.Data.Code

	case interface{ Unwrap() error }:
		return KindOf(e.Unwrap())

	default:
		return kind.Unknown
	}
}

func SeverityOf(err error) severity.Level {
	switch e := err.(type) {
	case *Error[severity.Level]:
		return e.meta

	case *internal.Error[severity.Level]:
		return e.Data

	case *internal.Error[Standard]:
		return e.Data.Level

	case interface{ Unwrap() error }:
		return SeverityOf(e.Unwrap())

	default:
		return severity.Unknown
	}
}

func CauseOf(err error) error {
	switch e := err.(type) {
	case *Error[CausedBy]:
		return e.meta.Error

	case *internal.Error[CausedBy]:
		return e.Data.Error

	case *internal.Error[Standard]:
		return e.Data.Cause

	case interface{ Unwrap() error }:
		return CauseOf(e.Unwrap())

	default:
		return nil
	}
}

func Unwrap[T metadata](err error) T {
	switch e := err.(type) {
	case *Error[T]:
		return e.meta

	default:
		return Unwrap[T](errors.Unwrap(err))
	}
}
