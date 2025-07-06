package v1_2

type data interface{ errData() }

func New[D data](msg string, options ...func(*D)) error {
	n := 0
	for i := range options {
		if options[i] != nil {
			n++
		}
	}
	if n == 0 {
		return &msgError{msg: msg}
	}
	err := &Error[D]{msg: msg}
	for i := range options {
		if options[i] != nil {
			options[i](&err.custom)
		}
	}
	return err
}

type null struct{}

func (null) errData() {}

type (
	msgError  = Error[null]
	DiagError = Error[diagnosis]
	StdError  = Error[standard]
)

type Error[D data] struct {
	msg    string
	custom D
}

func (err *Error[_]) Error() string { return err.msg }

func (err *Error[D]) Unwrap() error {
	switch d := any(err.custom).(type) {
	case interface{ Unwrap() error }:
		return d.Unwrap()

	default:
		return nil
	}
}

func (err *Error[T]) Data() T { return err.custom }

type standard struct {
	Identity identity
	Diag     diagnosis
}

func (standard) errData() {}

func (s standard) Unwrap() error { return s.Identity.Kind }

func Chain(id func(*identity), diag func(*diagnosis)) func(*standard) {
	return func(s *standard) {
		if id != nil {
			id(&s.Identity)
		}
		if diag != nil {
			diag(&s.Diag)
		}
	}
}
