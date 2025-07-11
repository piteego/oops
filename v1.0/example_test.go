package v1_0_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
	"github.com/piteego/oops/v1.0"
)

func ExampleNew_withKindOption() {
	errs := []error{
		v1_0.New("an error with kind of", kind.Unauthorized),
		v1_0.New("an error with sanitized kind of", kind.Code(7)),
		v1_0.New("an error with sanitized kind of", kind.Code(-2)),
		errors.New("a builtin error's kind considered as"),
	}
	for i := range errs {
		fmt.Printf("✓ %s %q\n", errs[i], kind.Of(errs[i]))
	}
	// Output:
	// ✓ an error with kind of "Unauthorized Error"
	// ✓ an error with sanitized kind of "Unknown Error"
	// ✓ an error with sanitized kind of "Unknown Error"
	// ✓ a builtin error's kind considered as "Unknown Error"
}

func ExampleNew_withSeverityLevelOption() {
	errs := []error{
		v1_0.New("an error with severity level of", severity.Critical),
		v1_0.New("an error with sanitized severity level of", severity.Level(6)),
		v1_0.New("an error with sanitized severity level of", severity.Level(7)),
		errors.New("a builtin error's severity level considered as"),
	}
	for i := range errs {
		fmt.Printf("✓ %s %q\n", errs[i], severity.Of(errs[i]))
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
		v1_0.New("an error with cause of", v1_0.CausedBy(example.ErrorCause)),
		errors.New("a builtin error's cause is"),
	}
	for i := range errs {
		fmt.Printf("✓ %s \"%v\"\n", errs[i], v1_0.CauseOf(errs[i]))
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
	_ = v1_0.New("✗ an error with cause of nil will panic:", v1_0.CausedBy(nil))
	// Output:
	// ✗ panic: trying to set nil cause

}

func ExampleNew_withStandardOptions() {
	errs := []error{
		v1_0.New("an error with", v1_0.Standard{Code: kind.NotFound, Level: 5, CausedBy: example.OsErrNotExist}),
		v1_0.New("an error with", v1_0.Standard{Code: kind.Duplication, Level: 1, CausedBy: example.GormErrDuplicatedKey}),
		v1_0.New("an error with", v1_0.Standard{Code: -1, Level: 6, CausedBy: example.ErrorCause}),
	}
	for i := range errs {
		fmt.Printf("✓ %v {kind: %q, level: %q, cause: \"%v\"}\n",
			errs[i], kind.Of(errs[i]), severity.Of(errs[i]), v1_0.CauseOf(errs[i]),
		)
	}
	// Output:
	// ✓ an error with {kind: "Not Found Error", level: "Informational", cause: "file does not exist"}
	// ✓ an error with {kind: "Duplication Error", level: "Critical", cause: "gorm duplicated key"}
	// ✓ an error with {kind: "Unknown Error", level: "Unknown", cause: "a cause error"}
}

func ExampleNew_withCustomMetadata() {
	err := v1_0.New("an error", v1_0.Append(example.Metadata{Id: "E109", Retry: true}))
	metadata := v1_0.Unwrap[example.Metadata](err)
	fmt.Printf("✓ %s with custom metadata {Id: %q, Retry: %v}", err, metadata.Id, metadata.Retry)
	// Output:
	// ✓ an error with custom metadata {Id: "E109", Retry: true}
}

func ExampleAppend_directlyThenUnwrap() {
	err := v1_0.Append(example.Metadata{Id: "E109", Retry: false})(
		v1_0.New("an error", kind.Forbidden),
	)
	metadata := v1_0.Unwrap[example.Metadata](err)
	fmt.Printf("%q with kind of %q, and custom metadata {Id: %q, Retry: %v}\n",
		err,
		kind.Of(err),
		metadata.Id, metadata.Retry,
	)
	// Output:
	// "an error" with kind of "Forbidden Error", and custom metadata {Id: "E109", Retry: false}
}

func ExampleAppend_asOptionOfNew() {
	err := v1_0.New("an error",
		v1_0.Append(
			example.MetaKindAndSeverity{Code: kind.Validation, Retry: true},
		),
	)
	metadata := v1_0.Unwrap[example.MetaKind](err)
	fmt.Printf("%q with kind of %q, and custom metadata {Code: %d, Retry: %v}\n",
		err,
		kind.Of(err),
		metadata.Code, metadata.Retry,
	)
	// Output
	// "an error" with kind of "Validation Error", and custom metadata {Code: 2, Retry: true}
}
