package oops

import "fmt"

type ErrorOption func(*Error)

func CausedBy(parent error) ErrorOption {
	return func(e *Error) {
		// TODO: implement a general Custom formatter for the client
		e.wrapper = fmt.Errorf("%v caused by (%w)", e.wrapper, parent)
	}
}
