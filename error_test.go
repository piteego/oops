package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"testing"
)

var (
	Internal      = oops.Label{Id: "Internal", Error: errors.New("something went wrong")}
	NotFound      = oops.Label{Id: "NotFound", Error: errors.New("resource not found")}
	Forbidden     = oops.Label{Id: "Forbidden", Error: errors.New("forbidden access")}
	Validation    = oops.Label{Id: "Validation", Error: errors.New("invalid input")}
	Duplication   = oops.Label{Id: "Duplication", Error: errors.New("duplicate entry")}
	Unauthorized  = oops.Label{Id: "Unauthorized", Error: errors.New("unauthorized access")}
	Unimplemented = oops.Label{Id: "Unimplemented", Error: errors.New("not implemented yet")}
	Unprocessable = oops.Label{Id: "Unprocessable", Error: errors.New("the request is unprocessable")}
	inputs        = []input{
		{Internal, "An internal error occurred"},
		{NotFound, "The requested resource was not found"},
		{Forbidden, "Access to this resource is forbidden"},
		{Validation, "The input provided is invalid"},
		{Unauthorized, "You are not authorized to perform this action"},
		{Unimplemented, "This feature is not implemented"},
		{Unprocessable, "The request could not be processed"},
	}
)

type input struct {
	label oops.Label
	msg   string
}

func TestNew(t *testing.T) {
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = inputs[i].label.Id
		testCases[i].label = inputs[i].label
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, oops.Tag(tc.label))
			if got == nil {
				t.Errorf("oops.New() never returns nil: expected %T {%s, %v}, got %v", &oops.Error{}, tc.label.Id, tc.label.Error, got)
			}
			if got.Error() != tc.msg {
				t.Errorf("oops.Error.Error() must lead to the client msg %q, got %q", tc.msg, got.Error())
			}
			if fmt.Sprintf("%v", got) != tc.msg {
				t.Errorf("Printing oops.Error with fmt.Sprintf must lead to the client msg %q, got %q", tc.msg, fmt.Sprintf("%v", got))
			}
			if !errors.Is(got, tc.label.Error) {
				t.Logf("##### got.(*oops.Error).Unwrap(): %v", got.(*oops.Error).Unwrap())
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
	cause := errors.New("cause error")
	testCases := make([]struct {
		name string
		input
	}, len(inputs))
	for i := range inputs {
		testCases[i].name = inputs[i].label.Id
		testCases[i].label = inputs[i].label
		testCases[i].msg = inputs[i].msg
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, oops.Tag(tc.label), oops.CausedBy(cause))
			if !errors.Is(got, cause) {
				t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNew_ClientCustomErrorSuccessfullyWrappedInOopsError(t *testing.T) {
	got := oops.New("The requested resource was not found", oops.Tag(NotFound))
	if !errors.Is(got, NotFound.Error) {
		t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
	}
}

func TestNew_IsRootCauseError(t *testing.T) {
	cause := errors.New("cause error")
	got := oops.New("An internal error occurred", oops.Tag(Internal), oops.CausedBy(cause))
	if !errors.Is(got, cause) {
		t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
	}
}

func TestNew_OptionsOrderIsNotImportant(t *testing.T) {
	process := func() error {
		lowLevelNotFound := oops.New("The requested resource was not found",
			oops.Tag(NotFound), oops.CausedBy(errors.New("a low-level error")),
		)
		return oops.New("Unprocessable entity",
			oops.CausedBy(lowLevelNotFound),
			oops.Tag(Internal),
		)
	}
	got := process()
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error.Unwrap():", got)
	for _, err := range oopsErr.Unwrap() {
		t.Logf(" - %s", err)
	}
}

func TestBuiltinErrorsAsOopsError(t *testing.T) {
	process := func() error {
		return oops.New("Something went wrong",
			oops.Tag(Internal), oops.CausedBy(errors.New("a low-level error")),
		)
	}
	got := process()
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error %+v", got, oopsErr)
}

func TestError_Unwrap(t *testing.T) {
	// TODO: need to be improved
	mainIssue := errors.New("main issue")
	got := oops.New("The request is unprocessable",
		oops.Tag(Unprocessable), oops.CausedBy(mainIssue),
	)
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("%+q", oopsErr.Unwrap())
}
