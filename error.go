package oops

import (
	"errors"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

func New[T interface {
	kind.Code | severity.Level | Cause | Standard | modifier
}](msg string, data T) error {
	switch opt := any(data).(type) {
	case modifier:
		return opt(errors.New(msg))

	default:
		return &richError[T]{source: errors.New(msg), meta: data}
	}
}

type richError[T any] struct {
	source error
	meta   T
}

func (err *richError[_]) Error() string { return err.source.Error() }
func (err *richError[_]) Unwrap() error { return err.source }
func (err *richError[_]) Kind(sanitize bool) kind.Code {
	switch meta := any(err.meta).(type) {
	case kind.Reader:
		return meta.Kind(sanitize)

	default:
		if sanitize {
			return kind.Of(err.source).Sanitize()
		}
		return kind.Of(err.source)
	}
}
func (err *richError[T]) Severity(sanitize bool) severity.Level {
	switch meta := any(err.meta).(type) {
	case severity.Reader:
		return meta.Severity(sanitize)

	default:
		return severity.Of(err.source).Severity(sanitize)
	}
}
func (err *richError[T]) Cause() error {
	switch meta := any(err.meta).(type) {
	case Cause:
		return meta.err

	case cause:
		return meta.Cause()

	default:
		return CauseOf(err.source)
	}
}
