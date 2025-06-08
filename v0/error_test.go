package v0_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v0"
	"testing"
)

var (
	commonTestCases = []struct {
		name string
		// inputs
		label v0.Label
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
			got := v0.New(tc.msg, v0.Tag(tc.label))
			if got == nil {
				t.Errorf("oops.New() never returns nil, got nil")
			}
			_ = got.(*v0.Error) // ensure that got is of type *oops.Error
			if got.Error() != tc.msg {
				t.Errorf("oops.Error.Error() must lead to the client msg %q, got %q", tc.msg, got.Error())
			}
			if fmt.Sprintf("%v", got) != tc.msg {
				t.Errorf("Printing oops.Error with fmt.Sprintf must lead to the client msg %q, got %q", tc.msg, fmt.Sprintf("%v", got))
			}
			if !errors.Is(got, tc.label) {
				t.Logf("##### got.(*oops.Error).Unwrap(): %v", got.(*v0.Error).Unwrap())
				t.Errorf("Comparing oops.Error with client custom Label using errors.Is() must lead to true, got false")
			}
		})
	}
	t.Log("oops.New(msg string, options ...option) in brief:")
	t.Log(" - oops.New never returns nil")
	t.Log(" - New() error is always of type *oops.Error")
	t.Log(" - oops.New(msg).Error() leads to the client msg")
	t.Log(" - Printing oops.New(msg) with fmt.Sprintf leads to the client msg")
	t.Log(" - Comparing oops.New(msg, oops.Tag(custom)) with client custom Label using errors.Is() leads to true")
}

func TestNew_CausedBySuccessfullyWrappedInOopsErrorWrapper(t *testing.T) {
	cause := errors.New("cause error")
	for _, tc := range commonTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := v0.New(tc.msg, v0.Tag(tc.label), v0.Because(cause))
			if !errors.Is(got, cause) {
				t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
			}
		})
	}
}

func TestNew_ClientCustomErrorSuccessfullyWrappedInOopsError(t *testing.T) {
	got := v0.New("The requested resource was not found", v0.Tag(example.NotFound.Error))
	if !errors.Is(got, example.NotFound.Error) {
		t.Errorf("Comparing oops.Error with client custom error using errors.Is() must lead to true, got false")
	}
}

func TestNew_IsRootCauseError(t *testing.T) {
	cause := errors.New("cause error")
	got := v0.New("An internal error occurred", v0.Tag(example.Internal.Error), v0.Because(cause))
	if !errors.Is(got, cause) {
		t.Errorf("Comparing oops.Error with its root cause error using errors.Is() must lead to true, got false")
	}
}

func TestNew_OptionsOrderIsNotImportant(t *testing.T) {
	process := func() error {
		lowLevelNotFound := v0.New("The requested resource was not found",
			v0.Tag(example.NotFound.Error), v0.Because(errors.New("a low-level error")),
		)
		return v0.New("Unprocessable entity",
			v0.Because(lowLevelNotFound),
			v0.Tag(example.Internal.Error),
		)
	}
	got := process()
	var oopsErr *v0.Error
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
		return v0.New("Something went wrong",
			v0.Tag(example.Internal.Error), v0.Because(errors.New("a low-level error")),
		)
	}
	got := process()
	var oopsErr *v0.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("Successfully parsed process() error %q as oops.Error %+v", got, oopsErr)
}

func TestError_Unwrap(t *testing.T) {
	// TODO: need to be improved
	mainIssue := errors.New("main issue")
	got := v0.New("The request is unprocessable",
		v0.Tag(example.Unprocessable.Error), v0.Because(mainIssue),
	)
	var oopsErr *v0.Error
	if !errors.As(got, &oopsErr) {
		t.Errorf("expected *oops.Error, got %T", got)
	}
	t.Logf("%+q", oopsErr.Unwrap())
}
