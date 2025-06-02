package oops

import (
	"fmt"
)

type ErrorOption interface{ apply(*Error) }

type causedBy struct{ parent error }

func (opt causedBy) apply(err *Error) {
	if opt.parent == nil {
		return
	}
	if err.wrapper != nil {
		err.wrapper = fmt.Errorf("%s caused by (%w)", err.wrapper, opt.parent)
	} else {
		err.wrapper = fmt.Errorf("%s caused by (%w)", err.msg, opt.parent)
	}
}

func CausedBy(parent error) ErrorOption { return causedBy{parent} }

type Custom = category

func (opt Custom) apply(err *Error) { err.Category = opt }
