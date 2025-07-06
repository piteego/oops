package oops

import "errors"

func SetMetadata[D data](err error, meta D) error { return &setError[D]{base: err, data: meta} }

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
