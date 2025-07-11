package v1_0

import (
	"errors"
)

type (
	metadata interface{ errorMetadata() }
	Metadata struct{}
)

func (Metadata) errorMetadata() {}

func Unwrap[T metadata](err error) T {
	switch e := err.(type) {
	case *richError[T]:
		return e.meta

	default:
		return Unwrap[T](errors.Unwrap(err))
	}
}

func CauseOf(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case cause:
		return e.Cause()

	case interface{ Unwrap() error }:
		return CauseOf(e.Unwrap())

	default:
		return nil
	}
}
