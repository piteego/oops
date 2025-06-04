package example

import (
	"errors"
	"github.com/piteego/oops"
)

// custom is an example struct that uses the oops.Label to define custom error categories.
type custom struct {
	Code        int
	Error       oops.Label
	Description string
	// more fields can be added as needed, for example:
	// httpStatus int
	// grpcCode   int
	// debug string
	// ...
}

var (
	Unimplemented = custom{0, oops.Label(errors.New("not implemented yet")), "This feature is not implemented yet."}
	Internal      = custom{1, oops.Label(errors.New("something went wrong")), "An internal error occurred."}

	Unauthorized = custom{10, oops.Label(errors.New("unauthorized access")), "You are not authorized to perform this action."}
	Forbidden    = custom{11, oops.Label(errors.New("forbidden access")), "You do not have permission to access this resource."}

	Unprocessable = custom{20, oops.Label(errors.New("the request is unprocessable")), "The request could not be processed."}
	Validation    = custom{21, oops.Label(errors.New("invalid input")), "The input provided is invalid."}

	NotFound    = custom{30, oops.Label(errors.New("resource not found")), "The requested resource was not found."}
	Duplication = custom{31, oops.Label(errors.New("duplicate entry")), "The entry already exists."}
)
