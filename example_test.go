package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

func ExampleNew_withPredefinedOptions() {
	errs := []error{
		errors.New("just use errors.New in case of no options"),
		oops.New("an error with kind of 3!", kind.Code(3)),
		oops.New("an error with severity level of 2!", severity.Level(2)),
		oops.New(`an error caused by "a cause error"`, oops.CausedBy{Error: example.ErrorCause}),
		oops.New(`an error with {kind:3, level: 2, cause: "a cause error"}`,
			oops.Standard{Code: 3, Level: 2, Cause: example.ErrorCause},
		),
	}
	for i := range errs {
		fmt.Println(errs[i])
	}
	// Output:
	// just use errors.New in case of no options
	// an error with kind of 3!
	// an error with severity level of 2!
	// an error caused by "a cause error"
	// an error with {kind:3, level: 2, cause: "a cause error"}
}

func ExampleNew_withCustomMetadata() {
	err := oops.New("an error", oops.With(example.Metadata{Id: "E109", Retry: true}))
	metadata := oops.Unwrap[example.Metadata](err)
	fmt.Printf("%q with custom metadata:{Id: %q, Retry: %v}", err, metadata.Id, metadata.Retry)
	// Output:
	// "an error" with custom metadata:{Id: "E109", Retry: true}
}

func ExampleWith() {
	err := oops.With(
		example.Metadata{Id: "E109", Retry: true},
	)(
		oops.New("an error", kind.Code(2)),
	)
	metadata := oops.Unwrap[example.Metadata](err)
	fmt.Printf("%q with code: %d, and custom metadata:{Id: %q, Retry: %v}",
		err,
		oops.KindOf(err),
		metadata.Id, metadata.Retry,
	)
	// Output:
	// "an error" with code: 2, and custom metadata:{Id: "E109", Retry: true}
}
