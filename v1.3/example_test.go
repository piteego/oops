package v1_3_test

import (
	"errors"
	"fmt"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
	"github.com/piteego/oops/v1.3"
)

func ExampleNew_withStandardOptions() {
	err := v1_3.New("this is a message",
		v1_3.WithKind(5),
		v1_3.WithSeverity(2),
		v1_3.WithCause(errors.New("this is a cause")),
	)
	fmt.Println(err)
	fmt.Println("code:", v1_3.Get[kind.Code](err))
	fmt.Println("severity:", v1_3.Get[severity.Level](err))
	fmt.Println("cause:", v1_3.Get[v1_3.Cause](err).Error)
	// Output:
	// this is a message
	// code: 5
	// severity: 2
	// cause: this is a cause
}

//func ExampleNew_withClientMetadata() {
//	type custom struct {
//		oops.Metadata
//		Id    string
//		Retry bool
//	}
//	err := oops.New("this is a message", oops.WithMetadata(custom{Id: "E10", Retry: true}))
//	fmt.Printf("%q with custom:%+v\n", err, oops.Get[custom](err))
//	// Output:
//	// "this is a message" with custom:{Id:E10 Retry:true}
//}

//func ExampleNew_withStandardOptionsAndClientMetadata() {
//	type custom struct {
//		oops.metadata
//		Id    string
//		Retry bool
//	}
//	err := oops.New("this is a message",
//		oops.WithCause(errors.New("this is a cause")),
//		oops.WithKind(5),
//		oops.WithSeverity(2),
//		oops.WithMetadata(custom{Id: "E10", Retry: true}),
//	)
//	fmt.Println(err)
//	fmt.Println("code:", oops.Get[kind.Code](err))
//	fmt.Println("severity:", oops.Get[severity.Level](err))
//	fmt.Println("cause:", oops.Get[oops.Cause](err).Error)
//	fmt.Printf("custom:%+v\n", oops.Get[custom](err))
//	// Output:
//	// this is a message
//	// code: 5
//	// severity: 2
//	// cause: this is a cause
//	// custom:{Id:E10 Retry:true}
//}
