package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
)

func ExampleNew() {
	err := oops.New("this is a message", oops.With(
		example.Internal.Error,
		oops.Low.Diag("this is a note").CausedBy(example.Unimplemented.Error),
	))
	if err != nil {
		fmt.Println(err)
		var oopsErr *oops.Error[oops.Standard]
		if errors.As(err, &oopsErr) {
			data := oopsErr.Data()
			fmt.Printf("kind: %q, cause: %q, severity: %q, note: %q\n", data.Kind, data.Diag.Cause(), data.Diag.Severity(), data.Diag.Note())
		}
		if errors.Is(err, example.Internal.Error) {
			fmt.Printf("oops error is of kind %q", example.Internal.Error)
		}
	}
	// Output:
	// this is a message
	// kind: "something went wrong", cause: "not implemented yet", severity: "Low", note: "this is a note"
	// oops error is of kind "something went wrong"
}
