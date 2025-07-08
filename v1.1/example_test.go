package v1_1_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/v1.1"
)

func ExampleNew() {
	kind := errors.New("this is a kind")
	type MyMeta struct {
		v1_1.Metadata `json:"-"`
		Code          int
	}
	errs := []error{
		v1_1.New("this is a basic error: no meta, no cause, no kind"),
		v1_1.New("this is a basic error: zero options are skipped",
			nil, &v1_1.Tag{}, v1_1.Metadata{}, &v1_1.Metadata{},
		),
		v1_1.New("this is a standard error including cause, and kind errors",
			&v1_1.Tag{Kind: kind, Cause: errors.New("this is a cause")},
		),
		v1_1.New("this is a meta error including client custom metadata", MyMeta{Code: 42}),
		v1_1.New("this is a rich error: standard + meta", MyMeta{Code: 42},
			&v1_1.Tag{Kind: kind, Cause: errors.New("this is a cause")},
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
		v1_1.Metadata
		Diag v1_1.Diag
	}
	errs := []error{
		v1_1.New("an error including diagnostic note, and low severity level",
			MyMeta{Diag: v1_1.Low.Diag("custom diag note...")},
		),
		v1_1.New("an error including diagnostic note, and medium severity level",
			MyMeta{Diag: v1_1.Medium.Diag("custom diag note...")},
		),
		v1_1.New("an error including diagnostic note, and high severity level",
			MyMeta{Diag: v1_1.High.Diag("custom diag note...")},
		),
		v1_1.New("an error including diagnostic note, and critical severity level",
			MyMeta{Diag: v1_1.Critical.Diag("custom diag note...")},
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
