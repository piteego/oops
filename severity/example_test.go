package severity_test

import (
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/severity"
)

func ExampleIs() {
	fmt.Println(severity.Is(2, severity.Unknown))
	// Output:
	// false
}

func ExampleOf() {
	errs := []error{
		oops.New("an error", oops.With(example.Metadata{Retry: true})),
		oops.New("an error", severity.Level(1)),
		oops.New("an error", oops.Standard{Level: 2}),
		oops.New("an error", oops.With(example.MetaStandard{Level: 3})),
	}
	for i := range errs {
		fmt.Printf("%q with severity level of %d\n", errs[i], severity.Of(errs[i]))
	}
	// Output:
	// "an error" with severity level of -1
	// "an error" with severity level of 1
	// "an error" with severity level of 2
	// "an error" with severity level of 3
}
