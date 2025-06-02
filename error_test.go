package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"testing"
)

var (
	Internal      = oops.Custom{Code: "Internal", Error: errors.New("internal error")}
	Unauthorized  = oops.Custom{Code: "Unauthorized", Error: errors.New("unauthorized access")}
	Unimplemented = oops.Custom{Code: "Unimplemented", Error: errors.New("not implemented")}
	Invalid       = oops.Custom{Code: "Invalid", Error: errors.New("invalid input")}
	Forbidden     = oops.Custom{Code: "Forbidden", Error: errors.New("forbidden access")}
	NotFound      = oops.Custom{Code: "NotFound", Error: errors.New("not found error")}
	Unprocessable = oops.Custom{Code: "Unprocessable", Error: errors.New("unprocessable")}
	inputs        = []input{
		{Internal, "An internal error occurred"},
		{Unimplemented, "This feature is not implemented"},
		{NotFound, "The requested resource was not found"},
		{Invalid, "The input provided is invalid"},
		{Unauthorized, "You are not authorized to perform this action"},
		{Forbidden, "Access to this resource is forbidden"},
		{Unprocessable, "The request could not be processed"},
	}
)

type input struct {
	custom oops.Custom
	msg    string
}

func TestNew(t *testing.T) {
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = inputs[i].custom.Code.String()
		testCases[i].custom = inputs[i].custom
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, tc.custom)
			if got == nil {
				t.Errorf("oops.New() never returns nil: expected %T {%s, %v}, got %v", &oops.Error{}, tc.custom.Code, tc.custom.Error, got)
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
	t.Log("oops.New(msg string, options ...option) in brief:")
	t.Log(" - oops.New never returns nil")
	t.Log(" - oops.New(msg).Error() leads to the client msg")
	t.Log(" - Printing oops.New(msg) with fmt.Sprintf leads to the client msg")
	t.Log(" - Comparing oops.New(msg, custom) with client custom category's Error using errors.Is() leads to true")
}

func TestNew_CausedBySuccessfullyWrappedInOopsErrorWrapper(t *testing.T) {
	parent := errors.New("parent error")
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = inputs[i].custom.Code.String()
		testCases[i].custom = inputs[i].custom
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, tc.custom, oops.CausedBy(parent))
			if !errors.Is(got, parent) {
				t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNew_ClientCustomErrorSuccessfullyWrappedInOopsError(t *testing.T) {
	got := oops.New("The requested resource was not found", NotFound)
	if !errors.Is(got, NotFound.Error) {
		t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
	}
}

func TestNew_IsRootCauseError(t *testing.T) {
	parent := errors.New("parent error")
	got := oops.New("An internal error occurred", Internal, oops.CausedBy(parent))
	if !errors.Is(got, parent) {
		t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
	}
}

func TestNew_OptionsOrderIsNotImportant(t *testing.T) {
	process := func() error {
		lowLevelNotFound := oops.New("The requested resource was not found",
			NotFound, oops.CausedBy(errors.New("a low-level error")),
		)
		return oops.New("Unprocessable entity",
			oops.CausedBy(lowLevelNotFound),
			Internal,
		)
	}
	got := process()
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error.Debug(3):", got)
	for _, err := range oopsErr.Debug(3) {
		t.Logf(" - %s", err)
	}
}

func TestBuiltinErrorAsOopsError(t *testing.T) {
	process := func() error {
		return oops.New("Something went wrong",
			Internal, oops.CausedBy(errors.New("a low-level error")),
		)
	}
	got := process()
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error.Debug(2) %q", got, oopsErr.Debug(2))
}

func TestError_Debug(t *testing.T) {
	// TODO: need to be improved
	//mainIssue := errors.New("main issue")
	got := oops.New("The request is unprocessable") //Unprocessable, oops.CausedBy{Parent: mainIssue},

	t.Logf("%+q", got.Debug(2))
}
