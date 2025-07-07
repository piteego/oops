package oops

type Option func(error) error

func New(msg string, options ...Option) error {
	var err error = &msgError{msg: msg}
	for i := range options {
		err = options[i](err)
	}
	return err
}

type msgError struct{ msg string }

func (err *msgError) Error() string { return err.msg }

type (
	data             = any
	setError[D data] struct {
		base error
		data D
	}
)

func (err *setError[_]) Error() string { return err.base.Error() }

func (err *setError[_]) Unwrap() error { return err.base }

type kindError = setError[Code]

func WithKind(kind Code) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *kindError:
			err.data = kind
			return err

		default:
			return &kindError{base: err, data: kind}
		}
	}
}

type causeError = setError[Cause]

func WithCause(cause Cause) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *causeError:
			err.data = cause
			return err

		default:
			return &causeError{base: err, data: cause}
		}
	}
}

type severityError = setError[Level]

func WithSeverity(severity Level) Option {
	return func(e error) error {
		switch err := e.(type) {
		case *severityError:
			err.data = severity
			return err

		default:
			return &severityError{base: err, data: severity}
		}
	}
}

func WithMetadata[D data](meta D) Option {
	return func(err error) error {
		return &setError[D]{base: err, data: meta}
	}
}
