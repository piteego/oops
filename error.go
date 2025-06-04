package oops

import (
	"errors"
)

var Untagged Label = errors.New("unknown error label")

// Label is a type alias for error, used to categorize application errors.
type Label error

// New creates a new *[Error] and return builtin error interface
// with the given message and list of [ErrorOption].
// You can use optional [ErrorOption] (e.g, [CausedBy] to benefit stack trace,
// or [Tag] a [Label] to categorize your application errors)
func New(msg string, options ...ErrorOption) error {
	err := Error{msg: msg}
	for i := range options {
		options[i](&err)
	}
	// If no label is set, use the default untagged label.
	if err.Label == nil {
		err.Label = Untagged
	}
	// append the label to the stack trace
	CausedBy(err.Label)(&err)
	return &err
}

// Error is a labeled error with stack trace implements the builtin error interface.
type Error struct {
	Label
	msg   string
	stack []error
}

// Error implements golang's builtin error interface. It returns the client's message given in the [New] function.
func (err *Error) Error() string { return err.msg }

// Unwrap returns the wrapped errors, to allow interoperability with
// [errors.Is](), [errors.As]() and [errors.Unwrap]()
func (err *Error) Unwrap() []error { return err.stack }
