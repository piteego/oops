package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
)

func ExampleNew() {
	kind := errors.New("this is a kind")
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
			&oops.Tag{Kind: kind, Cause: errors.New("this is a cause")},
		),
		oops.New("this is a meta error including client custom metadata", MyMeta{Code: 42}),
		oops.New("this is a rich error: standard + meta", MyMeta{Code: 42},
			&oops.Tag{Kind: kind, Cause: errors.New("this is a cause")},
		),
	}
	for i := range errs {
		fmt.Println(errs[i])
		if errors.Is(errs[i], kind) {
			fmt.Println("errors.Is(err, kind) = true")
		}
	}
	// Output:
	// this is a basic error: no meta, no cause, no kind
	// this is a basic error: zero options are skipped
	// this is a standard error including cause, and kind errors
	// errors.Is(err, kind) = true
	// this is a meta error including client custom metadata
	// this is a rich error: standard + meta
	// errors.Is(err, kind) = true
}

func ExampleDiag() {
	type MyMeta struct {
		oops.Metadata
		Diag oops.Diag
	}
	errs := []error{
		oops.New("an error including diagnostic note, and low severity level",
			MyMeta{Diag: oops.Low.Diag("custom diag note...")},
		),
		oops.New("an error including diagnostic note, and medium severity level",
			MyMeta{Diag: oops.Medium.Diag("custom diag note...")},
		),
		oops.New("an error including diagnostic note, and high severity level",
			MyMeta{Diag: oops.High.Diag("custom diag note...")},
		),
		oops.New("an error including diagnostic note, and critical severity level",
			MyMeta{Diag: oops.Critical.Diag("custom diag note...")},
		),
	}
	for i := range errs {
		fmt.Println(errs[i])
	}
	// Output:
	// an error including diagnostic note, and low severity level
	// an error including diagnostic note, and medium severity level
	// an error including diagnostic note, and high severity level
	// an error including diagnostic note, and critical severity level
}
