package oops_test

import (
	"errors"
	"github.com/piteego/oops"
)

func ExampleNew() {
	type MyMeta struct {
		oops.Metadata `json:"-"`
		Code          int
	}
	errs := []error{
		oops.New("this is a basic error: no meta, no cause, no label"),
		oops.New("this is a basic error: zero values are skipped",
			nil, &oops.Tag{}, oops.Metadata{},
			&oops.Metadata{}, &oops.Tag{},
		),
		oops.New("this is a standard error including cause, label, and diagnostic detail",
			&oops.Tag{Cause: errors.New("this is a cause")},
		),
		oops.New("this is a meta error including client custom metadata", MyMeta{Code: 42}),
		oops.New("this is a rich error: standard + meta",
			&oops.Tag{Cause: errors.New("this is a cause")},
			MyMeta{Code: 42},
		),
	}
	// Output:
	// oops! this is a basic error: no meta, no cause, no label
	// oops! this is a basic error: zero values are skipped
	// oops.StandardErr{msg:this is a standard error including cause, label, and diagnostic detail Standard:{Label:<nil> Cause:this is a cause Diag:{Note: Severity:0}}}
	// oops.MetaErr{msg:this is a meta error including client custom metadata Data:{Metadata:{} Code:42}}
	// oops.RichErr{msg:this is a rich error: standard + meta Standard:{Label:<nil> Cause:this is a cause Diag:{Note: Severity:0}} Data:{Metadata:{} Code:42}}
}
