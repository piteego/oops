package v0_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v0"
	"strconv"
)

func ExampleLabel() {
	ErrInternal := v0.Label(errors.New("something went wrong"))
	ErrNotFound := v0.Label(errors.New("resource not found"))
	ErrForbidden := v0.Label(errors.New("forbidden access"))
	ErrValidation := v0.Label(errors.New("invalid input"))
	ErrDuplication := v0.Label(errors.New("duplicate entry"))
	ErrUnauthorized := v0.Label(errors.New("unauthorized access"))
	ErrUnimplemented := v0.Label(errors.New("not implemented yet"))
	ErrUnprocessable := v0.Label(errors.New("the request is unprocessable"))
	fmt.Println(ErrInternal)
	fmt.Println(ErrNotFound)
	fmt.Println(ErrForbidden)
	fmt.Println(ErrValidation)
	fmt.Println(ErrDuplication)
	fmt.Println(ErrUnauthorized)
	fmt.Println(ErrUnimplemented)
	fmt.Println(ErrUnprocessable)
	// Output:
	// something went wrong
	// resource not found
	// forbidden access
	// invalid input
	// duplicate entry
	// unauthorized access
	// not implemented yet
	// the request is unprocessable
}

func ExampleNew_withNoOptions() {
	err := v0.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
		var oopsErr *v0.Error
		fmt.Println(errors.As(err, &oopsErr))
		fmt.Println(errors.Is(err, v0.Untagged))
	}
	// Output:
	// emit macho dwarf: elf header corrupted
	// true
	// true
}

func ExampleNew_tagCustomLabel() {
	customLabel := v0.Label(errors.New("custom error label"))
	err := v0.New("emit macho dwarf: elf header corrupted", v0.Tag(customLabel))
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, customLabel))
		var oopsErr *v0.Error
		fmt.Println(errors.As(err, &oopsErr))
		if oopsErr != nil {
			fmt.Printf("%q", oopsErr.Unwrap())
		}
	}
	// Output:
	// emit macho dwarf: elf header corrupted
	// true
	// true
	// ["custom error label"]
}

func ExampleNew_causedByStackErrors() {
	var ErrProcessFailed = errors.New("process failed")
	failedProcess := func() []error {
		_, a2iErr := strconv.Atoi("invalid data")
		return []error{ErrProcessFailed, a2iErr}
	}
	err := v0.New("emit macho dwarf: elf header corrupted",
		v0.Because(failedProcess()...),
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrProcessFailed))
		fmt.Println(errors.Is(err, strconv.ErrSyntax))
	}
	// Output:
	// emit macho dwarf: elf header corrupted
	// true
	// true
}

func ExampleError_Error() {
	err := v0.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// emit macho dwarf: elf header corrupted
}

func ExampleError_Unwrap() {
	customLabel := v0.Label(errors.New("custom error label"))
	err := v0.New("emit macho dwarf: elf header corrupted",
		v0.Tag(customLabel),
		v0.Because(
			errors.New("cause error 1"),
			errors.New("cause error 2"),
		),
	)
	if err != nil {
		var oopsErr *v0.Error
		if errors.As(err, &oopsErr) {
			fmt.Printf("%q\n", oopsErr.Unwrap())
		}
	}
	// Output:
	// ["cause error 1" "cause error 2" "custom error label"]
}

func ExampleTag() {
	err := v0.New("emit macho dwarf: elf header corrupted",
		v0.Tag(example.Internal.Error),
		v0.Tag(example.NotFound.Error), // multiple tags are not merged, the first tag is used
	)
	if err != nil {
		fmt.Println(errors.Is(err, example.Internal.Error))
		fmt.Println(errors.Is(err, example.NotFound.Error))
	}
	// Output:
	// true
	// false
}

func ExampleBecause() {
	reasons := []error{
		errors.New("cause error 1"),
		errors.New("cause error 2"),
	}
	err := v0.New("emit macho dwarf: elf header corrupted",
		v0.Because(
			reasons..., // multiple causes are merged into a single oops.Error's stack trace
		),
	)
	if err != nil {
		var oopsErr *v0.Error
		if errors.As(err, &oopsErr) {
			fmt.Printf("%q\n", oopsErr.Unwrap())
		}
		for i := range reasons {
			fmt.Println(errors.Is(err, reasons[i]))
		}
	}
	// Output:
	// ["cause error 1" "cause error 2" "untagged"]
	// true
	// true
}

func ExampleMap() {
	fmt.Println(example.ErrMap.Handle(example.RedisCacheMissed))
	fmt.Println(example.ErrMap.Handle(errors.New("unhandled error")))
	// Output:
	// cache key not found
	// unhandled error
}

func ExampleHandler_closure() {
	fmt.Println(example.HandleRepoErr("user")(example.GormErrRecordNotFound))
	fmt.Println(example.HandleRepoErr("user")(example.RedisCacheMissed))
	fmt.Println(example.HandleRepoErr("user")(errors.New("unhandled error")))
	// Output:
	// user not found
	// user not found
	// something went wrong
}

func ExampleHandler_asVariable() {
	err1 := errors.New("error.1")
	err2 := errors.New("error.2")
	// ... and some other errors in lower layers
	var handle v0.Handler = func(err error) *v0.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, err1) {
			return v0.New("handled error.1", v0.Tag(example.Forbidden.Error), v0.Because(err1)).(*v0.Error)
		}
		if errors.Is(err, err2) {
			return v0.New("handled error.2", v0.Tag(example.NotFound.Error), v0.Because(err2)).(*v0.Error)
		}
		return v0.New("unknown error", v0.Tag(example.Internal.Error), v0.Because(err)).(*v0.Error)
	}
	oopsErr := handle(err1)
	fmt.Println(oopsErr)
	fmt.Printf("%q\n", oopsErr.Unwrap())
	// Output:
	// handled error.1
	// ["error.1" "forbidden access"]
}

func ExampleHandle() {
	err1 := errors.New("error.1")
	err2 := errors.New("error.2")
	// ... and some other errors in lower layers
	var handler v0.Handler = func(err error) *v0.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, err1) {
			return v0.New("handled error.1", v0.Tag(example.Forbidden.Error), v0.Because(err1)).(*v0.Error)
		}
		if errors.Is(err, err2) {
			return v0.New("handled error.2", v0.Tag(example.NotFound.Error), v0.Because(err2)).(*v0.Error)
		}
		return v0.New("unknown error", v0.Tag(example.Internal.Error), v0.Because(err)).(*v0.Error)
	}

	fmt.Println(v0.Handle(err1, handler))
	fmt.Println(v0.Handle(err2, handler))
	fmt.Println(v0.Handle(errors.New("unhandled error"), handler))
	fmt.Println(v0.Handle(nil, handler)) // should return nil
	fmt.Println(v0.Handle(errors.New("an error with nil handler"), nil))
	fmt.Println(v0.Handle(v0.New("already an oops error", v0.Tag(example.Validation.Error))))
	// Output:
	// handled error.1
	// handled error.2
	// unknown error
	// <nil>
	// an error with nil handler
	// already an oops error
}
