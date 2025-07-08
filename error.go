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

	default:
		return &internal.Error[D]{Msg: msg, Data: data}
	}
}

type metaError[T any] struct {
	source error
	meta   T
}

func (err *metaError[_]) Error() string { return err.source.Error() }
func (err *metaError[T]) Unwrap() error { return err.source }
func (err *metaError[T]) Kind() (kind.Code, bool) {
	switch meta := any(err.meta).(type) {
	case kind.Code:
		return meta, true

	case kind.Reader:
		return meta.Kind()

	default:
		if code := kind.Of(err.source); code != kind.Unknown {
			return code, true
		}
		return kind.Unknown, false
	}
}

func (err *metaError[T]) Severity() (severity.Level, bool) {
	switch meta := any(err.meta).(type) {
	case severity.Level:
		return meta, true

	case severity.Reader:
		return meta.Severity()

	default:
		if level := severity.Of(err.source); level != severity.Unknown {
			return level, true
		}
		return severity.Unknown, false
	}
}
