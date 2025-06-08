package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
)

func ExampleAnalyze() {
	errs := []error{
		errors.New("builtin error using errors.New"),
		fmt.Errorf("wrap '%w' using fmt.Errorf", errors.New("another builtin error")),
		oops.New("rich error using oops.New",
			oops.Tag{Label: example.NotFound.Error},
			oops.Because{Error: errors.New("cause of the error")},
			oops.Medium.Diag("this is a diagnostic note"),
			example.Metadata{Retry: true, Code: 96},
		),
	}
	for i := range errs {
		analysis := oops.Analyze(errs[i])
		if analysis != nil {
			fmt.Printf("%+v\n", analysis)
		}
	}
	// Output:
	//&{Message:builtin error using errors.New Label:untagged oops error Cause:<nil> Diagnosis:{severity: Unknown, note: ""} Metadata:<nil>}
	//&{Message:wrap 'another builtin error' using fmt.Errorf Label:untagged oops error Cause:another builtin error Diagnosis:{severity: Unknown, note: ""} Metadata:<nil>}
	//&{Message:rich error using oops.New Label:resource not found Cause:cause of the error Diagnosis:{severity: Medium, note: "this is a diagnostic note"} Metadata:{Code: 96, Retry: true}}
}
