package v1_0_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v1_0"
)

func ExampleAnalyze() {
	errs := []error{
		errors.New("builtin error using errors.New"),
		fmt.Errorf("wrap '%w' using fmt.Errorf", errors.New("another builtin error")),
		v1_0.New("rich error using oops.New",
			v1_0.Tag{Label: example.NotFound.Error},
			v1_0.Because{Error: errors.New("cause of the error")},
			v1_0.Medium.Diag("this is a diagnostic note"),
			example.Metadata{Retry: true, Code: 96},
		),
	}
	for i := range errs {
		analysis := v1_0.Analyze(errs[i])
		if analysis != nil {
			fmt.Printf("%+v\n", analysis)
		}
	}
	// Output:
	//&{Message:builtin error using errors.New Label:untagged oops error Cause:<nil> Diagnosis:{severity: Unknown, note: ""} Metadata:<nil>}
	//&{Message:wrap 'another builtin error' using fmt.Errorf Label:untagged oops error Cause:another builtin error Diagnosis:{severity: Unknown, note: ""} Metadata:<nil>}
	//&{Message:rich error using oops.New Label:resource not found Cause:cause of the error Diagnosis:{severity: Medium, note: "this is a diagnostic note"} Metadata:{Code: 96, Retry: true}}
}
