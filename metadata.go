package oops

import "errors"

func GetMetadata[D data](err error) D {
	if err == nil {
		var zero D
		return zero
	}
	switch e := err.(type) {
	case *setError[D]:
		return e.data

	default:
		return GetMetadata[D](errors.Unwrap(err))
	}
}
