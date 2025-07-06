package oops_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops"
)

func ExampleNew_standardOptions() {
	err := oops.New("this is a message",
		oops.WithKind(5),
		oops.WithSeverity(2),
		oops.WithCause(errors.New("this is a cause")),
	)
	fmt.Println(err)
	fmt.Println("code:", oops.GetMetadata[oops.Code](err))
	fmt.Println("severity:", oops.GetMetadata[oops.Level](err))
	fmt.Println("cause:", oops.GetMetadata[oops.Cause](err))
	// Output:
	// this is a message
	// code: 5
	// severity: 2
	// cause: this is a cause
}

func ExampleNew_withClientCustomMetadata() {
	type custom struct {
		Id    string
		Retry bool
	}
	err := oops.SetMetadata(errors.New("this is a message"), custom{Id: "E10", Retry: true})
	fmt.Printf("%q with custom:%+v\n", err, oops.GetMetadata[custom](err))
	// Output:
	// "this is a message" with custom:{Id:E10 Retry:true}
}

func ExampleNew_withStandardOptionsAndClientCustomMetadata() {
	type custom struct {
		Id    string
		Retry bool
	}
	err := oops.SetMetadata(
		oops.New("this is a message",
			oops.WithCause(errors.New("this is a cause")),
			oops.WithKind(5),
			oops.WithSeverity(2),
		), custom{Id: "E10", Retry: true},
	)
	fmt.Println(err)
	fmt.Println("code:", oops.GetMetadata[oops.Code](err))
	fmt.Println("severity:", oops.GetMetadata[oops.Level](err))
	fmt.Println("cause:", oops.GetMetadata[oops.Cause](err))
	fmt.Printf("custom:%+v\n", oops.GetMetadata[custom](err))
	// Output:
	// this is a message
	// code: 5
	// severity: 2
	// cause: this is a cause
	// custom:{Id:E10 Retry:true}
}
