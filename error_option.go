package oops

import (
	"fmt"
)

type ErrorOption interface{ apply(*Error) }

type causedBy struct{ cause error }

func (opt causedBy) apply(err *Error) {
	if opt.cause == nil {
		return
	}
	if err.wrapper != nil {
		err.wrapper = fmt.Errorf("%w ==> %s", opt.cause, err.wrapper)
	} else {
		err.wrapper = fmt.Errorf("%w ==> ", opt.cause)
	}
}

func CausedBy(cause error) ErrorOption { return causedBy{cause} }

type Category = category

func (opt Category) apply(err *Error) {
	if err.Category.Error != nil {
		// Do not overwrite existing category
		return
	}
	err.Category = opt
}
