package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"testing"
)

var (
	commonTestCases = []struct {
		name string
		// inputs
		label oops.Label
		msg   string
	}{
		{"ExampleLabelInternal", example.Internal.Error, "An internal error occurred"},
		{"ExampleLabelFound", example.NotFound.Error, "The requested resource was not found"},
		{"ExampleLabelForbidden", example.Forbidden.Error, "Access to this resource is forbidden"},
		{"ExampleLabelValidation", example.Validation.Error, "The input provided is invalid"},
		{"ExampleLabelUnauthorized", example.Unauthorized.Error, "You are not authorized to perform this action"},
		{"ExampleLabelUnimplemented", example.Unimplemented.Error, "This feature is not implemented"},
		{"ExampleLabelUnprocessable", example.Unprocessable.Error, "The request could not be processed"},
	}
)

func TestNew(t *testing.T) {
	for _, tc := range commonTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, oops.Tag(tc.label))
			if got == nil {
				t.Errorf("oops.New() never returns nil, got nil")
			}
			if got.Error() != tc.msg {
				t.Errorf("oops.Error.Error() must lead to the client msg %q, got %q", tc.msg, got.Error())
			}
			if fmt.Sprintf("%v", got) != tc.msg {
				t.Errorf("Printing oops.Error with fmt.Sprintf must lead to the client msg %q, got %q", tc.msg, fmt.Sprintf("%v", got))
			}
			if !errors.Is(got, tc.label) {
				t.Logf("##### got.(*oops.Error).Unwrap(): %v", got.(*oops.Error).Unwrap())
				t.Errorf("Comparing oops.Error with client custom Label using errors.Is() must lead to true, got false")
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
	for _, tc := range commonTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := oops.New(tc.msg, oops.Tag(tc.label), oops.Because(cause))
			if !errors.Is(got, cause) {
				t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNew_ClientCustomErrorSuccessfullyWrappedInOopsError(t *testing.T) {
	got := oops.New("The requested resource was not found", oops.Tag(example.NotFound.Error))
	if !errors.Is(got, example.NotFound.Error) {
		t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
	}
}

func TestNew_IsRootCauseError(t *testing.T) {
	cause := errors.New("cause error")
	got := oops.New("An internal error occurred", oops.Tag(example.Internal.Error), oops.Because(cause))
	if !errors.Is(got, cause) {
		t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
	}
}

func TestNew_OptionsOrderIsNotImportant(t *testing.T) {
	process := func() error {
		lowLevelNotFound := oops.New("The requested resource was not found",
			oops.Tag(example.NotFound.Error), oops.Because(errors.New("a low-level error")),
		)
		return oops.New("Unprocessable entity",
			oops.Because(lowLevelNotFound),
			oops.Tag(example.Internal.Error),
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
			oops.Tag(example.Internal.Error), oops.Because(errors.New("a low-level error")),
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
		oops.Tag(example.Unprocessable.Error), oops.Because(mainIssue),
	)
	var oopsErr *oops.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("%+q", oopsErr.Unwrap())
}
