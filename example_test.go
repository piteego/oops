package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"strconv"
)

func ExampleLabel() {
	ErrInternal := oops.Label(errors.New("something went wrong"))
	ErrNotFound := oops.Label(errors.New("resource not found"))
	ErrForbidden := oops.Label(errors.New("forbidden access"))
	ErrValidation := oops.Label(errors.New("invalid input"))
	ErrDuplication := oops.Label(errors.New("duplicate entry"))
	ErrUnauthorized := oops.Label(errors.New("unauthorized access"))
	ErrUnimplemented := oops.Label(errors.New("not implemented yet"))
	ErrUnprocessable := oops.Label(errors.New("the request is unprocessable"))
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
	err := oops.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
		var oopsErr *oops.Error
		fmt.Println(errors.As(err, &oopsErr))
		fmt.Println(errors.Is(err, oops.Untagged))
	}
	// Output:
	// emit macho dwarf: elf header corrupted
	// true
	// true
}

func ExampleNew_tagCustomLabel() {
	customLabel := oops.Label(errors.New("custom error label"))
	err := oops.New("emit macho dwarf: elf header corrupted", oops.Tag(customLabel))
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, customLabel))
		var oopsErr *oops.Error
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
	err := oops.New("emit macho dwarf: elf header corrupted",
		oops.Because(failedProcess()...),
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
	err := oops.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// emit macho dwarf: elf header corrupted
}

func ExampleError_Unwrap() {
	customLabel := oops.Label(errors.New("custom error label"))
	err := oops.New("emit macho dwarf: elf header corrupted",
		oops.Tag(customLabel),
		oops.Because(
			errors.New("cause error 1"),
			errors.New("cause error 2"),
		),
	)
	if err != nil {
		var oopsErr *oops.Error
		if errors.As(err, &oopsErr) {
			fmt.Printf("%q\n", oopsErr.Unwrap())
		}
	}
	// Output:
	// ["cause error 1" "cause error 2" "custom error label"]
}

func ExampleTag() {
	err := oops.New("emit macho dwarf: elf header corrupted",
		oops.Tag(example.Internal.Error),
		oops.Tag(example.NotFound.Error), // multiple tags are not merged, the first tag is used
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
	err := oops.New("emit macho dwarf: elf header corrupted",
		oops.Because(
			reasons..., // multiple causes are merged into a single oops.Error's stack trace
		),
	)
	if err != nil {
		var oopsErr *oops.Error
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
	var handle oops.Handler = func(err error) *oops.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, err1) {
			return oops.New("handled error.1", oops.Tag(example.Forbidden.Error), oops.Because(err1)).(*oops.Error)
		}
		if errors.Is(err, err2) {
			return oops.New("handled error.2", oops.Tag(example.NotFound.Error), oops.Because(err2)).(*oops.Error)
		}
		return oops.New("unknown error", oops.Tag(example.Internal.Error), oops.Because(err)).(*oops.Error)
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
	var handler oops.Handler = func(err error) *oops.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, err1) {
			return oops.New("handled error.1", oops.Tag(example.Forbidden.Error), oops.Because(err1)).(*oops.Error)
		}
		if errors.Is(err, err2) {
			return oops.New("handled error.2", oops.Tag(example.NotFound.Error), oops.Because(err2)).(*oops.Error)
		}
		return oops.New("unknown error", oops.Tag(example.Internal.Error), oops.Because(err)).(*oops.Error)
	}

	fmt.Println(oops.Handle(err1, handler))
	fmt.Println(oops.Handle(err2, handler))
	fmt.Println(oops.Handle(errors.New("unhandled error"), handler))
	fmt.Println(oops.Handle(nil, handler)) // should return nil
	fmt.Println(oops.Handle(errors.New("an error with nil handler"), nil))
	fmt.Println(oops.Handle(oops.New("already an oops error", oops.Tag(example.Validation.Error))))
	// Output:
	// handled error.1
	// handled error.2
	// unknown error
	// <nil>
	// an error with nil handler
	// already an oops error
}
