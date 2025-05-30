package oops

import (
	"errors"
	"fmt"
)

// Custom implements Error interface and represents the Custom error of the application.
type custom struct {
	msg  string // wrapper message
	code Code   // your internal Code
	error
}

func (x *custom) Code() Code { return x.code }

// Error implements golang's builtin error interface and returns the wrapper message of the Custom error.
func (x *custom) Error() string { return x.msg }

// Unwrap returns the wrapped error, to allow interoperability with
// errors.Is(), errors.As() and errors.Unwrap()
func (x *custom) Unwrap() error {
	switch {
	case x == nil:
		return nil
	case errors.Unwrap(x.error) != x.code.Err():
		// x.CausedBy() called with non-nil parent error...
		return errors.Join(x.code.Err(), x.error)
	default:
		// New() called without calling CausedBy(), or called CausedBy(nil)
		return x.error
	}
}

func (x *custom) Trace() []error {
	if x == nil {
		return nil
	}
	result := []error{x.code.Err()}
	next := errors.Unwrap(x.error)
	if next == x.code.Err() {
		return result
	}
	for next != nil {
		result = append(result, next)
		next = errors.Unwrap(next)
	}
	return result
}

func (x *custom) CausedBy(parent error) Reporter {
	if parent == nil {
		return x
	}
	x.error = fmt.Errorf("%v caused by (%w)", x.error, parent)
	return x
}
