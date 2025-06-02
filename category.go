package oops

import "errors"

var Unknown = category{Code: "Unknown", Error: errors.New("unknown category")}

type (
	code     string
	category struct {
		Code  code
		Error error
	}
)

func (c code) String() string { return string(c) }
