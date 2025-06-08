package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
)

func ExampleAnalyze() {
	errs := []error{
		errors.New("builtin error using errors.New"),
		fmt.Errorf("wrap '%w' using fmt.Errorf", errors.New("another builtin error")),
		oops.New("custom error using oops.New", oops.Because{Error: errors.New("cause of the error")}),
	}
	for i := range errs {
		analysis := oops.Analyze(errs[i])
		if analysis != nil {
			fmt.Printf("%+v\n", analysis)
		}
	}

	// Output:
	//&{Message:builtin error using errors.New Label:untagged oops error Cause:<nil> Diagnosis:{Note: Severity:Undefined} Metadata:<nil>}
	//&{Message:wrap 'another builtin error' using fmt.Errorf Label:untagged oops error Cause:another builtin error Diagnosis:{Note: Severity:Undefined} Metadata:<nil>}
	//&{Message:custom error using oops.New Label:untagged oops error Cause:cause of the error Diagnosis:{Note: Severity:Undefined} Metadata:<nil>}
}
