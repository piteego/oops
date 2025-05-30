package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"testing"
)

var (
	Internal       = oops.Identifier{Code: "Internal", Error: errors.New("internal error")}
	NotFound       = oops.Identifier{Code: "NotFound", Error: errors.New("not found")}
	Invalid        = oops.Identifier{Code: "Invalid", Error: errors.New("invalid input")}
	NotImplemented = oops.Identifier{Code: "NotImplemented", Error: errors.New("not implemented")}
	Unauthorized   = oops.Identifier{Code: "Unauthorized", Error: errors.New("unauthorized access")}
	Forbidden      = oops.Identifier{Code: "Forbidden", Error: errors.New("forbidden access")}
	Unprocessable  = oops.Identifier{Code: "Unprocessable", Error: errors.New("unprocessable")}
	inputs         = []input{
		{Internal, "An internal error occurred"},
		{NotFound, "The requested resource was not found"},
		{Invalid, "The input provided is invalid"},
		{NotImplemented, "This feature is not implemented"},
		{Unauthorized, "You are not authorized to perform this action"},
		{Forbidden, "Access to this resource is forbidden"},
		{Unprocessable, "The request could not be processed"},
	}
)

type input struct {
	custom oops.Identifier
	msg    string
}

func TestNew(t *testing.T) {
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = inputs[i].custom.Code
		testCases[i].custom = inputs[i].custom
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.custom, tc.msg)
			if got == nil {
				t.Fatalf("oops.New() never returns nil: expected %T {%s, %v}, got %v", &oops.Error{}, tc.custom.Code, tc.custom.Error, got)
			}
			if got.Error() != tc.msg {
				t.Errorf("oops.Error.Error() must lead to the client msg %q, got %q", tc.msg, got.Error())
			}
			if fmt.Sprintf("%v", got) != tc.msg {
				t.Errorf("Printing oops.Error with fmt.Sprintf must lead to the client msg %q, got %q", tc.msg, fmt.Sprintf("%v", got))
			}
			if !errors.Is(got, tc.custom.Error) {
				t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNewCausedBy(t *testing.T) {
	parent := errors.New("parent error")
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = fmt.Sprintf("New%sCausedByParentError", inputs[i].custom.Code)
		testCases[i].custom = inputs[i].custom
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.custom, tc.msg, oops.CausedBy(parent))
			if !errors.Is(got, parent) {
				t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNewIsClientCustomError(t *testing.T) {
	got := oops.New(NotFound, "The requested resource was not found")
	if !errors.Is(got, NotFound.Error) {
		t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
	}
}

func TestNewIsRootCauseError(t *testing.T) {
	parent := errors.New("parent error")
	got := oops.New(Internal, "An internal error occurred", oops.CausedBy(parent))
	if !errors.Is(got, parent) {
		t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
	}
}

func TestBuiltinErrorAsOopsError(t *testing.T) {
	process := func() error {
		return oops.New(Internal, "Something went wrong",
			oops.CausedBy(errors.New("a low-level error")),
		)
	}
	got := process()
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error.Debug(2) %q", got, oopsErr.Debug(2))
}
