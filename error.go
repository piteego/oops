package oops

func New[T any](msg string, options ...func(*T)) error {
	n := 0
	for i := range options {
		if options[i] != nil {
			n++
		}
	}
	if n == 0 {
		return &Error[struct{}]{msg: msg}
	}
	err := &Error[T]{msg: msg}
	for i := range options {
		if options[i] != nil {
			options[i](&err.data)
		}
	}
	return err
}

type Error[T any] struct {
	msg  string
	data T
}

func (err *Error[_]) Error() string { return err.msg }

func (err *Error[T]) Unwrap() error {
	switch t := any(err.data).(type) {
	case interface{ Unwrap() error }:
		return t.Unwrap()

	default:
		return nil
	}
}

func (err *Error[T]) Data() T { return err.data }
