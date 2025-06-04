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
	Unimplemented = custom{0, errors.New("not implemented yet"), "This feature is not implemented yet."}
	Internal      = custom{1, errors.New("something went wrong"), "An internal error occurred. Please try again later."}

	Unauthorized = custom{10, errors.New("unauthorized access"), "You are not authorized to perform this action."}
	Forbidden    = custom{11, errors.New("forbidden access"), "You do not have permission to access this resource."}

	Unprocessable = custom{20, errors.New("the request is unprocessable"), "The request could not be processed due to semantic errors."}
	Validation    = custom{21, errors.New("invalid input"), "The input provided is invalid. Please check your data and try again."}

	NotFound    = custom{30, errors.New("resource not found"), "The requested resource was not found. Please check the identifier and try again."}
	Duplication = custom{31, errors.New("duplicate entry"), "The entry already exists. Please check for duplicates and try again."}
)
