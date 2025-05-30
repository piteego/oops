package oops

import "errors"

var Oops = identifier{Code: "Oops!", Error: errors.New("something went wrong")}

type (
	code       string
	identifier struct {
		Code  code
		Error error
	}
)

func (c code) String() string { return string(c) }
