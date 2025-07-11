package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
)

func ExampleNew_withKindOption() {
	errs := []error{
		oops.New("an error with kind of", oops.Unauthorized),
		oops.New("an error with sanitized kind of", oops.Kind(8)),
		oops.New("an error with sanitized kind of", oops.Kind(-2)),
		errors.New("a builtin error's kind considered as"),
	}
	for i := range errs {
		fmt.Printf("✓ %s %q\n", errs[i], oops.KindOf(errs[i]))
	}
	// Output:
	// ✓ an error with kind of "Unauthorized Error"
	// ✓ an error with sanitized kind of "Unknown Error"
	// ✓ an error with sanitized kind of "Unknown Error"
	// ✓ a builtin error's kind considered as "Unknown Error"
}

func ExampleNew_withSeverityLevelOption() {
	errs := []error{
		oops.New("an error with severity level of", oops.Critical),
		oops.New("an error with sanitized severity level of", oops.Level(6)),
		oops.New("an error with sanitized severity level of", oops.Level(7)),
		errors.New("a builtin error's severity level considered as"),
	}
	for i := range errs {
		fmt.Printf("✓ %s %q\n", errs[i], oops.LevelOf(errs[i]))
	}
	// Output:
	// ✓ an error with severity level of "Critical"
	// ✓ an error with sanitized severity level of "Unknown"
	// ✓ an error with sanitized severity level of "Unknown"
	// ✓ a builtin error's severity level considered as "Unknown"
}

func ExampleNew_withCauseOption() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()
	errs := []error{
		oops.New("an error with cause of", oops.CausedBy(example.ErrorCause)),
		errors.New("a builtin error's cause is"),
	}
	for i := range errs {
		fmt.Printf("✓ %s \"%v\"\n", errs[i], oops.CauseOf(errs[i]))
	}
	// Output:
	// ✓ an error with cause of "a cause error"
	// ✓ a builtin error's cause is "<nil>"
}

func ExampleNew_withCauseNil() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("✗ panic: %v\n", rec)
		}
	}()
	_ = oops.New("✗ an error with cause of nil will panic:", oops.CausedBy(nil))
	// Output:
	// ✗ panic: trying to set nil cause
}

func ExampleNew_withStandardOptions() {
	errs := []error{
		oops.New("an error with", oops.Standard(oops.NotFound, 5, example.OsErrNotExist)),
		oops.New("an error with", oops.Standard(oops.Duplication, 1, example.GormErrDuplicatedKey)),
		oops.New("an error with", oops.Standard(-1, 6, example.ErrorCause)),
	}
	for i := range errs {
		fmt.Printf("✓ %v {kind: %q, level: %q, cause: \"%v\"}\n",
			errs[i], oops.KindOf(errs[i]), oops.LevelOf(errs[i]), oops.CauseOf(errs[i]),
		)
	}
	// Output:
	// ✓ an error with {kind: "Not Found Error", level: "Informational", cause: "file does not exist"}
	// ✓ an error with {kind: "Duplication Error", level: "Critical", cause: "gorm duplicated key"}
	// ✓ an error with {kind: "Unknown Error", level: "Unknown", cause: "a cause error"}
}

func ExampleNew_withCustomMetadata() {
	err := oops.New("an error", example.Metadata{Id: "E109", Retry: true})
	metadata := new(example.Metadata)
	if oops.As(err, metadata) {
		fmt.Printf("✓ %s with custom metadata {Id: %q, Retry: %v}\n", err, metadata.Id, metadata.Retry)
	}
	kind := new(oops.Kind)
	if !oops.As(err, kind) {
		fmt.Printf("✓ %s with kind of %q\n", err, *kind)
	}
	// Output:
	// ✓ an error with custom metadata {Id: "E109", Retry: true}
	// ✓ an error with kind of "Unknown Error"
}
