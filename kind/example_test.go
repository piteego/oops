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
		example.BuiltinErr,
		oops.New("an oops meta error", oops.With(example.Metadata{Retry: true})),
		oops.New("an oops kind error", kind.Code(1)),
		oops.New("an oops standard error", oops.Standard(2, 0, nil)),
		oops.New("an oops meta standard error", oops.With(example.MetaStandard{Code: 3})),
	}
	for i := range errs {
		fmt.Printf("%q with kind of %d\n", errs[i], kind.Of(errs[i]))
	}
	// Output:
	// "a builtin error" with kind of -1
	// "an oops meta error" with kind of -1
	// "an oops kind error" with kind of 1
	// "an oops standard error" with kind of 2
	// "an oops meta standard error" with kind of 3
}
