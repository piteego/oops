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

func ExampleAnalysis() {
	errs := []error{
		errors.New("builtin error using errors.New"),
		fmt.Errorf("wrap '%w' using fmt.Errorf", errors.New("another builtin error")),
		oops.New("rich error using oops.New",
			oops.Tag{Label: example.NotFound.Error},
			oops.Because{Error: errors.New("cause of the error")},
			oops.Diagnosis{Note: "this is a diagnostic note", Severity: oops.Medium},
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
	//&{Message:builtin error using errors.New Label:untagged oops error Cause:<nil> Diagnosis:{Note: Severity:Undefined} Metadata:<nil>}
	//&{Message:wrap 'another builtin error' using fmt.Errorf Label:untagged oops error Cause:another builtin error Diagnosis:{Note: Severity:Undefined} Metadata:<nil>}
	//&{Message:rich error using oops.New Label:resource not found Cause:cause of the error Diagnosis:{Note:this is a diagnostic note Severity:Medium} Metadata:{Code: 96, Retry: true}}
}
