package oops

import (
	"errors"
	"slices"
)

func New(msg string, options ...ErrorOption) *Error {
	err := &Error{msg: msg}
	for i := range options {
		options[i].apply(err)
	}
	if err.Category.Error == nil {
		err.Category = Unknown
	}
	return err
}

type Error struct {
	Category category
	msg      string
	wrapper  error // TODO: fmt.Errorf("[%d][%w]::%q", code.Index(), code.Err(), msg)
}

// Error implements golang's builtin error interface. It returns the client's message of the Error.
func (x *Error) Error() string { return x.msg }

// Unwrap returns the wrapped error, to allow interoperability with
// errors.Is(), errors.As() and errors.Unwrap()
func (x *Error) Unwrap() error {
	switch {
	case x == nil:
		return nil

	case errors.Unwrap(x.wrapper) != x.Category.Error:
		// CausedBy applied with non-nil parent error...
		return errors.Join(x.Category.Error, x.wrapper)

	default:
		// New called without CausedBy, or called CausedBy(nil)
		return x.wrapper
	}
}

func (x *Error) Debug(depth int) []error {
	if x == nil {
		return nil
	}
	if depth <= 0 {
		panic("Oops! debug depth must be positive")
	}
	debug := make([]error, 0, depth)
	if x.Category.Error != nil {
		debug = append(debug, x.Category.Error)
	}
	next := errors.Unwrap(x.wrapper)
	if next == x.Category.Error || depth == 1 {
		return debug
	}
	for next != nil && len(debug) <= depth {
		debug = append(debug, next)
		next = errors.Unwrap(next)
	}
	return slices.Clip(debug)
}
