package oops

import "fmt"

type Code interface {
	Index() int
	Err() error
}

type Reporter interface {
	error
	// Code returns Custom identifier of an Error.
	Code() Code
	// Trace returns all the wrapped errors of an Error.
	Trace() []error
}

type Error interface {
	Reporter
	//CausedBy wraps the given error in an Error.
	CausedBy(parent error) Reporter
}

func New(code Code, msg string) Error {
	if code == nil {
		return nil
	}
	return &custom{
		msg: msg, code: code,
		error: fmt.Errorf("[%d][%w]::%q", code.Index(), code.Err(), msg),
	}
}
