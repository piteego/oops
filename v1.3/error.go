package v1_3

import (
	"github.com/piteego/oops/internal"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type Option func(error) error

func New(msg string, options ...Option) error {
	var err error = &msgError{msg: msg}
	for i := range options {
		if options[i] == nil {
			continue
		}
		err = options[i](err)
	}
	return err
}

type msgError struct{ msg string }

func (err *msgError) Error() string { return err.msg }

type wrap[D any] struct {
	head error
	data D
}

func (err *wrap[_]) Error() string { return err.head.Error() }

func (err *wrap[_]) Unwrap() error { return err.head }

type kindError = wrap[kind.Code]

func WithKind(kind kind.Code) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *kindError:
			err.data = kind
			return err

		default:
			return &kindError{head: err, data: kind}
		}
	}
}

type severityError = wrap[severity.Level]

func WithSeverity(severity severity.Level) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *severityError:
			err.data = severity
			return err

		default:
			return &severityError{head: err, data: severity}
		}
	}
}

type (
	causeError = wrap[internal.CauseOption]
)

func WithCause(cause error) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *causeError:
			err.data = internal.CauseOption{Error: cause}
			return err

		default:
			return &causeError{head: err, data: internal.CauseOption{Error: cause}}
		}
	}
}

func WithMetadata[D metadata](meta D) Option {
	return func(err error) error {
		return &wrap[D]{head: err, data: meta}
	}
}
