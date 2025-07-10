package oops

import (
	"errors"
	"github.com/piteego/oops/internal"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

func New[D interface{ internal.Option | metaOption }](msg string, data D) error {
	switch opt := any(data).(type) {
	case metaOption:
		if opt == nil { // Unreachable until expose metaOption type!
			return errors.New(msg)
		}
		return opt(errors.New(msg))

	default: // unreachable!
		return &internal.Error[D]{Msg: msg, Data: data}
	}
}

type metaError[T any] struct {
	source error
	meta   T
}

func (err *metaError[_]) Error() string { return err.source.Error() }
func (err *metaError[T]) Unwrap() error { return err.source }
func (err *metaError[T]) Kind() kind.Code {
	switch meta := any(err.meta).(type) {
	case kind.Code:
		return meta

	case kind.Reader:
		return meta.Kind()

	default:
		if code := kind.Of(err.source); code != kind.Unknown {
			return code
		}
		return kind.Unknown
	}
}
func (err *metaError[T]) Severity() severity.Level {
	switch meta := any(err.meta).(type) {
	case severity.Level:
		return meta

	case severity.Reader:
		return meta.Severity()

	default:
		return severity.Of(err.source)
	}
}
