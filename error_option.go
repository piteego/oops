package oops

import (
	"fmt"
)

type option interface{ apply(*Error) }

type CausedBy struct{ Parent error }

func (opt CausedBy) apply(err *Error) {
	if opt.Parent == nil {
		return
	}
	if err.wrapper != nil {
		err.wrapper = fmt.Errorf("%s caused by (%w)", err.wrapper, opt.Parent)
	} else {
		err.wrapper = fmt.Errorf("%s caused by (%w)", err.msg, opt.Parent)
	}
}

type Custom = identifier

func (opt Custom) apply(err *Error) { err.Custom = opt }
