package kind_test

import (
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/kind"
)

func ExampleIs() {
	fmt.Println(kind.Is(2, kind.Unknown))
	// Output:
	// false
}

func ExampleOf() {
	errs := []error{
		oops.New("an error", oops.With(example.Metadata{Retry: true})),
		oops.New("an error", kind.Code(1)),
		oops.New("an error", oops.Standard{Code: 2}),
		oops.New("an error", oops.With(example.MetaStandard{Code: 3})),
	}
	for i := range errs {
		fmt.Printf("%q with kind of %d\n", errs[i], kind.Of(errs[i]))
	}
	// Output:
	// "an error" with kind of -1
	// "an error" with kind of 1
	// "an error" with kind of 2
	// "an error" with kind of 3
}
