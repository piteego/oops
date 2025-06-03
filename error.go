package oops

import (
	"errors"
)

var Untagged = Label{Id: "Oops!", Error: errors.New("unknown error label")}

type Label struct {
	Id    string
	Error error
}

func New(msg string, options ...ErrorOption) error {
	err := Error{msg: msg}
	for i := range options {
		options[i](&err)
	}
	if err.Label == (Label{}) {
		err.Label = Untagged
	}
	CausedBy(err.Label.Error)(&err)
	return &err
}

type Error struct {
	Label
	msg   string
	stack []error
}

// Error implements golang's builtin error interface. It returns the client's message given in the New function.
func (err *Error) Error() string { return err.msg }

// Unwrap returns the wrapped errors, to allow interoperability with
// errors.Is(), errors.As() and errors.Unwrap()
func (err *Error) Unwrap() []error { return err.stack }
