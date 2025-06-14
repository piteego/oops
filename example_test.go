package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
)

func ExampleNew() {
	type MyMeta struct {
		oops.Metadata `json:"-"`
		Code          int
	}
	errs := []error{
		oops.New("this is a basic error: no meta, no cause, no kind"),
		oops.New("this is a basic error: zero options are skipped",
			nil, &oops.Tag{}, oops.Metadata{}, &oops.Metadata{},
		),
		oops.New("this is a standard error including cause, and kind errors",
			&oops.Tag{Cause: errors.New("this is a cause"), Kind: errors.New("this is a kind")},
		),
		oops.New("this is a meta error including client custom metadata", MyMeta{Code: 42}),
		oops.New("this is a rich error: standard + meta", MyMeta{Code: 42},
			&oops.Tag{Cause: errors.New("this is a cause"), Kind: errors.New("this is a kind")},
		),
	}
	for i := range errs {
		fmt.Printf("%+v\n", errs[i])
	}
	// Output:
	//this is a basic error: no meta, no cause, no kind
	//this is a basic error: zero options are skipped
	//this is a standard error including cause, and kind errors
	//this is a meta error including client custom metadata
	//this is a rich error: standard + meta
}
