package oops

import (
	"errors"
)

func Unwrap[T metadata](err error) T {
	switch e := err.(type) {
	case *metaError[T]:
		return e.meta

	default:
		return Unwrap[T](errors.Unwrap(err))
	}
}
