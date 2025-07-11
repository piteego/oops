package v1_0

import (
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

func CausedBy(err error) Cause {
	if err == nil {
		panic("trying to set nil cause")
	}
	return Cause{err: err}
}

func Append[T metadata](metadata T) modifier {
	return func(e error) error {
		if e == nil {
			panic("trying to set metadata to nil error")
		}
		switch err := e.(type) {
		case *richError[T]:
			err.meta = metadata
			return err

		default:
			return &richError[T]{source: err, meta: metadata}
		}
	}
}

type (
	cause    interface{ Cause() error }
	Cause    struct{ err error }
	Standard struct {
		kind.Code
		severity.Level
		CausedBy error
	}
	modifier func(err error) error
)

func (c Cause) Cause() error    { return c.err }
func (c Standard) Cause() error { return c.CausedBy }
