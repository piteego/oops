package oops

import "errors"

var Unknown = category{Id: "Unknown", Error: errors.New("unknown error category")}

type category struct {
	Id    id
	Error error
}

type id string

func (c id) String() string { return string(c) }
