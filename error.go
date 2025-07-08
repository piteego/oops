package oops

import (
	"errors"
	"github.com/piteego/oops/internal"
)

func New[O interface{ internal.Option | metaOption }](msg string, option O) error {
	switch opt := any(option).(type) {
	case metaOption:
		if opt == nil {
			return errors.New(msg)
		}
		return opt(errors.New(msg))

	default:
		return &internal.Error[O]{Msg: msg, Data: option}
	}
}

type Error[T any] struct {
	source error
	meta   T
}

func (err *Error[_]) Error() string { return err.source.Error() }
func (err *Error[T]) Unwrap() error { return err.source }
