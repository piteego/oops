package v1_2_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v1_2"
)

func ExampleNew() {
	err := v1_2.New("this is a message",
		v1_2.Chain(
			v1_2.WithIdentity(example.Internal.Error, 2),
			v1_2.SeverityLow.Diag(example.Unimplemented.Error, "this is a note"),
		),
	)
	if err != nil {
		fmt.Println(err)
		var oopsErr *v1_2.StdError
		if errors.As(err, &oopsErr) {
			std := oopsErr.Data()
			fmt.Printf("diag{cause: %q, severity: %q, note: %q}\n", std.Diag.Cause, std.Diag.Severity, std.Diag.Note)
			fmt.Printf("identity{code: %d, kind: %s}\n", std.Identity.Code, std.Identity.Kind)
		}
		if errors.Is(err, example.Internal.Error) {
			fmt.Printf("oops error is of kind %q", example.Internal.Error)
		}
	}
	// Output:
	// this is a message
	// diag{cause: "not implemented yet", severity: "Low", note: "this is a note"}
	// identity{code: 2, kind: something went wrong}
	// oops error is of kind "something went wrong"
}
