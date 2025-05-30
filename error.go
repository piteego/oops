package oops

import (
	"errors"
	"slices"
)

type Identifier struct {
	Code  string
	Error error
}

func New(custom Identifier, msg string, options ...ErrorOption) *Error {
	e := &Error{
		Custom: custom,
		msg:    msg,
	}
	for i := range options {
		options[i](e)
	}
	return e
}

type Error struct {
	Custom  Identifier
	msg     string
	wrapper error
}

// Error implements golang's builtin error interface. It returns the client's message of the Error.
func (x *Error) Error() string { return x.msg }

// Unwrap returns the wrapped error, to allow interoperability with
// errors.Is(), errors.As() and errors.Unwrap()
func (x *Error) Unwrap() error {
	switch {
	case x == nil:
		return nil

	case errors.Unwrap(x.wrapper) != x.Custom.Error:
		// CausedBy called with non-nil parent error...
		return errors.Join(x.Custom.Error, x.wrapper)

	default:
		// New called without calling CausedBy, or called CausedBy(nil)
		return x.wrapper
	}
}

func (x *Error) Debug(depth int) []error {
	if x == nil {
		return nil
	}
	if depth <= 0 {
		panic("[Oops!] debug depth must be positive")
	}
	debug := make([]error, 1, depth)
	debug[0] = x.Custom.Error
	next := errors.Unwrap(x.wrapper)
	if next == x.Custom.Error || depth == 1 {
		return debug
	}
	for next != nil && len(debug) <= depth {
		debug = append(debug, next)
		next = errors.Unwrap(next)
	}
	return slices.Clip(debug)
}
