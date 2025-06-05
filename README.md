# oops: Structured Error Handling for Go Clean Architecture
[![License](http://img.shields.io/:license-Apache_2.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/piteego/oops)](https://godoc.org/github.com/piteego/oops)
[![Coverage Status](https://coveralls.io/repos/github/piteego/oops/badge.svg?branch=main)](https://coveralls.io/github/piteego/oops?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/piteego/oops)](https://goreportcard.com/report/github.com/piteego/oops)

**`oops`** provides a straightforward and structured approach to error handling specifically designed for Go applications adopting a clean architecture. It empowers you to create, categorize, and manage errors effectively using a system of labels and handlers, thereby offering enhanced debuggability, maintainability, and clear separation of concerns across your architectural layers. This approach respects the Dependency Rule, ensuring that inner, higher-level policy layers remain independent of outer, lower-level implementation details, thus preventing undesirable import cycles and promoting a robust, testable codebase.

---

## Installation

```bash
go get github.com/piteego/oops
```

---

## Features

* **Categorize Errors with [`Label`](https://pkg.go.dev/github.com/piteego/oops#Label)**: 
 
    Define custom error categories using `Label` errors. This allows you to classify application errors consistently. Examples demonstrating how to define these custom categories can be found in the [`example`](https://pkg.go.dev/github.com/piteego/oops/example) package.


* **Create Labeled Errors**: 

    Use the [`New`](https://pkg.go.dev/github.com/piteego/oops#New) function to create new errors and associate them with your predefined labels. For instance:

    ```go
    import (
        "errors"
        "fmt"

        "[github.com/piteego/oops](https://github.com/piteego/oops)"
        "[github.com/piteego/oops/example](https://github.com/piteego/oops/example)" // Assuming your example package is here
    )

    func processData() error {
        // Simulate a data processing error that is a duplicate
        if true { // Replace with actual condition
            return oops.New("failed to process", oops.Tag(example.Duplicated.Error))
        }
        return nil
    }

    func main() {
        err := processData()
        if err != nil {
            if errors.Is(err, example.Duplicated.Error) {
                fmt.Println("Caught a duplicated error:", err)
            } else {
                fmt.Println("An unexpected error occurred:", err)
            }
        }
    }
    ```

* **Flexible Error Options**:

  [`ErrorOption`](https://pkg.go.dev/github.com/piteego/oops#ErrorOption) is a function type that modifies an [`Error`](https://pkg.go.dev/github.com/piteego/oops#Error) instance. It allows you to set options like tagging the error with a [`Label`](https://pkg.go.dev/github.com/piteego/oops#Label) or adding a stack trace with [`Because`](https://pkg.go.dev/github.com/piteego/oops#Because).


* **Stack Traces**: 

  Use [`Because`](https://pkg.go.dev/github.com/piteego/oops#Because) in the [`New`](https://pkg.go.dev/github.com/piteego/oops#New) function to append stack traces to your errors, providing valuable context for debugging.


* **Structured Error Handling**:

    * **[`errors.Is`](https://pkg.go.dev/errors#Is)**: Handle errors in higher layers of your application using `errors.Is` to check against specific labeled errors:

        ```go
        if errors.Is(err, example.Duplicated.Error) {
            // Handle the duplicated error specifically
        }
        ```

    * **[`Map`](https://pkg.go.dev/github.com/piteego/oops#Map) Type**: 
  
      The [`Map`](https://pkg.go.dev/github.com/piteego/oops#Map) type provides a structured way to handle built-in errors. Just define a map of errors to their corresponding `*Error` instances, and use the `Map.Handle` method to process errors. 
  
      The `Handle` method will append the original error to the stack of the returned `*Error`.
      ```go
        import (
            "errors"
            "github.com/piteego/oops"
            "github.com/piteego/oops/example"
        )
        func main() {
           err := example.ErrMap.Handle(example.GormErrRecordNotFound)
           
           _ = err.(*oops.Error) // will not panic
           
           errors.Is(err, example.GormErrNotFound) // true 
		   
           errors.Is(err, example.NotFound.Error) // true
        }
      ```
      
    * **Custom [`Handler`](https://pkg.go.dev/github.com/piteego/oops#Handler)**:
  
      Define complex `Handler` functions in different application layers. The [`Handle`](https://pkg.go.dev/github.com/piteego/oops#Handle) function can then be used to process a given error by invoking a series of these custom handlers.
      ```go
        import (
            "errors"
            "github.com/piteego/oops"
            "github.com/piteego/oops/example"
        )
    
        func main() {
            fmt.Println(example.HandleRepoErr("user")(example.GormErrRecordNotFound)) // user not found
            fmt.Println(example.HandleRepoErr("user")(example.RedisCacheMissed)) // user not found
            fmt.Println(example.HandleRepoErr("user")(errors.New("unhandled error"))) // something went wrong
        }