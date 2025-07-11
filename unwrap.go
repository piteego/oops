package oops

import "errors"

func As[M metadata](err error, metadata *M) bool {
	if err == nil {
		return false
	}
	if metadata == nil {
		panic("trying to unwrap using nil metadata")
	}
	switch e := err.(type) {
	case *richError[M]:
		*metadata = e.meta
		return true

	default:
		return As[M](errors.Unwrap(err), metadata)
	}
}
